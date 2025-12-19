const { REST, Routes } = require('discord.js');
require('dotenv').config();

// 1. Define your commands exactly as they should appear in Discord
const commands = [
    { name: 'health', description: 'Check API status' },
    { name: 'help', description: 'Show bot commands' },
    { name: 'theraroutes', description: 'Get Thera connections' },
    { name: 'turnurroutes', description: 'Get Turnur connections' }, // Added!
    { name: 'allroutes', description: 'Get all EVE Scout connections' }
];

// 2. Prepare the REST manager
const rest = new REST({ version: '10' }).setToken(process.env.BOT_TOKEN);

// 3. The Deployment Logic
(async () => {
    try {
        console.log(`Started refreshing ${commands.length} application (/) commands.`);

        // Routes.applicationCommands(CLIENT_ID) registers them globally
        // This can take up to an hour to update across all servers, 
        // but usually happens instantly.
        await rest.put(
            Routes.applicationCommands(process.env.CLIENT_ID),
            { body: commands },
        );

        console.log('Successfully reloaded application (/) commands.');
    } catch (error) {
        console.error(error);
    }
})();