package api

import (
	"net/http"

	***REMOVED*** "***REMOVED***/ecp-cluster-tool/pkg/utils"
)

func init() {
	client = ***REMOVED***.CreateK8sClient()

}

func Run() {
	http.HandleFunc("/api/health/up", HealthAPIHandler)
	http.HandleFunc("/api/cluster/objects", ObjectsAPIHandler)
}
