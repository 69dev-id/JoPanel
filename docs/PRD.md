# JoPanel PRD (Product Requirements Document)

## 1. Introduction
JoPanel is a modern, open-source web hosting control panel designed for Linux servers. It aims to provide a cPanel/WHM-like experience with a modern UI and robust security, built on a scalable architecture.

## 2. Goals
- **Functionality:** Match core cPanel/WHM features (User management, File management, Database, DNS, SSH).
- **Security:** Secure default configuration, RBAC, isolated execution (Agent).
- **Usability:** Modern, responsive UI (React + Tailwind).
- **Extensibility:** Open source and modular.

## 3. User Roles
- **Super Admin:** Full system access, manage all users and server settings.
- **Reseller:** Manage own users (future scope, but RBAC should support it).
- **User:** Manage own hosting account (Files, DB, Domains).

## 4. System Components
### 4.1. Backend (API Server)
- **Tech:** Go, Gin, GORM.
- **Responsibility:** Business logic, DB interaction, Auth, API serving.
- **Port:** 8080 (internal), exposed via Nginx/Reverse Proxy.

### 4.2. Linux Agent
- **Tech:** Go (Native Binary).
- **Responsibility:** Executing privileged system commands (useradd, chmod, service restart).
- **Security:** Runs as root, accepts commands *only* from Backend via authenticated channel.

### 4.3. Frontend
- **Tech:** React, TypeScript, Vite, TailwindCSS, Shadcn/ui.
- **Responsibility:** UI for Admin and User.

## 5. Functional Requirements
### 5.1. Admin Panel
- **User Management:** Create, Delete, Suspend, Unsuspend, Modify Users.
- **Impersonation:** Login as User.
- **Packages:** Define limits (Disk, BW, Domains).

### 5.2. User Panel
- **Dashboard:** Stats (Disk, BW), Account Info.
- **File Manager:** Upload, Download, Edit, Zip/Unzip, Permissions.
- **Databases:** Create DB, User, Access.
- **Domains:** Manage Zones, Subdomains.
- **SSH:** Manage Keypairs.

## 6. Non-Functional Requirements
- **Performance:** Fast API response (<100ms for non-system ops).
- **Security:** Zero-trust between Frontend and Agent. Input validation everywhere. System isolation.
- **Compatibility:** Ubuntu 22.04 LTS / Debian 12.
