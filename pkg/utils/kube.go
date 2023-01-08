package utils

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var Client *clientset.Clientset
var err error

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

func CreateK8sClient() *kubernetes.Clientset {
	log.Println("Setting up connection to Kubernetes API")
	kubeconfig, _ := os.LookupEnv("KUBECONFIG")

	if kubeconfig == "" {
		log.Println("KUBECONFIG Variable not set. Checking default kubeconfig: $HOME.kube/config ")
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
