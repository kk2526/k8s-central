package controller

import (
	"log"
	"time"

	***REMOVED*** "github.com/kk2526/k8s-central/pkg/utils"
	clientset "k8s.io/client-go/kubernetes"
)

var client *clientset.Clientset

func init() {
	client = ***REMOVED***.CreateK8sClient()

}

func Run() {
	log.Printf("Running Controller")
	dur := 1 * time.Hour
	ticker := time.NewTicker(dur)

	for _ = range ticker.C {
		d := staleObjData{}
		c := SetUnusedObj()

		d.CreateAnnotationData()
		c.AnnotateStaleObj(&d)
	}

}
