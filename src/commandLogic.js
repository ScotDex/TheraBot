const api = require (`./api`)

async function handleHealth(interaction) {
    await interaction.deferReply();
    const version = await api.checkHealth();

    if (version) {
        await interaction.editReply(`✅ EVE Scout API Status: **Online** (Version: ${version})`);
    } else {
        await interaction.editReply("❌ Error checking API health. The scout might be cloaked.");
    }
}

async function handleHelp(interaction) {
    const embed = new EmbedBuilder()
        .setTitle("WiNGSPAN Thera Bot Commands")
        .setDescription("Modernized Thera tracking for 2025.")
        .setColor(0xFFA500) // Orange
        .addFields(
            { name: '/health', value: 'Checks if the API is alive.' },
            { name: '/theraroutes', value: 'Fetch Thera-specific wormhole routes.' },
            { name: '/turnurroutes', value: 'Fetch Turnur-specific wormhole routes.' }, // Added!
            { name: '/allroutes', value: 'Fetch all active EVE Scout signatures.' },
            { name: 'I want new features, what do?', value: 'Post on my GitHub: https://github.com/ScotDex/TheraBot' }
        );

    await interaction.reply({ embeds: [embed] });
}

async function handleRoutes(interaction, systemName) {
    await interaction.deferReply();
    
    const signatures = await api.getRoutes(systemName);
    
    if (!signatures || signatures.length === 0) {
        return interaction.editReply("No active signatures found. Space is quiet...");
    }

    const embed = new EmbedBuilder()
        .setTitle(`🛰️ Live ${systemName || 'EVE Scout'} Connections`)
        .setColor(0xFFA500)
        .setDescription(`Found **${signatures.length}** active signatures. Showing newest 25.`);

    // Map through signatures and add fields (replaces Go's for loop)
    signatures.slice(0, 25).forEach(sig => {
        const expiry = new Date(sig.expires_at);
        const diffMins = Math.round((expiry - new Date()) / 60000);

        embed.addFields({
            name: `${sig.wh_type}: ${sig.out_system_name} ↔ ${sig.in_system_name}`,
            value: `**Sig:** \`${sig.in_signature}\` | **Expires:** ${diffMins}m | **Class:** ${sig.in_system_class}`,
            inline: true
        });
    });

    await interaction.editReply({ embeds: [embed] });
}


module.exports = { handleHealth, handleHelp, handleRoutes };