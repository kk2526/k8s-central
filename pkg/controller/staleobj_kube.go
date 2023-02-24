package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type K8sObj struct {
	Deployments            []appsv1.Deployment        `json:"Deployments"`
	Secrets                []v1.Secret                `json:"Secrets"`
	Configmaps             []v1.ConfigMap             `json:"ConfigMaps"`
	Services               []v1.Service               `json:"Services"`
	Endpoints              []v1.Endpoints             `json:"Endpoints"`
	PersistentVolumeClaims []v1.PersistentVolumeClaim `json:"PVCs"`
	Pods                   []v1.Pod                   `json:"Pods"`
	Ingress                []networking.Ingress       `json:"Ingress"`
}

// Struct internally used for comparision
type K8sObjMeta struct {
	apiobject string // Name of K8s Object
	namespace string
}

type staleObjData struct {
	LastCheckTime string `json:"Last-Check-Time"`
	ObjStatus     string `json:"Status"`
}

var (
	excludens, _ = regexp.Compile("kube.*|cattle.*|ingress-nginx")
)

func (ob *K8sObj) FetchUsedObjs() {
	log.Printf("Getting All Used Objects\n")

	log.Printf("Getting Used Endpoints..")
	eps, err := client.CoreV1().Endpoints("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error loading Endpoints %v", err)
	}
	for _, ep := range eps.Items {
		if !excludens.MatchString(ep.Namespace) {
			if ep.Subsets != nil {
				ob.Endpoints = append(ob.Endpoints, ep)
			}
		}
	}

	pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error loading pods %v", err)
	}
	log.Printf("Iterating Pods For Used Storage Objects")

	// Function Literal to Fetch Used Secrets and ConfigMap
	usedSecretConfigMap := func(item v1.Container, ns string) {
		if item.Env != nil {
			for _, env := range item.Env {
				if env.ValueFrom != nil {
					if env.ValueFrom.SecretKeyRef != nil {
						fetchsecret, _ := client.CoreV1().Secrets(ns).Get(context.TODO(), env.ValueFrom.SecretKeyRef.LocalObjectReference.Name, metav1.GetOptions{})
						ob.Secrets = append(ob.Secrets, *fetchsecret)
					} else if env.ValueFrom.ConfigMapKeyRef != nil {
						fetchconfigmap, _ := client.CoreV1().ConfigMaps(ns).Get(context.TODO(), env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name, metav1.GetOptions{})
						ob.Configmaps = append(ob.Configmaps, *fetchconfigmap)
					}
				}

			}
		}
		if item.EnvFrom != nil {
			for _, envfrom := range item.EnvFrom {
				if envfrom.SecretRef != nil {
					fetchsecret, _ := client.CoreV1().Secrets(ns).Get(context.TODO(), envfrom.SecretRef.LocalObjectReference.Name, metav1.GetOptions{})
					ob.Secrets = append(ob.Secrets, *fetchsecret)

				} else if envfrom.ConfigMapRef != nil {
					fetchconfigmap, _ := client.CoreV1().ConfigMaps(ns).Get(context.TODO(), envfrom.ConfigMapRef.LocalObjectReference.Name, metav1.GetOptions{})
					if fetchconfigmap.Name != "" {
						ob.Configmaps = append(ob.Configmaps, *fetchconfigmap)
					}
				}

			}
		}
	}

	for _, i := range pods.Items {
		if !excludens.MatchString(i.Namespace) {
			container := i.Spec.Containers
			initcontainer := i.Spec.InitContainers
			for _, item := range container {
				log.Printf("Fetching Used Secrets & ConfigMaps")
				usedSecretConfigMap(item, i.Namespace)
			}
			for _, item := range initcontainer {
				usedSecretConfigMap(item, i.Namespace)
			}

			if i.Spec.Volumes != nil {
				for _, volume := range i.Spec.Volumes {
					log.Printf("Fetching Used PVCs")
					if volume.VolumeSource.Secret != nil {
						fetchsecret, _ := client.CoreV1().Secrets(i.Namespace).Get(context.TODO(), volume.VolumeSource.Secret.SecretName, metav1.GetOptions{})
						ob.Secrets = append(ob.Secrets, *fetchsecret)
					} else if volume.VolumeSource.ConfigMap != nil {
						fetchconfigmap, _ := client.CoreV1().ConfigMaps(i.Namespace).Get(context.TODO(), volume.VolumeSource.ConfigMap.Name, metav1.GetOptions{})
						ob.Configmaps = append(ob.Configmaps, *fetchconfigmap)
					} else if volume.VolumeSource.PersistentVolumeClaim != nil {
						fetchpvc, _ := client.CoreV1().PersistentVolumeClaims(i.Namespace).Get(context.TODO(), volume.VolumeSource.PersistentVolumeClaim.ClaimName, metav1.GetOptions{})
						ob.PersistentVolumeClaims = append(ob.PersistentVolumeClaims, *fetchpvc)
					}

				}
			}
		}
	}
}

