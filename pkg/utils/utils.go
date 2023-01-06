package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// func TestDNSHost(dnsIP string, url string) (bool, error) {
// 	const timeout = 3000 * time.Millisecond
// 	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
// 	defer cancel() // important to avoid a resource leak
// 	log.Printf("Setting custom resolver for %v", dnsIP)
// 	r := &net.Resolver{
// 		PreferGo: true,
// 		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
// 			d := net.Dialer{
// 				Timeout: timeout,
// 			}
// 			return d.DialContext(ctx, network, dnsIP+":53")
// 		},
// 	}
// 	log.Printf("Making request to %v", url)
// 	ip, err := r.LookupHost(ctx, url)
// 	if err != nil {
// 		return false, err
// 	}
// 	log.Printf("Response %v", ip)
// 	return true, nil
// }

// func TestConn(url string) (bool, error) {
// 	const timeout = 3000 * time.Millisecond
// 	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
// 	defer cancel() // important to avoid a resource leak
// 	log.Printf("Calling endpoint %v", url)
// 	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		if resp != nil {
// 			log.Printf("Endpoint: %v, Status code: %v", url, resp.StatusCode)
// 			return false, err
// 		} else {
// 			log.Printf("Endpoint: %v, Status code: %v", url, "UNKNOWN")
// 			return false, err
// 		}
// 	}

// 	log.Printf("Endpoint: %v, Status code: %v", url, resp.StatusCode)
// 	return true, nil

// }
func ListenServe() {
	log.Println("Listing for requests at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func EncodeJSON(data any) (string, error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}
