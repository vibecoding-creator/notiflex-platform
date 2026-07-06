package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	counter  int
	counterM sync.Mutex
	podName  string
)

func init() {
	podName = os.Getenv("HOSTNAME")
	if podName == "" {
		podName = "unknown"
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ok","timestamp":"%s"}`, time.Now().UTC().Format(time.RFC3339))
}

func id(w http.ResponseWriter, r *http.Request) {
	counterM.Lock()
	defer counterM.Unlock()

	counter++
	currentID := counter

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"id":%d,"pod":"%s","timestamp":"%s"}`, currentID, podName, time.Now().UTC().Format(time.RFC3339))
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/id", id)

	log.Println("Notiflex API server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
