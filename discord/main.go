package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/halworsen/grug"
	"github.com/halworsen/grug/discord/discordgrug"
	"github.com/halworsen/grug/util"
)

var grugSession *grug.GrugSession

func discordMsgHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, discordgrug.DISCORDSESSION, s)
	ctx = context.WithValue(ctx, discordgrug.DISCORDMSG, m)
	grugSession.HandleMessage(ctx, m.Content)
}

func main() {
	log.Println("[INFO] Initializing Grug")
	grugSession = &grug.GrugSession{}
	grugSession.New(os.Getenv("GRUG_CONFIG_FILE"))

	log.Println("[INFO] Starting liveness probe...")
	go util.RunLivenessProbe("5700")

	log.Println("[INFO] Establishing Discord session")
	session, err := discordgo.New("Bot " + os.Getenv("GRUG_DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln("[ERROR] Failed to create Discord session -", err)
	}
	session.AddHandler(discordMsgHandler)

	// Start listening for commands
	err = session.Open()
	if err != nil {
		log.Fatalln("[ERROR] Failed to open connection -", err)
	}

	log.Println("[INFO] Grug is now running (CTRL-C to exit)")
	// wait for a termination signal
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
