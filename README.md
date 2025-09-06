
# WiNGSPAN Thera Bot 🚀

A simple, efficient Discord bot written in Go that monitors the EVE Scout API for new wormhole connections and provides real-time notifications. It also allows users to query for currently active routes.

This project was built as a practical learning exercise in Go, covering API interaction, concurrency with goroutines, and Discord bot development.

* **Deployment Steps**: [View the Wiki Page](https://github.com/ScotDex/TheraBot/wiki/Deployment)
* **Visuals**: [View the Gallery](https://github.com/ScotDex/TheraBot/wiki/Gallery)
* **Contact**: [voidsparkadmin@proton.me](mailto:voidsparkadmin@proton.me)

***

## ✨ Features

* **Real-time Notifications**: A background poller checks the EVE Scout API every 120 seconds and announces any new wormhole signatures in a designated channel.
* **Command Interface**: Users can interact with the bot using simple commands.
    * `!allroutes`: Lists the 25 most recent, active wormhole connections (Discord has a hardlimit)
    * `!health`: Checks the current status of the EVE Scout API.
    * `!theraroutes`: Shows all Thera specific WHs
    * `!turnurroutes`: Shows all Turnur specific WHs
    * `!help`: Provides a list of available commands (Should you forget)
* **Graceful Shutdown**: The bot is designed to shut down cleanly when the process is terminated (e.g., via CTRL-C).
* **Configuration via `.env`**: All sensitive information and settings are managed through a `.env` file, keeping secrets out of the code.

***
