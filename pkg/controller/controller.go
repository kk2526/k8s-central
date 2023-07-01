package controller

import (
	"log"
	"time"

	utils "github.com/kube-stale/pkg/utils"

	clientset "k8s.io/client-go/kubernetes"
)

var client *clientset.Clientset

func init() {
	client = utils.CreateK8sClient()
}

func Run() {
	log.Printf("Running Controller")
	dur := 1 * time.Hour
	ticker := time.NewTicker(dur)

	for _ = range ticker.C {
		d := staleObjData{}
		c := SetUnusedObj()

		d.BuildAnnotationData()
		c.AnnotateStaleObj(&d)
	}

}
