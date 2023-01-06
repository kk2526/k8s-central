package main

import (
	"log"
	"os"

	api "***REMOVED***/ecp-cluster-tool/pkg/api"
	controller "***REMOVED***/ecp-cluster-tool/pkg/controller"
	"***REMOVED***/ecp-cluster-tool/pkg/utils"
)

var profile string

func init() {
	profile, _ = os.LookupEnv("ECP_APP_PROFILE")
	if profile == "daemonset" || profile == "ui" || profile == "api" || profile == "controller" {
		log.Println("ECP_APP_PROFILE ENV VAR IS SET")
	} else {
		log.Println("ECP_APP_PROFILE ENV VAR IS NOT SET, defaulting to dev")
		profile = "dev"
	}

}

func main() {

	log.Printf("Starting Application")
	log.Printf("Profile Selected: %s", profile)

	switch profile {

	case "api":
		api.Run()
		utils.ListenServe()

	case "controller":
		controller.Run()

	case "dev":
		api.Run()
		utils.ListenServe()

	default:
		log.Printf("No profile chosen")
	}
}
