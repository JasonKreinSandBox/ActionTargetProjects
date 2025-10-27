package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"
)

var (
	hosts    *string
	port     *int
	interval *int
	counter  *int
	avgPing  *[]time.Duration
)

func init() {
	hosts = flag.String("hosts", "localhost", "Comma-separated list of hostnames")
	port = flag.Int("port", 8080, "Port number")
	interval = flag.Int("interval", 10, "Interval in seconds")
	counter = flag.Int("counter", 0, "Ping counter")
	avgPing = &[]time.Duration{}
}

func main() {
	flag.Parse()
	http.HandleFunc("/events", sseHandler)
	fmt.Printf("Starting server on port :8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}

func PingHost(host string) (time.Duration, error) {
	start := time.Now()
	_, err := net.LookupHost(host)
	if err != nil {
		return 0, err
	}
	return time.Since(start), nil
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	clientGone := r.Context().Done()

	rc := http.NewResponseController(w)
	t := time.NewTicker(time.Duration(*interval) * time.Second)
	defer t.Stop()

	for {
		select {
		case <-clientGone:
			fmt.Println("Client disconnected")
			return
		case <-t.C:
			pingTime, err := PingHost(*hosts)
			if err != nil {
				fmt.Fprintf(w, "data: Error: %v\n\n", err)
				return
			}
			fmt.Fprintf(w, "data: Ping time to %s is %v\n\n", *hosts, pingTime)
			fmt.Fprintf(w, "data: Interval: %d seconds\n\n", *interval)
			*counter++
			fmt.Fprintf(w, "data: Ping count: %d\n\n", *counter)
			*avgPing = append(*avgPing, pingTime)
			avg := avgPingTotal(*avgPing)
			fmt.Fprintf(w, "data: Avg response time: %v\n\n", avg)
		}
		rc.Flush()
	}
}

func avgPingTotal(durations []time.Duration) time.Duration {
	var total time.Duration
	for _, d := range durations {
		total += d
	}
	return total / time.Duration(len(durations))
}
