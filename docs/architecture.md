# JoPanel Architecture

## High Level Overview

```mermaid
graph TD
    UserClient[User Browser]
    AdminClient[Admin Browser]
    
    ReverseProxy[Nginx / Apache]
    
    subgraph "Server (JoPanel Host)"
        API[Backend API (Go/Gin) - Port 8080]
        Agent[Linux Agent (Go) - Root Privileges]
        DB[(MariaDB)]
        Redis[(Redis - Optional Cache)]
        
        System[Linux System Services]
        Cron[Cron Jobs]
        FS[File System]
    end

    UserClient -->|HTTPS| ReverseProxy
    AdminClient -->|HTTPS| ReverseProxy
    
    ReverseProxy -->|Proxy /api| API
    ReverseProxy -->|Serve Static| FrontendFiles
    
    API -->|SQL| DB
    API -->|Local Socket / RPC| Agent
    
    Agent -->|Exec| System
    Agent -->|Exec| Cron
    Agent -->|Files| FS
```

## Component Details

### 1. Backend API (Go)
The brain of the operation. It does NOT run as root.
- **Auth:** JWT for stateless auth.
- **Database:** Stores user meta-data, domains info, packages, invoices (if any).
- **Communication with Agent:** When a system change is needed (e.g., "Create User"), the API sends a structured command to the Agent.

### 2. Linux Agent (Go)
The muscle. Runs as `root` (Systemd service).
- **Listening:** Listens on a Unix Domain Socket (most secure for local IPC) or a localhost-only TCP port with mTLS.
- **Validation:** Strictly validates all incoming commands against a whitelist.
- **Action:** Executes low-level Linux commands (`useradd`, `passwd`, `nginx reload`, `openvpn`, etc.).
- **Response:** Returns success/failure logic to API.

### 3. Frontend (React)
A Single Page Application (SPA).
- **Admin View:** `admin.jopanel.com` or `/admin`
- **User View:** `domain.com/cpanel` or `/panel`
- **State:** React Query / Zustand.

## Security Model
1. **Separation of Privileges:** The public-facing API does not have root access. If it is compromised, the attacker cannot immediately root the server (they are limited to the `jopanel` user).
2. **Strict Agent Protocol:** The Agent accepts only specific JSON payloads.
    - `{"op": "create_user", "username": "foo", "shell": "..."}` -> Allowed.
    - `{"op": "exec", "cmd": "rm -rf /"}` -> DENIED. No arbitrary command execution.
3. **Jailed Users:** Hosted users are jailed (Chroot or strict permissions) so they can't see each other's files.
