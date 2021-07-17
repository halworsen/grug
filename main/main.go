package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/halworsen/grug"
)

var grugSession *grug.GrugSession

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("i'm alive! :)\n"))
}

func livenessProbe(liveChan chan string) {
	liveChan <- "[INFO] Liveness probe up on :5700/health"
	http.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(":5700", nil)
	if err != nil {
		log.Fatalln("[ERROR] Live- and readiness probe failed")
	}
}

func main() {
	log.Println("[INFO] Starting Grug session")
	grugSession = &grug.GrugSession{}
	grugSession.New("/etc/grug/grug.yaml")

	log.Println("[INFO] Starting liveness probe...")
	liveChan := make(chan string)
	go livenessProbe(liveChan)
	log.Println(<-liveChan)

	log.Println("[INFO] Grug is now running (CTRL-C to exit)")
	// wait for a termination signal
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	grugSession.Close()
}
