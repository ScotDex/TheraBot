const api = require (`./api`)
const { EmbedBuilder } = require('discord.js');

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
    
    const signatures = await api.getSignatures(systemName);
    
    if (!signatures || signatures.length === 0) {
        return interaction.editReply("🌑 **Space is quiet...** No active signatures found.");
    }

    const titleSystem = systemName ? systemName.toUpperCase() : 'EVE SCOUT';
    const embed = new EmbedBuilder()
        .setTitle(`🛰️ Live ${titleSystem} Connections`)
        .setURL('https://www.eve-scout.com/thera/') // Link to the actual map
        .setColor(0xFFA500) // EVE Scout Orange
        .setThumbnail('https://www.eve-scout.com/images/eve-scout-logo.png')
        .setDescription(`Found **${signatures.length}** signatures. *Newest 25 listed below.*`)
        .setTimestamp()
        .setFooter({ text: 'o7 Fly Safe • Data via EVE Scout API' });

    signatures.slice(0, 25).forEach(sig => {
        // Convert ISO date to Unix Timestamp for Discord's dynamic timers
        const unixExpiry = Math.floor(new Date(sig.expires_at).getTime() / 1000);
        
        // Add dotlan links to the system names for utility
        const outLink = `[${sig.out_system_name}](https://evemaps.dotlan.net/system/${sig.out_system_name.replace(' ', '_')})`;
        const inLink = `[${sig.in_system_name}](https://evemaps.dotlan.net/system/${sig.in_system_name.replace(' ', '_')})`;

        embed.addFields({
            name: `🔹 ${sig.wh_type} | ${sig.out_system_name} ↔ ${sig.in_system_name}`,
            value: [
                `**Routes:** ${outLink} ↔ ${inLink}`,
                `**Signatures:** \`${sig.out_signature}\` ↔ \`${sig.in_signature}\``,
                `**Class:** \`${sig.in_system_class}\` | **Expires:** <t:${unixExpiry}:R>`
            ].join('\n'),
            inline: false // False makes it look much better on mobile
        });
    });

    await interaction.editReply({ embeds: [embed] });
}


module.exports = { handleHealth, handleHelp, handleRoutes };