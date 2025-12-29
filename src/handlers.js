const commandLogic = require('./commandLogic');

async function routeInteractions(interaction) {
    // 1. Safety Guard
    if (!interaction.isChatInputCommand()) return;

    const { commandName } = interaction;

    // 2. The Logic Map
    // We map the slash command name to the function in commandLogic.js
    const commands = {
        'health': commandLogic.handleHealth,
        'help': commandLogic.handleHelp,
        'theraroutes': (int) => commandLogic.handleRoutes(int, 'thera'),
        'turnurroutes': (int) => commandLogic.handleRoutes(int, 'turnur'),
        'allroutes': (int) => commandLogic.handleRoutes(int, null),
    };

    try {
        const execute = commands[commandName];

        if (execute) {
            await execute(interaction);
        } else {
            await interaction.reply({ content: "Command not found.", ephemeral: true });
        }
    } catch (error) {
        console.error(`Error executing ${commandName}:`, error);
        
        // Always check if we already replied to avoid "Interaction already acknowledged" errors
        const errorMessage = { content: 'There was an error while executing this command!', ephemeral: true };
        if (interaction.deferred || interaction.replied) {
            await interaction.editReply(errorMessage);
        } else {
            await interaction.reply(errorMessage);
        }
    }
}

module.exports = { routeInteractions };