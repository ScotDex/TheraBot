const axios = require (`axios`)
const API_BASE = "https://api.eve-scout.com/v2/public/signatures";

async function checkAPIHealth() {
    try {
        const { data } = await axios.get("https://api.eve-scout.com/v2/health");
        return data.api_version;
    } catch (err) {
        console.error("API is down:", err.message);
        return null;
    }
}
async function getSignatures(system = "") {
    try {
        // Construct URL: if system is "thera", it adds the query param
        const url = system ? `${API_BASE}?system_name=${system}` : API_BASE;
        
        const response = await axios.get(url);
        return response.data; // This is your '[]Route' slice equivalent
    } catch (err) {
        console.error("Fetch failed:", err.message);
        return [];
    }

}

module.exports = { checkAPIHealth, getRoutes: getSignatures };