package api

import (
	"net/http"

	***REMOVED*** "github.com/kk2526/k8s-central/pkg/utils"
)

func init() {
	client = ***REMOVED***.CreateK8sClient()

}

func Run() {
	http.HandleFunc("/api/health/up", HealthAPIHandler)
	http.HandleFunc("/api/cluster/objects", ObjectsAPIHandler)
}
