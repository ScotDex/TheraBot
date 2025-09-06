Of course. The text from the previous answer is already in the correct Markdown (`.md`) format.

Here it is again, wrapped in a single block for easy copying.

---
### ## Instructions

1.  In your project's main folder, create a new file.
2.  Name the file exactly `README.md`.
3.  Copy the text below and paste it into your new file.

```markdown
# EVE Scout Discord Bot 🚀

A simple, efficient Discord bot written in Go that monitors the EVE Scout API for new wormhole connections and provides real-time notifications. It also allows users to query for currently active routes.

This project was built as a practical learning exercise in Go, covering API interaction, concurrency with goroutines, and Discord bot development.

***

## ✨ Features

* **Real-time Notifications**: A background poller checks the EVE Scout API every 60 seconds and announces any new wormhole signatures in a designated channel.
* **Command Interface**: Users can interact with the bot using simple commands.
    * `!allroutes`: Lists the 10 most recent, active wormhole connections.
    * `!health`: Checks the current status of the EVE Scout API.
* **Graceful Shutdown**: The bot is designed to shut down cleanly when the process is terminated (e.g., via CTRL-C).
* **Configuration via `.env`**: All sensitive information and settings are managed through a `.env` file, keeping secrets out of the code.

***

## 🔧 Setup and Installation

Follow these steps to get the bot running on your own machine.

### **1. Prerequisites**

* You must have **Go** installed on your system.
* You need a **Discord Bot Token**. You can get one from the [Discord Developer Portal](https://discord.com/developers/applications).

### **2. Installation**

```bash
# 1. Clone the repository (replace with your own git URL)
git clone https://github.com/your-username/thera-bot.git

# 2. Navigate into the project directory
cd thera-bot

# 3. Tidy up the Go modules to ensure all dependencies are present
go mod tidy


# Your secret bot token from the Discord Developer Portal
BOT_TOKEN="your_discord_bot_token_here"

# The ID of the Discord channel where you want notifications to be sent
NOTIFICATION_CHANNEL_ID="your_notification_channel_id_here"

# The base URL for the EVE Scout API (pre-configured)
EVE_SCOUT_BASE_URL="https://api-eve.scout.com/v2"

go run main.go

## 🐳 Deploying with Docker Compose (Recommended)

The easiest way to run this bot is with Docker and Docker Compose.

### **1. Prerequisites**
* You must have **Docker** and **Docker Compose** installed on your server.
* You need to have your configured `.env` file in the project directory.

### **2. Running the Bot**
With your `.env` file in place, you can build and start the bot with a single command:
```bash
docker-compose up -d