package main

import (
	"fmt"
	"log"
	"thera-bot/api"
	"time"

	"github.com/bwmarrin/discordgo"
)

func handleHealthCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	health, err := api.CheckApiHealth()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error checking API health: "+err.Error())
		return
	}
	reply := fmt.Sprintf("EVE Scout API Status: **%s**", health.Version)
	s.ChannelMessageSend(m.ChannelID, reply)
}

func handleHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "WiNGSPAN Thera Bot Commands",
		Description: "Here are the commands you can use, and some questions you may have that have been answered.",
		Color:       0xFFA500, // Your brand's orange color
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "!health",
				Value:  "Checks if the API is alive (If it isnt you will have to start scanning)",
				Inline: false,
			},
			{
				Name:   "!allroutes",
				Value:  "Fetch all wormhole routes (Discord has a hard limit of 25 routes per message so you wont get all of them)",
				Inline: false,
			},
			{
				Name:   "!theraroutes",
				Value:  "Fetch Thera-specific wormhole routes.",
				Inline: false,
			},
			{
				Name:   "!turnurroutes",
				Value:  "Fetch Turnur-specific wormhole routes.",
				Inline: false,
			},
			{
				Name:   "Dex, why did you make this bot when there are so many others?",
				Value:  "I am a tech nerd and I like to make things, im also lazy so I made this thing to avoid having to check the EVE Scout website all the time, it also annoyed me that people make stuff then abandon it.",
				Inline: false,
			},
			{
				Name:   "Who is Dex?",
				Value:  "Dexomus Viliana, if you like the bot send me isk, if you dont like the bot send me isk anyway, I like isk and it pays for all my Loki's",
				Inline: false,
			},
			{
				Name:   "I want new features, what do?",
				Value:  "Sure, go post your ideas on my GitHub. I’ll read them, nod, and then ignore you https://github.com/ScotDex/TheraBot.",
				Inline: false,
			},
		},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func handleAllRoutesCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "🔍 Fetching EVE Scout signatures...")
	signatures, err := api.AllRoutes()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error fetching signatures: "+err.Error())
		return
	}
	if len(signatures) == 0 {
		s.ChannelMessageSend(m.ChannelID, "No signatures found.")
		return
	}
	embed := &discordgo.MessageEmbed{
		Title:       "🛰️ Live Wormhole Connections",
		Description: fmt.Sprintf("Found **%d** active signatures, here are the newest 25 ", len(signatures)),
		Color:       0xFFA500,
		Fields:      []*discordgo.MessageEmbedField{},
		Timestamp:   time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "WiNGSPAN Thera Bot | Powered by EVE Scout API ",
		},
	}
	for i, sig := range signatures {
		if i >= 25 {
			break
		}
		remainingTime := time.Until(sig.ExpiresAt).Round(time.Minute)
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name: fmt.Sprintf(sig.WhType+":%s <-> %s", sig.OutSystemName, sig.InSystemName),
			Value: fmt.Sprintf(
				"**Thera/Turnur:** `%s`\n**Entry Sig:** `%s`\n**Expires in:** %s \n**Class:** %s",
				sig.OutSignature, sig.InSignature, remainingTime, sig.InSystemClass,
			),
			Inline: true,
		})
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func handleTheraRoutesCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "🔍 Fetching EVE Scout signatures...")
	signatures, err := api.TheraRoutes()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error fetching signatures: "+err.Error())
		return
	}
	if len(signatures) == 0 {
		s.ChannelMessageSend(m.ChannelID, "No signatures found.")
		return
	}
	embed := &discordgo.MessageEmbed{
		Title:       "🛰️ Live Thera Wormhole Connections",
		Description: fmt.Sprintf("Found **%d** active signatures.", len(signatures)),
		Color:       0xFFA500,
		Fields:      []*discordgo.MessageEmbedField{},
		Timestamp:   time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "WiNGSPAN Thera Bot | Powered by EVE Scout API ",
		},
	}
	for i, sig := range signatures {
		if i >= 25 {
			break
		}
		remainingTime := time.Until(sig.ExpiresAt).Round(time.Minute)
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name: fmt.Sprintf(sig.WhType+":%s <-> %s", sig.OutSystemName, sig.InSystemName),
			Value: fmt.Sprintf(
				"**Thera:** `%s`\n**Entry Sig:** `%s`\n**Expires in:** %s \n**Class:** %s",
				sig.OutSignature, sig.InSignature, remainingTime, sig.InSystemClass,
			),
			Inline: true,
		})
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func handleTurnurRoutesCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "🔍 Fetching EVE Scout signatures...")
	signatures, err := api.TurnurRoutes()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error fetching signatures: "+err.Error())
		return
	}
	if len(signatures) == 0 {
		s.ChannelMessageSend(m.ChannelID, "No signatures found.")
		return
	}
	embed := &discordgo.MessageEmbed{
		Title:       "🛰️ Live Turnur Wormhole Connections",
		Description: fmt.Sprintf("Found **%d** active signatures. ", len(signatures)),
		Color:       0xFFA500,
		Fields:      []*discordgo.MessageEmbedField{},
		Timestamp:   time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "WiNGSPAN Thera Bot | Powered by EVE Scout API ",
		},
	}
	for i, sig := range signatures {
		if i >= 25 {
			break
		}
		remainingTime := time.Until(sig.ExpiresAt).Round(time.Minute)
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name: fmt.Sprintf(sig.WhType+":%s <-> %s", sig.OutSystemName, sig.InSystemName),
			Value: fmt.Sprintf(
				"**Turnur:** `%s`\n**Entry Sig:** `%s`\n**Expires in:** %s \n**Class:** %s",
				sig.OutSignature, sig.InSignature, remainingTime, sig.InSystemClass,
			),
			Inline: true,
		})
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

// In commands.go

// handleSetChannelCommand allows a server admin to set the notification channel.
func handleSetChannelCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		log.Printf("Error getting user permissions: %v", err)
		return
	}
	// Check if the user has "Manage Server" permissions.
	if p&discordgo.PermissionManageGuild == 0 {
		s.ChannelMessageSend(m.ChannelID, "You must have the `Manage Server` permission to use this command.")
		return
	}

	// Save the setting for this server (Guild).
	err = SetNotificationChannel(m.GuildID, m.ChannelID)
	if err != nil {
		log.Printf("Error saving config: %v", err)
		s.ChannelMessageSend(m.ChannelID, "Sorry, there was an error saving the configuration.")
		return
	}

	// Send a confirmation message.
	confirmationMsg := fmt.Sprintf("✅ This channel has been set for notifications.")
	s.ChannelMessageSend(m.ChannelID, confirmationMsg)
}
