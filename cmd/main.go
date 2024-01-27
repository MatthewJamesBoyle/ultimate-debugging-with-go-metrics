package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	// basic counter - A cumulative metric that represents a single numerical value that only ever goes up.
	transactionSuccess = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "transaction_success",
		Help: "Successful transactions",
	})

	// Gauge: A metric that represents a single numerical value that can arbitrarily go up and down.
	BloodSugarGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "matts_blood_sugar",
		Help: "Current blood Sugar",
	})

	// Histogram: A metric that samples observations (usually things like request durations or response sizes) and
	//counts them in configurable buckets. It also provides a sum of all observed values.
	httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duration of HTTP requests in seconds",
		// Define suitable buckets for measuring request duration.
		Buckets: []float64{0.001, 0.003, 0.01, 0.03, 0.1, 0.3, 1, 3, 10},
	}, []string{"path"})

	// Summary: Similar to a histogram, a summary samples observations. It calculates configurable quantiles over a sliding time window.
	httpResponseSize = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "http_response_size_bytes",
		Help:       "Size of HTTP responses in bytes",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
	}, []string{"path"})
)

func main() {

	if err := prometheus.Register(transactionSuccess); err != nil {
		log.Fatal(err)
	}
	if err := prometheus.Register(BloodSugarGauge); err != nil {
		log.Fatal(err)
	}
	if err := prometheus.Register(httpDuration); err != nil {
		log.Fatal(err)
	}
	if err := prometheus.Register(httpResponseSize); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/tx", func(w http.ResponseWriter, r *http.Request) {
		transactionSuccess.Inc()

		w.Write([]byte("counter incremented"))
	})

	http.HandleFunc("/blood", func(w http.ResponseWriter, r *http.Request) {

		BloodSugarGauge.Set(30 + 270*rand.Float64())

		w.Write([]byte("counter incremented"))
	})

	http.HandleFunc("/duration", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))

		duration := time.Since(startTime).Seconds()

		httpDuration.WithLabelValues("/duration").Observe(duration)

		w.Write([]byte("counter incremented"))
	})

	http.HandleFunc("/size", func(w http.ResponseWriter, r *http.Request) {
		responseSize := rand.Intn(2000) // Random response size between 0 and 2000 bytes

		fmt.Fprintf(w, "This is a response of random size: %d bytes", responseSize)

		httpResponseSize.WithLabelValues("/size").Observe(float64(responseSize))
	})

	// Expose the registered Prometheus metrics at the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())

	// Start the server.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
