package app

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

var (
	listenAddr string
)

func StartApplication() {
	flag.StringVar(&listenAddr, "listen-addr", ":8080", "server listen address")
	flag.Parse()

	router := chi.NewRouter()
	mapUrls(router)

	server := http.Server{
		Addr:           listenAddr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Running server...")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
