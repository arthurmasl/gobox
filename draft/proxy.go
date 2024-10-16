package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("Starting proxy server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())

	if r.Method == http.MethodConnect {
		handleConnect(w, r)
		return
	}

	handleHTTP(w, r)
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	targetURL := r.URL.String()

	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	req.Header = r.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error forwarding request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Error copying response body", http.StatusInternalServerError)
	}
}

func handleConnect(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))

	conn, err := net.Dial("tcp", r.URL.Host)
	if err != nil {
		log.Printf("Error connecting to target server: %v", err)
		return
	}
	defer conn.Close()

	go io.Copy(conn, r.Body)
	io.Copy(w, conn)
}
