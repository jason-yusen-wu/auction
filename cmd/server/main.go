package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// renders a blank screen -> expected behavior!
}

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func main() {
	r := chi.NewRouter()

	// chi middleware
	r.Use(metricsMiddleware)

	r.Get("/healthz", healthzHandler)
	r.Handle("/metrics", promhttp.Handler())

	if ok := http.ListenAndServe(":3000", r); ok != nil {
		os.Exit(-1)
	}
}
