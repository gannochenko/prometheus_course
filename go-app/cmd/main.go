package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// RequestCount is a sample counter metrics
	RequestCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_app_request_count",
		Help: "Total APP HTTP Request count",
	})
	// RequestInProgress is a sample gauge metrics
	RequestInProgress = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_app_request_in_progress",
		Help: "HTTP Request in progress",
	})
	// RequestResponseTime is a sample summary metrics
	// promql average latency:
	RequestResponseTime = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: "go_app_response_latency",
		Help: "Response latency",
	}, []string{"path"})
	// RequestResponseTimeHistogram is a sample histogram metrics
	RequestResponseTimeHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "go_app_response_latency_histogram",
		Help: "Response latency (histogram)",
	}, []string{"path"})
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func latencySampler() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		RequestResponseTime.WithLabelValues(c.Request.RequestURI).Observe(time.Since(startTime).Seconds())
		RequestResponseTimeHistogram.WithLabelValues(c.Request.RequestURI).Observe(time.Since(startTime).Seconds())
	}
}

// https://stackoverflow.com/questions/65608610/how-to-use-gin-as-a-server-to-write-prometheus-exporter-metrics
func main() {
	r := gin.Default()
	r.Use(latencySampler())

	r.GET("/ping", func(c *gin.Context) {
		RequestInProgress.Inc()

		time.Sleep(2 * time.Second)
		RequestCount.Inc()

		RequestInProgress.Dec()
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/metrics", prometheusHandler())
	r.Run()
}
