package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/halworsen/grug"
)

var grugSession *grug.GrugSession

func main() {
	log.Println("[INFO] Starting Grug session")
	grugSession = &grug.GrugSession{}
	grugSession.New("grug.yaml")

	log.Println("[INFO] Grug is now running (CTRL-C to exit)")
	// wait for a termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc

	grugSession.Close()
}
