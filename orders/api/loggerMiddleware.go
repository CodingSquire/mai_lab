package api

import (
	"fmt"
	"io/ioutil"
	"orders/http"
	"time"
)

func LoggerMiddleware(c *http.RouteContext) {
	body, _ := ioutil.ReadAll(c.Body())
	bodyString := string(body)

	fmt.Printf("Path:\t%s\n", c.R.URL.Path)
	fmt.Printf("Method:\t%s\n", c.R.Method)
	fmt.Printf("Body:\t%s\n", bodyString)

	start := time.Now()
	c.Next()

	fmt.Printf("Time spent: %d\n", time.Since(start))
}