// Takes Struct Of Used K8s Objects and returns Struct of UN-Used K8s Objects
func GetUnusedObjs(used K8sObj) K8sObj {
	var staleObjs K8sObj
	//Get UnUsed ConfigMaps
	allconfigs, err := client.CoreV1().ConfigMaps("").List(context.TODO(), metav1.ListOptions{FieldSelector: "metadata.name!=kube-root-ca.crt"})

	if err != nil {
		log.Printf("Error retrieving all defined Configmaps: %v", err)
	}
	log.Println("Calculating all Un-Used Configmaps")
	for _, i := range allconfigs.Items {
		var match bool
		if !excludens.MatchString(i.Namespace) {
			allcm := K8sObjMeta{i.Name, i.Namespace}
			for _, j := range used.Configmaps {
				usedcm := K8sObjMeta{j.Name, j.Namespace}
				if allcm == usedcm {
					match = true
					_ = match
				}
			}
			if !match {
				staleObjs.Configmaps = append(staleObjs.Configmaps, i)
			}
		}
	}

	//Get UnUsed Services
	allsvcs, err := client.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Printf("Error retrieving all defined Services: %v", err)
	}
	log.Println("Calculating all Un-Used Services")
	for _, i := range allsvcs.Items {
		var match bool
		if !excludens.MatchString(i.Namespace) {
			allsvc := K8sObjMeta{i.Name, i.Namespace}
			for _, j := range used.Endpoints {
				usedsvc := K8sObjMeta{j.Name, j.Namespace}
				if allsvc == usedsvc {
					match = true
					_ = match
				}
			}
			if !match {
				staleObjs.Services = append(staleObjs.Services, i)
			}
		}

	}

	//Get UnUsed Ingresses
	ings, err := client.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error loading Ingress %v", err)
	}
	for _, i := range ings.Items {
		if i.Spec.Rules != nil {
			for _, rule := range i.Spec.Rules {
				m := map[string]bool{}
				for _, path := range rule.HTTP.Paths {
					var match bool
					if _, ok := m[path.Backend.Service.Name]; !ok { //Make sure unique service name within Path are being compared
						m[path.Backend.Service.Name] = true
						svcinIng := K8sObjMeta{path.Backend.Service.Name, i.Namespace}
						for _, j := range used.Endpoints {
							usedsvc := K8sObjMeta{j.Name, j.Namespace}
							if svcinIng == usedsvc {
								match = true
								_ = match
							}
						}
						if !match {
							staleObjs.Ingress = append(staleObjs.Ingress, i)
						}
					}
				}
			}
		}
	}

	//Get UnUsed Secrets
	allsecrets, err := client.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Printf("Error retrieving all defined Secrets: %v", err)
	}
	log.Println("Calculating all Un-Used Secrets")
	for _, i := range allsecrets.Items {
		var match bool
		if !excludens.MatchString(i.Namespace) {
			if i.Type != "kubernetes.io/service-account-token" && i.Type != "kubernetes.io/tls" {
				allsecret := K8sObjMeta{i.Name, i.Namespace}
				for _, j := range used.Secrets {
					usedsecret := K8sObjMeta{j.Name, j.Namespace}
					if allsecret == usedsecret {
						match = true
						_ = match
					}
				}
				if !match {
					staleObjs.Secrets = append(staleObjs.Secrets, i)
				}
			}
		}
	}

	//Get UnUsed PVCs
	allpvc, err := client.CoreV1().PersistentVolumeClaims("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Printf("Error retrieving all defined PersistentVolumeClaims: %v", err)
	}
	log.Println("Calculating all Un-Used PVCs")
	for _, i := range allpvc.Items {
		var match bool
		if !excludens.MatchString(i.Namespace) {
			allpvc := K8sObjMeta{i.Name, i.Namespace}
			for _, j := range used.PersistentVolumeClaims {
				usedpvc := K8sObjMeta{j.Name, j.Namespace}
				if allpvc == usedpvc {
					match = true
					_ = match
				}
			}
			if !match {
				staleObjs.PersistentVolumeClaims = append(staleObjs.PersistentVolumeClaims, i)
			}
		}
	}
	return staleObjs
}

