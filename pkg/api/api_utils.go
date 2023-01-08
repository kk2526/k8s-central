package api

import (
	"fmt"
	"net/http"
	"strconv"
)

func (rp *APIRequestParams) Parse(r *http.Request) {
	var err error
	r.ParseForm()
	rp.namespace = r.FormValue("namespace")
	rp.objectType = r.FormValue("type")
	rp.podState = r.FormValue("podstate")
	if r.FormValue("stale") != "" {
		rp.stale, err = strconv.ParseBool(r.FormValue("stale"))
		if err != nil {
			fmt.Printf("Error parsing boolean object: %v. Error: %v", rp.stale, err)
		}
	}
}
