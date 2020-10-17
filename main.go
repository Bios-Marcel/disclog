package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Bios-Marcel/discordgo"
)

func main() {
	var token string
	var channels string
	flag.StringVar(&token, "token", "", "Token for authentication with discord backend")
	flag.StringVar(&channels, "channels", "", "comma separated list of channel IDs to log messages from")
	flag.Parse()

	token = strings.TrimSpace(token)
	channels = strings.TrimSpace(channels)

	if token == "" {
		fmt.Fprintf(os.Stderr, "Please supply a token via --token")
		os.Exit(1)
	}
	if channels == "" {
		fmt.Fprintf(os.Stderr, "Please supply a list of channel IDs via")
		os.Exit(1)
	}

	channelIDs := strings.Split(channels, ",")

	session, sessionError := discordgo.NewWithToken(token)
	if sessionError != nil {
		fmt.Fprintf(os.Stderr, "Error establishing connection: %s\n", sessionError.Error())
		os.Exit(1)
	}

	session.AddHandler(func(s *discordgo.Session, messageCreate *discordgo.MessageCreate) {
		if !contains(channelIDs, messageCreate.ChannelID) {
			return
		}

		data, jsonError := json.Marshal(messageCreate)
		if jsonError != nil {
			fmt.Fprintf(os.Stderr, "Error marshalling data: %s\n", jsonError)
		}

		fmt.Fprintln(os.Stdout, string(data))
	})

	gatewayError := session.Open()
	if gatewayError != nil {
		fmt.Fprintf(os.Stderr, "Error opening gateway connection: %s\n", gatewayError)
		os.Exit(1)
	}
	defer session.Close()

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-killSignal
}

func contains(arr []string, element string) bool {
	for _, arrayElement := range arr {
		if arrayElement == element {
			return true
		}
	}

	return false
}
