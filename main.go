package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	addr := flag.String("addr", "127.0.0.1", "bind address")
	port := flag.String("port", "9000", "bind port")
	dir := flag.String("dir", "./", "web server root directory")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.Printf("Exiting...")
		os.Exit(1)
	}()

	log.Printf("Serving %s at: http://%s:%s\n", *dir, *addr, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *addr, *port), nil))
}
