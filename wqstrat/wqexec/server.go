package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	api "wquant/back/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	server := api.Engine("test")
	defer server.Shutdown()

	// kis
	kis := server.Conn.Group("/api/v1/kis")
	server.MountServiceKIS(kis)

	// binance
	bnc := server.Conn.Group("/api/v1/bn")
	server.MountServiceBinance(bnc)

	// Host webserver
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	log.Printf("Hosting on %s:%s", host, port)
	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", host, port),
		Handler:        server.Conn,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe())
}
