// handlers.go
package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// messageCreate is the central hub for all incoming messages.
// It acts as a router, directing commands to their appropriate handlers.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself to prevent loops.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// This switch statement checks the message content and calls the correct function.
	switch {
	case m.Content == "!help":
		handleHelpCommand(s, m)

	case m.Content == "!health":
		handleHealthCommand(s, m)

	case m.Content == "!setchannel":
		handleSetChannelCommand(s, m)

	case strings.HasPrefix(m.Content, "!allroutes"):
		handleAllRoutesCommand(s, m)

	case strings.HasPrefix(m.Content, "!help"):
		handleHelpCommand(s, m)

	case strings.HasPrefix(m.Content, "!turnurroutes"):
		handleTurnurRoutesCommand(s, m)

	case strings.HasPrefix(m.Content, "!theraroutes"):
		handleTheraRoutesCommand(s, m)
	}
}
