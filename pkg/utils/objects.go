package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

func (c *K8sObj) GetConfigMaps(Client *clientset.Clientset, ns string, st bool) {

	var tmpObjData staleObjData
	if !st {
		m, err := Client.CoreV1().ConfigMaps(ns).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("Error while fetching all Configmaps. Error: %v", err)
		}
		c.Configmaps = m.Items
	} else if st {
		log.Println("Getting all stale Configmaps")
		m, err := Client.CoreV1().ConfigMaps(ns).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("Error while fetching stale Configmaps. Error: %v", err)
		}
		for _, i := range m.Items {
			a := i.GetAnnotations()
<<<<<<< HEAD
			js := a["kubernetes.io/stale-object"]
=======
			js := a["***REMOVED***/stale-object"]
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
			json.Unmarshal([]byte(js), &tmpObjData)
			if js != "" && tmpObjData.ObjStatus == "NotInUse" {
				c.Configmaps = append(c.Configmaps, i)
			}
		}
	}
}

func (c *K8sObj) GetSecrets(Client *clientset.Clientset, ns string, st bool) {
	var tmpObjData staleObjData

	if !st {

		m, err := Client.CoreV1().Secrets(ns).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("Error while fetching all Secrets. Error: %v", err)
		}
		c.Secrets = m.Items
	} else if st {
		log.Println("Getting all stale Secrets")
		m, err := Client.CoreV1().Secrets(ns).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("Error while fetching stale Secrets. Error: %v", err)
		}
		for _, i := range m.Items {
			a := i.GetAnnotations()
<<<<<<< HEAD
			js := a["kubernetes.io/stale-object"]
=======
			js := a["***REMOVED***/stale-object"]
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
			json.Unmarshal([]byte(js), &tmpObjData)
			if js != "" && tmpObjData.ObjStatus == "NotInUse" {
				c.Secrets = append(c.Secrets, i)
			}
		}
	}
}

func (c *K8sObj) GetPVC(Client *clientset.Clientset, ns string, st bool) {
	var tmpObjData staleObjData

	if !st {

		m, err := Client.CoreV1().PersistentVolumeClaims(ns).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("Error while fetching all PersistentVolumeClaims. Error: %v", err)
		}
		c.PersistentVolumeClaims = m.Items
	} else if st {
		log.Println("Getting all stale PVCs")
		m, err := Client.CoreV1().PersistentVolumeClaims(ns).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("Error while fetching stale PersistentVolumeClaims. Error: %v", err)
		}
		for _, i := range m.Items {
			a := i.GetAnnotations()
<<<<<<< HEAD
			js := a["kubernetes.io/stale-object"]
=======
			js := a["***REMOVED***/stale-object"]
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
			json.Unmarshal([]byte(js), &tmpObjData)
			if js != "" && tmpObjData.ObjStatus == "NotInUse" {
				c.PersistentVolumeClaims = append(c.PersistentVolumeClaims, i)
			}
		}
	}
}

func (c *K8sObj) GetServices(Client *clientset.Clientset, ns string, st bool) {
	var tmpObjData staleObjData

	if !st {
		m, err := Client.CoreV1().Services(ns).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("Error while fetching all Services. Error: %v", err)
		}
		c.Services = m.Items
	} else if st {
		log.Println("Getting all stale Services")
		m, err := Client.CoreV1().Services(ns).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("Error while fetching stale Services. Error: %v", err)
		}

		for _, i := range m.Items {
			a := i.GetAnnotations()
<<<<<<< HEAD
			js := a["kubernetes.io/stale-object"]
=======
			js := a["***REMOVED***/stale-object"]
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
			json.Unmarshal([]byte(js), &tmpObjData)
			if js != "" && tmpObjData.ObjStatus == "NotInUse" {
				c.Services = append(c.Services, i)
			}
		}
	}
}

func (c *K8sObj) GetDeployment(Client *clientset.Clientset, ns string, st bool) {
	var tmpObjData staleObjData

	if !st {
		m, err := Client.AppsV1().Deployments(ns).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("Error while fetching all Deployments. Error: %v", err)
		}
		c.Deployments = m.Items
	} else if st {
		log.Println("Getting all stale Deployments")
		m, err := Client.AppsV1().Deployments(ns).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("Error while fetching all Deployments. Error: %v", err)
		}

		for _, i := range m.Items {
			a := i.GetAnnotations()
<<<<<<< HEAD
			js := a["kubernetes.io/stale-object"]
=======
			js := a["***REMOVED***/stale-object"]
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
			json.Unmarshal([]byte(js), &tmpObjData)
			if js != "" && tmpObjData.ObjStatus == "NotInUse" {
				c.Deployments = append(c.Deployments, i)
			}
		}
	}
}

func GetNamespaces(Client *clientset.Clientset) []string {
	var result []string
	ns, err := Client.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("Error fetching all namespaces: %v", err)
	}

	for _, i := range ns.Items {
		result = append(result, i.Name)
	}
	return result
}

func (c *K8sObj) GetPods(Client *clientset.Clientset, ns string, podState string) {
	pods, err := Client.CoreV1().Pods(ns).List(context.TODO(), v1.ListOptions{})

	if err != nil {
		fmt.Printf("Error while fetching all Deployments. Error: %v", err)
	}

	if podState == "" {
		c.Pods = pods.Items

	} else if podState == "healthy" {
		for _, i := range pods.Items {
			if len(i.Status.ContainerStatuses) != 0 {
				if i.Status.ContainerStatuses[0].State.Running != nil {
					c.Pods = append(c.Pods, i)
				}
			}
		}

	} else if podState == "unhealthy" {
		for _, i := range pods.Items {
			if len(i.Status.ContainerStatuses) != 0 {
				if i.Status.ContainerStatuses[0].State.Running == nil {
					c.Pods = append(c.Pods, i)
				}
			}
		}
	}

}
