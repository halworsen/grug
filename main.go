package main

import (
	"log"
	"main/grug"
	"os"
	"os/signal"
	"syscall"
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
