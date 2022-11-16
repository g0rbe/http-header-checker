package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {

	listenAddress := flag.String("address", ":8080", "Listen address")

	flag.Parse()

	if *listenAddress == "" {
		fmt.Fprintf(os.Stderr, "address is empty!\n")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("New request from %s\n", r.RemoteAddr)

		for h := range r.Header {
			fmt.Printf("	%s -> %v\n", h, r.Header[h])
		}

		fmt.Printf("---------------\n")
	})

	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP Server: %s\n", err)
	}
}
