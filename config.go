package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Config struct {
	NotificationChannels map[string]string `json:"notification_channels"`

	mu sync.RWMutex
}

var config Config

const configFile = "config.json"

func loadConfig() error {
	config.mu.Lock()
	defer config.mu.Unlock()

	config.NotificationChannels = make(map[string]string)

	file, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Print("Config file not found, creating a new one.")
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &config.NotificationChannels)

}

func saveConfig() error {
	config.mu.RLock()
	defer config.mu.RUnlock()

	data, err := json.MarshalIndent(config.NotificationChannels, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configFile, data, 0644)

}

func setNotificationChannel(guildID, ChannelID string) error {
	config.mu.Lock()
	defer config.mu.Unlock()

	config.NotificationChannels[guildID] = ChannelID
	return saveConfig()
}
