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
)

func init() {
	hosts = flag.String("hosts", "localhost", "Comma-separated list of hostnames")
	port = flag.Int("port", 80, "Port number")
	interval = flag.Int("interval", 10, "Interval in seconds")
}

func main() {
	flag.Parse()
	counter = new(int)
	*counter = 0
	tk := time.NewTicker(time.Duration(*interval) * time.Second)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for ; ; <-tk.C {
			time, err := PingHost(*hosts, *port, *interval)
			if err != nil {
				fmt.Fprintf(w, "Error: %v\n", err)
				return
			}
			fmt.Fprintf(w, "Host is reachable. Ping time to %s:%d is %v\n", *hosts, *port, time)
			fmt.Fprintf(w, "Ping count: %d\n", *counter)
			fmt.Fprintf(w, "Interval: %d seconds\n", *interval)
			*counter++
		}
	})

	http.ListenAndServe(":80", nil)

}

func PingHost(host string, port int, interval int) (time.Duration, error) {
	start := time.Now()
	_, err := net.LookupHost(host)
	if err != nil {
		return 0, err
	}
	return time.Since(start), nil
}
