const { Client, GatewayIntentBits } = require ('discord.js')
const { routeInteractions } = require (`./handlers`)
const { startMonitor } = require (`./monitor`)
const http = require('http');
require('dotenv').config();

const botToken = process.env.BOT_TOKEN;

if (!botToken) {
    console.log("No Bot Token Provided")
    process.exit(1);
}

const client = new Client({
    intents: [GatewayIntentBits.Guilds] 
});

const port = process.env.PORT || 8080; // Cloud Run injects the PORT variable
http.createServer((req, res) => {
    res.writeHead(200, { 'Content-Type': 'text/plain' });
    res.end('Bot is undocked and scouting! o7');
}).listen(port, () => {
    console.log(`📡 Health check server listening on port ${port}`);
});

// Modern Philosophy: Respond to a specific user action (Interaction)
client.on('interactionCreate', async (interaction) => {
    // Check if the interaction is a Slash Command
    if (!interaction.isChatInputCommand()) return;
    await routeInteractions(interaction)
});

client.once('ready', () => {
    console.log(`Logged in as ${client.user.tag}!`);
    console.log("Bot is now running. Press CTRL-C to exit.");

    startMonitor();
});

client.login(botToken).catch(err => {
    console.error("error opening connection,", err);
});

process.on('SIGINT', () => {
    console.log("Shutting down the bot.");
    client.destroy(); // Properly close the discord connection
    process.exit(0);
});
