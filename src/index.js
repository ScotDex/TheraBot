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

// Left health checker in anyways

const port = process.env.PORT || 8080; 
http.createServer((req, res) => {
    res.writeHead(200, { 'Content-Type': 'text/plain' });
    res.end('Bot is undocked and scouting! o7');
}).listen(port, () => {
    console.log(`📡 Health check server listening on port ${port}`);
});


client.on('interactionCreate', async (interaction) => {
    if (!interaction.isChatInputCommand()) return;
    await routeInteractions(interaction)
});

client.once('clientReady', () => {
    console.log(`Logged in as ${client.user.tag}!`);
    console.log("Bot is now running. Press CTRL-C to exit.");

    startMonitor();
});

client.login(botToken).catch(err => {
    console.error("error opening connection,", err);
});

process.on('SIGINT', () => {
    console.log("Shutting down the bot.");
    client.destroy(); 
    process.exit(0);
});
