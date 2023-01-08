package main

import (
	"log"
	"os"

	api "github.com/kk2526/k8s-central/pkg/api"
	controller "github.com/kk2526/k8s-central/pkg/controller"
	"github.com/kk2526/k8s-central/pkg/utils"
)

var profile string

func init() {
	profile, _ = os.LookupEnv("APP_PROFILE")
	if profile == "ui" || profile == "api" || profile == "controller" {
		log.Println("APP_PROFILE ENV VAR IS SET")
	} else {
		log.Println("APP_PROFILE ENV VAR IS NOT SET, defaulting to api")
		profile = "api"
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
		controller.Run()
	default:
		log.Printf("No profile chosen")
	}
}
