package api

import (
	"fmt"
	"net/http"

<<<<<<< HEAD
	utils "github.com/kk2526/k8s-central/pkg/utils"
=======
	***REMOVED*** "github.com/kk2526/k8s-central/pkg/utils"
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
)

func ObjectsAPIHandler(w http.ResponseWriter, r *http.Request) {

	var (
<<<<<<< HEAD
		objResponseData utils.K8sObj
=======
		objResponseData ***REMOVED***.K8sObj
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
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
<<<<<<< HEAD
		resp.Data = utils.GetNamespaces(client)
=======
		resp.Data = ***REMOVED***.GetNamespaces(client)
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
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

<<<<<<< HEAD
	respJSON, _ := utils.EncodeJSON(resp)
=======
	respJSON, _ := ***REMOVED***.EncodeJSON(resp)
>>>>>>> 493881d5b0235aa47da0912003042c00c4526d6e
	fmt.Fprint(w, respJSON)
}