// Builds annotation json value struct that will be applied to stale objects
func (n *staleObjData) BuildAnnotationData() {
	overrideStatus, _ := os.LookupEnv("OVERRIDE_STATUS")
	if overrideStatus == "" {
		n.ObjStatus = "NotInUse"

	} else {
		n.ObjStatus = overrideStatus
	}
	n.LastCheckTime = time.Now().Format(time.RFC3339Nano)
}

// Set stale annotation to an Object
func SetCustomAnnotate(object *metav1.ObjectMeta, data *staleObjData) {
	annotations := object.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
		jsonStr, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
<<<<<<< HEAD
		annotations["kubernetes.io/stale-object"] = string(jsonStr)
=======
		annotations["***REMOVED***/stale-object"] = string(jsonStr)
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
	} else {
		jsonStr, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
<<<<<<< HEAD
		annotations["kubernetes.io/stale-object"] = string(jsonStr)
=======
		annotations["***REMOVED***/stale-object"] = string(jsonStr)
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
	}
	object.SetAnnotations(annotations)
}

// Annotate using above struct
func (n *K8sObj) AnnotateStaleObj(m *staleObjData) {
	// Annotate ConfigMap
	cms := n.Configmaps
	for _, i := range cms {
		SetCustomAnnotate(&i.ObjectMeta, m)
		client.CoreV1().ConfigMaps(i.Namespace).Update(context.TODO(), &i, metav1.UpdateOptions{})
		log.Printf("Updated Annotation on CM: %v", i.Name)
	}

	// Annotate Services
	svc := n.Services
	for _, i := range svc {
		SetCustomAnnotate(&i.ObjectMeta, m)
		client.CoreV1().Services(i.Namespace).Update(context.TODO(), &i, metav1.UpdateOptions{})
		log.Printf("Updated Annotation on Service: %v", i.Name)
	}

	// Annotate Secrets
	secret := n.Secrets
	for _, i := range secret {
		SetCustomAnnotate(&i.ObjectMeta, m)
		client.CoreV1().Secrets(i.Namespace).Update(context.TODO(), &i, metav1.UpdateOptions{})
		log.Printf("Updated Annotation on Secret: %v", i.Name)
	}

	// Annotate PVC
	pvc := n.PersistentVolumeClaims
	for _, i := range pvc {
		SetCustomAnnotate(&i.ObjectMeta, m)
		client.CoreV1().PersistentVolumeClaims(i.Namespace).Update(context.TODO(), &i, metav1.UpdateOptions{})
		log.Printf("Updated Annotation on PVC: %v", i.Name)
	}

	// Annotate Ingress
	ingress := n.Ingress
	for _, i := range ingress {
		SetCustomAnnotate(&i.ObjectMeta, m)
		client.NetworkingV1().Ingresses(i.Namespace).Update(context.TODO(), &i, metav1.UpdateOptions{})
		log.Printf("Updated Annotation on Ingress: %v", i.Name)
	}
}

// Main Function
func SetUnusedObj() K8sObj {
	var usedObj = K8sObj{}

	usedObj.FetchUsedObjs()
	c := GetUnusedObjs(usedObj)

	return c
}
