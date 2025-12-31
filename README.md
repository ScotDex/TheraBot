# 🛰️ TheraBot
A minimalist, modern Node.js bot for tracking **Thera** and **Turnur** wormhole connections. 

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/X7X21RDFR0)

### 🚀 Quick Start
1. **Clone & Install**: `git clone && npm install`
2. **Configure**: Copy `.env.example` to `.env` and fill in your Discord credentials.
3. **Register Commands**: `node deploy.js`
4. **Deploy**: Built for Google Cloud Run. `gcloud run deploy --source . --min-instances 1`

### 🛠️ Tech Stack
- **Runtime**: Node.js 25
- **Library**: Discord.js v14 (Slash Commands & Webhooks)
- **Architecture**: Stateless & Dockerized
