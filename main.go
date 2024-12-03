package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const backendURL = "http://test-back.iot-edge.svc.cluster.local/data"

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Data endpoint
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Calling backend at: %s", backendURL)

		// Call backend
		resp, err := http.Get(backendURL)
		if err != nil {
			log.Printf("Detailed error: %+v", err)
			http.Error(w, fmt.Sprintf("Failed to fetch data from backend: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Log response status code
		log.Printf("Backend response status: %d", resp.StatusCode)

		if resp.StatusCode != http.StatusOK {
			log.Printf("Backend returned non-200 status: %d", resp.StatusCode)
			http.Error(w, fmt.Sprintf("Backend returned status: %d", resp.StatusCode), resp.StatusCode)
			return
		}

		// Read response data
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read backend response", http.StatusInternalServerError)
			log.Printf("Error reading backend response: %v", err)
			return
		}

		// Set JSON content type
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	// Start frontend server
	port := "8081"
	log.Printf("Frontend server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
