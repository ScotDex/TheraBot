const { WebhookClient, EmbedBuilder } = require('discord.js');
const api = require('./api');
require('dotenv').config();

const webhook = new WebhookClient({ url: process.env.WEBHOOK_URL });
const seenSigs = new Set();

async function startMonitor() {
    console.log("🛰️ Initializing Thera Poller...");

    try {
        const initialSigs = await api.getSignatures();
        initialSigs.forEach(sig => seenSigs.add(sig.id));
        console.log(`Mapped ${seenSigs.size} existing signatures.`);
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
                    seenSigs.add(newSig.id);
                });
                await webhook.send({ embeds: [embed] });
            }
        }, 120000); // 120 seconds

    } catch (error) {
        console.error("Monitor Error:", error.message);
    }
}
module.exports = { startMonitor };