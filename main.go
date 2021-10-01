package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func prometheusHandler() gin.HandlerFunc {
	handler := promhttp.Handler()

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	engine := gin.Default()
	engine.Static("/", "/var/www")
	go metricsMain()
	err := engine.Run()
	if err != nil {
		log.Fatalf("Gin ran into an error: %v", err)
	}
}

func metricsMain() {
	engine := gin.New()
	engine.GET("/metrics", prometheusHandler())
	engine.Run(":8081")
}
