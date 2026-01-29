# JoPanel - Modern Open Source Hosting Control Panel

JoPanel is a lightweight, modern web hosting control panel for Linux, designed to replace cPanel/WHM for personal and small-to-medium hosting environments.

## Features
- **Admin Panel (WHM-like)**: Manage users, packages, and system status.
- **User Panel (cPanel-like)**: File Manager, Database Management, Domain Handling.
- **Architecture**: Secure separation of concerns with a core API and a privileged Linux Agent.
- **Tech Stack**: Go (Backend/Agent), React/Shadcn (Frontend), MariaDB, Nginx.

## Directory Structure
- `/backend`: Go API Server (Gin + GORM).
- `/frontend`: React SPA (Vite + Tailwind).
- `/agent`: Privileged System Agent (Go).
- `/docs`: Documentation and Planning.

## Development Setup

### Prerequisites
- Go 1.21+
- Node.js 18+ & NPM
- Docker & Docker Compose

### Running Locally
1. **Start Database**:
   ```bash
   docker-compose up -d
   ```
2. **Backend**:
   ```bash
   cd backend
   go run main.go
   ```
3. **Agent** (Dry Run Mode):
   ```bash
   cd agent
   export DRY_RUN=true
   go run main.go
   ```
4. **Frontend**:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

## Installation (Production)
Run the installer script on a fresh Ubuntu 22.04 server:
```bash
sudo ./install.sh
```

## Security
- The Agent runs as root but accepts commands ONLY from the Backend via a secure token/mTLS.
- The Backend handles all user input validation and business logic.
- File system isolation is enforced via standard Linux permissions (and future Chroot).
