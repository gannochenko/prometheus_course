package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_app_request_count",
		Help: "Total APP HTTP Request count",
	})
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// https://stackoverflow.com/questions/65608610/how-to-use-gin-as-a-server-to-write-prometheus-exporter-metrics
func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		RequestCount.Inc()

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/metrics", prometheusHandler())
	r.Run()
}
