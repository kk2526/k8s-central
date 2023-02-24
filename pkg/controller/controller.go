package controller

import (
	"log"
	"time"

<<<<<<< HEAD
	utils "github.com/kk2526/k8s-central/pkg/utils"
=======
	***REMOVED*** "github.com/kk2526/k8s-central/pkg/utils"
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
	clientset "k8s.io/client-go/kubernetes"
)

var client *clientset.Clientset

func init() {
<<<<<<< HEAD
	client = utils.CreateK8sClient()
=======
	client = ***REMOVED***.CreateK8sClient()
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e

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
