package api

import (
	"fmt"
	"net/http"
)

func HealthAPIHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "200 OK")

}
