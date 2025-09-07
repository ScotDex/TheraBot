package main

import (
	"fmt"
	"log"
	"time"

	"thera-bot/api"

	"github.com/bwmarrin/discordgo"
)

func StartNotifier(s *discordgo.Session, channelID string) {
	log.Println("Starting poller...")
	seenSigs := make(map[string]bool)
	log.Println("logging existing signatures...")
	initialSigs, err := api.AllRoutes()
	if err != nil {
		log.Println("Error fetching initial signatures:", err)
		return
	}

	for _, sig := range initialSigs {
		seenSigs[sig.ID] = true

	}
	log.Printf("Mapped %d existing signatures", len(seenSigs))
	ticker := time.NewTicker(120 * time.Second) // Check for new signatures every 120 seconds
	defer ticker.Stop()
	for range ticker.C {
		log.Println("Checking for new signatures...")
		currentSigs, err := api.AllRoutes()
		if err != nil {
			log.Println("Error fetching current signatures:", err)
			continue
		}
		var newSigs []api.Route
		for _, sig := range currentSigs {
			if _, found := seenSigs[sig.ID]; !found {
				newSigs = append(newSigs, sig)
			}
		}

		if len(newSigs) > 0 {
			log.Printf("Found %d new signatures", len(newSigs))
			embed := &discordgo.MessageEmbed{
				Title:  "🛰️ New Wormhole Connections Detected!",
				Color:  0xFFA500, // On brand orange colour.
				Fields: []*discordgo.MessageEmbedField{},
			}
			for _, newSig := range newSigs {
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
					Name: fmt.Sprintf("Route: %s <-> %s", newSig.OutSystemName, newSig.InSystemName),
					Value: fmt.Sprintf(
						"**Thera/Turnur Sig:** `%s`\n**In Sig:** `%s`\n**Expires at:** %s\n**Type:** %s",
						newSig.OutSignature,
						newSig.InSignature,
						newSig.ExpiresAt.Format("2006-01-02 15:04 MST"),
						newSig.WhType,
					),
					Inline: false,
				})

			}
			config.mu.RLock()
			var channelsToNotify []string
			for _, channelID := range config.NotificationChannels {
				channelsToNotify = append(channelsToNotify, channelID)
			}
			config.mu.RUnlock()

			// Loop through every registered channel and send the alert
			for _, channelID := range channelsToNotify {
				s.ChannelMessageSendEmbed(channelID, embed)
			}

			for _, newSig := range newSigs {
				seenSigs[newSig.ID] = true
			}

		}

	}
}
