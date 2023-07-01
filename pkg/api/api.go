package api

import (
	"net/http"

	utils "github.com/kube-stale/pkg/utils"
)

func init() {
	client = utils.CreateK8sClient()

}

func Run() {
	http.HandleFunc("/api/health/up", HealthAPIHandler)
	http.HandleFunc("/api/cluster/objects", ObjectsAPIHandler)
}
