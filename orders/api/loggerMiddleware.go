package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func LoggerMiddleware(_ http.ResponseWriter, r *http.Request) {
	fmt.Printf("Path:\t%s\n", r.URL.Path)
	fmt.Printf("Method:\t%s\n", r.Method)

	if r.Method != http.MethodGet {
		body, _ := ioutil.ReadAll(r.Body)
		bodyString := string(body)
		fmt.Printf("Body:\t%s\n", bodyString)
	}
}
