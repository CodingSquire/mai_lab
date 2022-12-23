package api

import (
	"github.com/innerave/mai_lab/orders/internal/http"
	"log"
	"time"
)

func LoggerMiddleware(c *http.RouteContext) {
	log.Printf("Path:\t%s\n", c.R.URL.Path)
	log.Printf("Method:\t%s\n", c.R.Method)

	start := time.Now()
	c.Next()

	log.Printf("Time spent: %d\n", time.Since(start))
}
