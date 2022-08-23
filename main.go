package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}

func state() {
	fmt.Println("Started Golang Webserver")
}

func main() {
	go func() {
		state()
	}()

	fileName := "./logs/log.txt"
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	log.SetOutput(logFile)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public_html")))

	log.Fatal(http.ListenAndServe(":8888", RequestLogger(mux)))

}
