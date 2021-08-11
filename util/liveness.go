package util

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("i'm alive! :)\n"))
}

func RunLivenessProbe(port string) {
	log.Printf("[INFO] Liveness probe up on :%s/health\n", port)
	http.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalln("[ERROR] Live- and readiness probe failed")
	}
}
