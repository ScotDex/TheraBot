


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
