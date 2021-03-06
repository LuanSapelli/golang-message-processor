package controllers

import (
	"fmt"
	"net/http"
)

func HealthCheck() {
	http.HandleFunc("/message-processor/health-check", HealthCheckServer)
	_ = http.ListenAndServe(":80", nil)
}

func HealthCheckServer(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "alive and kicking")
}