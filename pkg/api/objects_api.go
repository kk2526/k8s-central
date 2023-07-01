package api

import (
	"fmt"
	"net/http"

	utils "github.com/kube-stale/pkg/utils"
)

func ObjectsAPIHandler(w http.ResponseWriter, r *http.Request) {

	var (
		objResponseData utils.K8sObj
		resp            APIResponse
	)

	reqParam := APIRequestParams{}
	reqParam.Parse(r)
	switch reqParam.objectType {
	case "configmaps":
		objResponseData.GetConfigMaps(client, reqParam.namespace, reqParam.stale)
	case "secrets":
		objResponseData.GetSecrets(client, reqParam.namespace, reqParam.stale)
	case "pvcs":
		objResponseData.GetPVC(client, reqParam.namespace, reqParam.stale)
	case "services":
		objResponseData.GetServices(client, reqParam.namespace, reqParam.stale)
	case "deployments":
		objResponseData.GetDeployment(client, reqParam.namespace, reqParam.stale)
	case "pods":
		objResponseData.GetPods(client, reqParam.namespace, reqParam.podState)
	case "namespaces":
		resp.Data = utils.GetNamespaces(client)
	default:
		fallthrough
	case "all":
		objResponseData.GetServices(client, reqParam.namespace, reqParam.stale)
		objResponseData.GetConfigMaps(client, reqParam.namespace, reqParam.stale)
		objResponseData.GetSecrets(client, reqParam.namespace, reqParam.stale)
		objResponseData.GetPVC(client, reqParam.namespace, reqParam.stale)
		objResponseData.GetDeployment(client, reqParam.namespace, reqParam.stale)
		objResponseData.GetPods(client, reqParam.namespace, reqParam.podState)
	}

	if reqParam.objectType != "namespaces" {
		resp.Data = objResponseData
	}

	respJSON, _ := utils.EncodeJSON(resp)
	fmt.Fprint(w, respJSON)
}
