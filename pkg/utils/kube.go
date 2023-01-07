package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type ClusterData struct {
	Name   string     `json:"name"`
	Errors []string   `json:"errors"`
	Nodes  []NodeData `json:"nodes"`
}

type NodeData struct {
	Name              string            `json:"name"`
	IP                string            `json:"ip"`
	LastCheckTime     string            `json:"LastCheckTime"`
	DnsHealthy        bool              `json:"dnsHealthy"`
	DnsHealthStatus   map[string]string `json:"dnsHealthStatus"`
	PodNetworkHealthy bool              `json:"podNetworkHealthy"`
	PodNetworkStatus  map[string]string `json:"podNetworkStatus"`
}

var Client *clientset.Clientset
var err error

func init() {
	// TEST OTHER ENV VARS
	ip, _ := os.LookupEnv("POD_IP")
	if ip == "" {
		log.Println("POD_IP ENV VAR NOT SET")
	}
	node, _ := os.LookupEnv("NODE_NAME")
	if node == "" {
		log.Println("NODE_NAME ENV VAR NOT SET")
	}
	ns, _ := os.LookupEnv("NAMESPACE")
	if ns == "" {
		log.Println("NAMESPACE ENV VAR NOT SET")
	}

}

func NewConfig(kubeconfig string) (*rest.Config, error) {

	if kubeconfig != "" {
		log.Println("KUBECONFIG Env Var Set: Using out of cluster config")
		cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}

	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	log.Println("KUBECONFIG Env Var NOT Set:  using in cluster config")
	return cfg, nil
}

func GetCoreDNSPods(k8client *clientset.Clientset) (*v1.PodList, error) {
	pods, err := k8client.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{LabelSelector: "k8s-app=kube-dns"})
	if err != nil {
		return nil, err
	}
	return pods, nil
}

func GetECPPods(k8client *clientset.Clientset) (*v1.PodList, error) {

	pods, err := k8client.CoreV1().Pods("ecp-tools").List(context.TODO(), metav1.ListOptions{LabelSelector: "***REMOVED***/component=node-ds"})
	if err != nil {
		return nil, err
	}
	return pods, nil
}

func GetClusterHealth(k8client *clientset.Clientset, c *ClusterData) {
	nodes, err := k8client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	for _, n := range nodes.Items {
		tmpNode := NodeData{}
		name := n.Name
		annotations := n.GetAnnotations()
		nodeHealth := annotations["***REMOVED***/healthcheck"]
		json.Unmarshal([]byte(nodeHealth), &tmpNode)
		if nodeHealth != "" {
			log.Printf("Node: %v has status: %v", name, tmpNode)
			c.Nodes = append(c.Nodes, tmpNode)
		} else {
			log.Printf("Node: %s MISSING STATUS", name)
			errStr := fmt.Sprintf("ERROR: No healthcheck data for: %s", name)
			c.Errors = append(c.Errors, errStr)
		}
	}
}

func (n *NodeData) AnnotateStatus() {
	log.Printf("Setting Annotations for : %v", n.Name)
	node, err := Client.CoreV1().Nodes().Get(context.TODO(), n.Name, metav1.GetOptions{})
	if err != nil {
		log.Printf("ERROR: %v", err)
		return
	}
	annotations := node.GetAnnotations()
	jsonStr, err := json.Marshal(n)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	annotations["***REMOVED***/healthcheck"] = string(jsonStr)
	node.SetAnnotations(annotations)
	Client.CoreV1().Nodes().Update(context.TODO(), node, metav1.UpdateOptions{})
	log.Printf("Updated Annotation on Node: %v", n.Name)
}

func CreateK8sClient() *kubernetes.Clientset {
	log.Println("Setting up connection to Kubernetes API")
	kubeconfig, _ := os.LookupEnv("KUBECONFIG")
	log.Println("SOS")
	if kubeconfig == "" {
		log.Println("KUBECONFIG HERE")
		kubeconfig = filepath.Join(
			os.Getenv("HOME"), ".kube", "config",
		)
	}

	config, err := NewConfig(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	return clientset.NewForConfigOrDie(config)
}
