package api

import (
	"fmt"
	"orders/internal/http"
	"time"
)

func LoggerMiddleware(c *http.RouteContext) {
	fmt.Printf("Path:\t%s\n", c.R.URL.Path)
	fmt.Printf("Method:\t%s\n", c.R.Method)

	start := time.Now()
	c.Next()

	fmt.Printf("Time spent: %d\n", time.Since(start))
}
