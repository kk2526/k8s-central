package api

import (
	clientset "k8s.io/client-go/kubernetes"
)

var client *clientset.Clientset

type APIRequestParams struct {
	namespace  string
	objectType string
	podState   string
	stale      bool
}

type APIResponse struct {
	Errors []string    `json:"errors"`
	Data   interface{} `json:"data"`
}
