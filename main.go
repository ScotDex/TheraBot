// main.go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal" // Required for signal handling
	"syscall"   // Required for system signals

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	// Import the local api package
)

// Main function to start the bot
func main() {
	godotenv.Load()                    //Load .env file
	botToken := os.Getenv("BOT_TOKEN") // Get the bot token from environment variables
	if botToken == "" {                // error handling for missing token
		fmt.Println("No bot token provided")
		return
	}

	if err := loadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// Create a new Discord session using the provided bot token.

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
