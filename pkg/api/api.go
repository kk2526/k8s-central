package api

import (
	"net/http"

<<<<<<< HEAD
	utils "github.com/kk2526/k8s-central/pkg/utils"
)

func init() {
	client = utils.CreateK8sClient()
=======
	utils "github.com/kk2526/k8s-central/pkg/utils"
)

func init() {
	client = utils.CreateK8sClient()
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e

}

func Run() {
	http.HandleFunc("/api/health/up", HealthAPIHandler)
	http.HandleFunc("/api/cluster/objects", ObjectsAPIHandler)
}
