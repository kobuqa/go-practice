package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kobuqa/hello-world-service/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	serverAddr := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	log.Printf("Server starting on %s in %s mode...", serverAddr, cfg.Env)

	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatal(err)
	}
}
