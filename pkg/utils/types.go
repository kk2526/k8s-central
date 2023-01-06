package utils

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

type K8sObj struct {
	Deployments            []appsv1.Deployment        `json:"Deployments"`
	Secrets                []v1.Secret                `json:"Secrets"`
	Configmaps             []v1.ConfigMap             `json:"ConfigMaps"`
	Services               []v1.Service               `json:"Services"`
	Endpoints              []v1.Endpoints             `json:"Endpoints"`
	PersistentVolumeClaims []v1.PersistentVolumeClaim `json:"PVCs"`
	Pods                   []v1.Pod                   `json:"Pods"`
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
