async function routeInteractions(interaction) {
    if (!interaction.isChatInputCommand()) return;

    const { commandName } = interaction;
    try {
        switch (commandName) {
            case 'health':
                await handleHealth(interaction);
                break;
            case 'help':
                await handleHelp(interaction);
                break;
            case 'theraroutes':
                await handleRoutes(interaction, 'thera');
                break;
            case 'turnurroutes':
                await handleRoutes(interaction, 'turnur');
                break;
            case 'allroutes':
                await handleRoutes(interaction, null);
                break;
            default:
                await interaction.reply({ content: "Unknown Command", ephemeral: true });
        } 
    } catch (error) {
        console.error(`Execution Error in ${commandName}:`, error);
        if (!interaction.replied) {
            await interaction.reply({ content: "There was an error executing this command!", ephemeral: true });
        }
    }
}