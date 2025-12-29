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
        
        const url = system ? `${API_BASE}?system_name=${system}` : API_BASE;
        
        const response = await axios.get(url);
        return response.data; 
    } catch (err) {
        console.error("Fetch failed:", err.message);
        return [];
    }

}

module.exports = { checkHealth: checkAPIHealth, getSignatures };