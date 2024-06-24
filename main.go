package main

import (
	"net/http"
	"time"
	"fmt"
)

func main() {
	srv := http.Server{
		Addr:              fmt.Sprintf("%s:%d", "0.0.0.0", "8080"),
		ReadHeaderTimeout: 5 * time.Second,
	}
	http.HandleFunc("/", p.handler)

	if err := srv.ListenAndServe(); err != nil {
		p.log.Fatalf("Webhook listener: %v", err)
	}
}