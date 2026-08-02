package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// renders a blank screen -> expected behavior!
}

func main() {
	r := chi.NewRouter()

	r.Get("/healthz", healthzHandler)

	if ok := http.ListenAndServe(":3000", r); ok != nil {
		fmt.Fprintln(os.Stderr, "Error starting server: %v", ok.Error())
		os.Exit(-1)
	}
}
