const { WebhookClient, EmbedBuilder } = require('discord.js');
const api = require('./api');
require('dotenv').config();

// Use the Webhook URL from your .env file
const webhook = new WebhookClient({ url: process.env.WEBHOOK_URL });

// In Go: seenSigs := make(map[string]bool)
// In Node: A Set is the modern, faster way to track unique IDs
const seenSigs = new Set();

async function startMonitor() {
    console.log("🛰️ Initializing Thera Poller...");

    try {
        // 1. Log existing signatures so we don't ping them on startup
        const initialSigs = await api.getSignatures();
        initialSigs.forEach(sig => seenSigs.add(sig.id));
        console.log(`Mapped ${seenSigs.size} existing signatures.`);

        // 2. Start the ticker (In Go: time.NewTicker)
        setInterval(async () => {
            console.log("Checking for new signatures...");
            
            const currentSigs = await api.getSignatures();
            if (!currentSigs) return;

            // Find signatures not in our 'seen' Set
            const newSigs = currentSigs.filter(sig => !seenSigs.has(sig.id));

            if (newSigs.length > 0) {
                console.log(`Found ${newSigs.length} new signatures!`);
                const limitedSigs = newSigs.slice(0, 25);
                const embed = new EmbedBuilder()
                    .setTitle("🛰️ New Wormhole Connections Detected!")
                    .setColor(0xFFA500) // Brand Orange
                    .setTimestamp();

                newSigs.forEach(newSig => {
                    embed.addFields({
                        name: `Route: ${newSig.out_system_name} <-> ${newSig.in_system_name}`,
                        value: `**Thera/Turnur Sig:** \`${newSig.out_signature}\`\n**In Sig:** \`${newSig.in_signature}\`\n**Type:** ${newSig.wh_type}`,
                        inline: false
                    });
                    
                    // Add to seen list so we don't notify again
                    seenSigs.add(newSig.id);
                });

                // Send via Webhook (Replaces the loop through channels)
                await webhook.send({ embeds: [embed] });
            }
        }, 120000); // 120 seconds

    } catch (error) {
        console.error("Monitor Error:", error.message);
    }
}

module.exports = { startMonitor };