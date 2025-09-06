// main.go
package main

import (
	"fmt"
	"log"

	"os"
	"os/signal" // Required for signal handling
	"syscall"   // Required for system signals
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"thera-bot/api" // Import the local api package
)

// Main function to start the bot
func main() {
	godotenv.Load()                    //Load .env file
	botToken := os.Getenv("BOT_TOKEN") // Get the bot token from environment variables
	if botToken == "" {                // error handling for missing token
		fmt.Println("No bot token provided")
		return
	}

	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("Unable to start session")
		return // Exit if there's an error creating the Discord session
	}

	dg.AddHandler(messageCreate)                         // Register the messageCreate func as a callback for MessageCreate events
	dg.Identify.Intents = discordgo.IntentsGuildMessages // Set the bot's intents to listen for guild messages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	// Add a defer statement to ensure the connection is closed when main exits.
	defer dg.Close()

	// Start the notifier in a separate goroutine.
	notificationChannelID := os.Getenv("NOTIFICATION_CHANNEL_ID")
	go StartNotifier(dg, notificationChannelID)

	// Go Routine to keep the program running until interrupted.

	// Notification to kill bot.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	// Notify the channel for Interrupt, SIGTERM, or SIGINT signals.
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	// Block execution until a signal is received.
	<-sc
	fmt.Println("Shutting down the bot.")
}

// messageCreate handles incoming Discord messages.
// In main.go
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!health" {
		health, err := api.CheckApiHealth()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error checking API health: "+err.Error())
			return
		}
		// This part is correct
		reply := fmt.Sprintf("EVE Scout API Status: **%s**", health.Version)
		s.ChannelMessageSend(m.ChannelID, reply)
	}

	if m.Content == "!help" {
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
	// Message handler for allroutes ##################################
	if m.Content == "!allroutes" {
		// Call the function that fetches signatures.
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
			Color:       0xFFA500,                         // On brand orange colour.
			Fields:      []*discordgo.MessageEmbedField{}, // Initialize the fields slice
			Timestamp:   time.Now().Format(time.RFC3339),  // Adds a timestamp to the footer
			Footer: &discordgo.MessageEmbedFooter{
				Text: "WiNGSPAN Thera Bot | Powered by EVE Scout API ",
			},
		}

		for i, sig := range signatures {
			if i >= 25 {
				break // Discord has hard limit of 25 fields per embed
			}

			// Calculate time remaining until expiry.
			remainingTime := time.Until(sig.ExpiresAt).Round(time.Minute)

			// Add a field to the embed for this signature.
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name: fmt.Sprintf(sig.WhType+":%s <-> %s", sig.OutSystemName, sig.InSystemName),
				Value: fmt.Sprintf(
					"**Thera/Turnur:** `%s`\n**Entry Sig:** `%s`\n**Expires in:** %s \n**Class:** %s",
					sig.OutSignature,
					sig.InSignature,
					remainingTime,
					sig.InSystemClass,
				),
				Inline: true, // Set to 'true' if you want fields to appear side-by-side
			})
		}

		// 5. Send the completed embed to the channel.
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
	// Message handler for thera routes ##################################

	if m.Content == "!theraroutes" {
		// Call the function that fetches signatures.
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
			Color:       0xFFA500,                         // On brand orange colour.
			Fields:      []*discordgo.MessageEmbedField{}, // Initialize the fields slice
			Timestamp:   time.Now().Format(time.RFC3339),  // Adds a timestamp to the footer
			Footer: &discordgo.MessageEmbedFooter{
				Text: "WiNGSPAN Thera Bot | Powered by EVE Scout API ",
			},
		}

		for i, sig := range signatures {
			if i >= 25 {
				break // Discord has hard limit of 25 fields per embed
			}

			// Calculate time remaining until expiry.
			remainingTime := time.Until(sig.ExpiresAt).Round(time.Minute)

			// Add a field to the embed for this signature.
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name: fmt.Sprintf(sig.WhType+":%s <-> %s", sig.OutSystemName, sig.InSystemName),
				Value: fmt.Sprintf(
					"**Thera:** `%s`\n**Entry Sig:** `%s`\n**Expires in:** %s \n**Class:** %s",
					sig.OutSignature,
					sig.InSignature,
					remainingTime,
					sig.InSystemClass,
				),
				Inline: true, // Set to 'true' if you want fields to appear side-by-side
			})
		}

		// 5. Send the completed embed to the channel.
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	// ####################################################################################################################

	if m.Content == "!turnurroutes" {
		// Call the function that fetches signatures.
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
			Color:       0xFFA500,                         // On brand orange colour.
			Fields:      []*discordgo.MessageEmbedField{}, // Initialize the fields slice
			Timestamp:   time.Now().Format(time.RFC3339),  // Adds a timestamp to the footer
			Footer: &discordgo.MessageEmbedFooter{
				Text: "WiNGSPAN Thera Bot | Powered by EVE Scout API ",
			},
		}

		for i, sig := range signatures {
			if i >= 25 {
				break // Discord has hard limit of 25 fields per embed
			}

			// Calculate time remaining until expiry.
			remainingTime := time.Until(sig.ExpiresAt).Round(time.Minute)

			// Add a field to the embed for this signature.
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name: fmt.Sprintf(sig.WhType+":%s <-> %s", sig.OutSystemName, sig.InSystemName),
				Value: fmt.Sprintf(
					"**Turnur:** `%s`\n**Entry Sig:** `%s`\n**Expires in:** %s \n**Class:** %s",
					sig.OutSignature,
					sig.InSignature,
					remainingTime,
					sig.InSystemClass,
				),
				Inline: true, // Set to 'true' if you want fields to appear side-by-side
			})
		}

		// 5. Send the completed embed to the channel.
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}

// ####################################################################################################################
// StartNotifier periodically checks for new wormhole signatures and sends notifications to a specified Discord channel.
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

		for _, newSig := range newSigs {
			seenSigs[newSig.ID] = true
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
			// Send the embed notification to the channel
			s.ChannelMessageSendEmbed(channelID, embed)
		}
	}
}
