# Architecture V2: System-Centric & Filesystem Source of Truth

## Core Principles
1.  **Linux User = Hosting Account**: Real `/etc/passwd` users.
2.  **Filesystem Config**: `/etc/jopanel/users/<username>/*.json` is the Single Source of Truth.
3.  **DB for Indexing**: MariaDB stores metadata for Panel Auth/Search only.

## Directory Structure
Root: `/etc/jopanel/users/<username>/` (700 root:root)
Files (600 root:root):

### 1. `account.json`
```json
{
  "schema_version": 1,
  "username": "joko",
  "uid": 12010,
  "gid": 12010,
  "primary_domain": "example.com",
  "home_dir": "/home/joko",
  "shell": "/bin/bash",
  "status": "active",
  "created_at": "2026-01-29T10:00:00+07:00",
  "updated_at": "2026-01-29T10:00:00+07:00",
  "package": "starter",
  "contact_email": "owner@example.com",
  "last_login_ip": null
}
```

### 2. `limits.json`
```json
{
  "schema_version": 1,
  "disk_quota_mb": 10240,
  "bandwidth_quota_mb": 102400,
  "max_addon_domains": 5,
  "max_subdomains": 20,
  "max_alias_domains": 20,
  "max_databases": 10,
  "max_db_users": 10,
  "max_ftp_accounts": 5,
  "max_cron_jobs": 20,
  "cpu_limit_percent": 200,
  "memory_limit_mb": 2048,
  "io_limit_mb_s": 20
}
```

### 3. `domains.json`
```json
{
  "schema_version": 1,
  "primary": "example.com",
  "addon_domains": [
    { "domain": "addon.com", "docroot": "/home/joko/domains/addon.com/public_html" }
  ],
  "subdomains": [
    { "subdomain": "app", "parent": "example.com", "docroot": "/home/joko/domains/example.com/subdomains/app" }
  ],
  "alias_domains": ["www.example.com"]
}
```

### 4. `services.json`
```json
{
  "schema_version": 1,
  "web_server": "nginx",
  "php": {
    "enabled": true,
    "version": "8.2",
    "fpm_pool": "jopanel-joko",
    "open_basedir": ["/home/joko", "/tmp"]
  },
  "nodejs": {
    "enabled": true,
    "apps": []
  },
  "ftp": { "enabled": true },
  "cron": { "enabled": true }
}
```

### 5. `ssh.json`
```json
{
  "schema_version": 1,
  "enabled": true,
  "mode": "sftp_chroot",
  "authorized_keys": [],
  "force_command": "internal-sftp",
  "chroot_dir": "/home/joko",
  "allow_port_forwarding": false,
  "allow_agent_forwarding": false,
  "allow_x11_forwarding": false
}
```

### 6. `meta.json`
```json
{
  "schema_version": 1,
  "notes": "",
  "last_applied_revision": "sha256:....",
  "audit": {
    "last_action": "create_user",
    "last_actor": "admin:root",
    "last_time": "2026-01-29T10:00:00+07:00"
  }
}
```

## System Operations (Agent)

### Authentication
- **Agent API**: accepts JSON commands from Backend.
- **Port:** Localhost only or Unix Socket.

### User Creation Flow
1.  **Validate**: Username regex ^[a-z][a-z0-9]{2,15}$.
2.  **Prepare**: Generate UID/GID.
3.  **Filesystem**:
    - `mkdir -p /etc/jopanel/users/<username>` (0700)
    - Write all `.json` files (0600)
4.  **System**:
    - `groupadd <username>`
    - `useradd -m -d /home/<username> -s /bin/bash -g <username> <username>`
    - `chmod 750 /home/<username>`
5.  **Services**:
    - Trigger Nginx config generation.
    - Trigger PHP-FPM pool generation.
    - Reload services.

### User Edit Flow
1.  **Lock**: `/run/jopanel/locks/user-<username>.lock`.
2.  **Read**: Load current JSONs.
3.  **Update Config**:
    - Apply changes to struct.
    - Write to `.tmp`.
    - Atomic Rename.
4.  **Apply System**:
    - If shell changed -> `usermod -s`.
    - If password changed -> `chpasswd`.
    - If suspended -> `usermod -L` / `passwd -l`.
5.  **Unlock**.

