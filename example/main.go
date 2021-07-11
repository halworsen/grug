package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/halworsen/grug"
)

var grugSession *grug.GrugSession

func main() {
	log.Println("[INFO] Starting Grug session")
	grugSession = &grug.GrugSession{}
	grugSession.New("grug.yaml")

	log.Println("[INFO] Grug is now running (CTRL-C to exit)")
	// wait for a termination signal
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	grugSession.Close()
}
