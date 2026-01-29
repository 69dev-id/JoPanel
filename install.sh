#!/bin/bash
set -e

# JoPanel Installer for Ubuntu 22.04 / Debian 12

echo "Starting JoPanel Installation..."

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit 1
fi

# Update System
apt update && apt upgrade -y

# Install Dependencies
apt install -y nginx mariadb-server curl git unzip redis-server

# Install Go (latest)
if ! command -v go &> /dev/null; then
    echo "Installing Go..."
    wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
fi

# Setup Database
echo "Configuring Database..."
mysql -e "CREATE DATABASE IF NOT EXISTS jopanel;"
mysql -e "CREATE USER IF NOT EXISTS 'jopanel'@'localhost' IDENTIFIED BY 'secure_password_please_change';"
mysql -e "GRANT ALL PRIVILEGES ON jopanel.* TO 'jopanel'@'localhost';"
mysql -e "FLUSH PRIVILEGES;"

# Setup Project Directory
mkdir -p /opt/jopanel
echo "Copying files to /opt/jopanel..."
cp -r backend agent frontend /opt/jopanel/
# (In real scenario, git clone here. For now we assume files are present or copied)

# Build Backend
echo "Building Backend..."
cd /opt/jopanel/backend
go build -o jopanel-api

# Build Agent
echo "Building Agent..."
cd /opt/jopanel/agent
go build -o jopanel-agent

# Create Systemd Services

# Backend Service
cat <<EOF > /etc/systemd/system/jopanel-api.service
[Unit]
Description=JoPanel API Server
After=network.target mysql.service

[Service]
User=root
WorkingDirectory=/opt/jopanel/backend
ExecStart=/opt/jopanel/backend/jopanel-api
Restart=always
EnvironmentFile=/opt/jopanel/backend/.env

[Install]
WantedBy=multi-user.target
EOF

# Agent Service
cat <<EOF > /etc/systemd/system/jopanel-agent.service
[Unit]
Description=JoPanel Linux Agent
After=network.target

[Service]
User=root
WorkingDirectory=/opt/jopanel/agent
ExecStart=/opt/jopanel/agent/jopanel-agent
Restart=always
EnvironmentFile=/opt/jopanel/agent/.env

[Install]
WantedBy=multi-user.target
EOF

# Start Services
systemctl daemon-reload
systemctl enable jopanel-api jopanel-agent
systemctl start jopanel-api jopanel-agent

echo "JoPanel Installed Successfully!"
echo "API running on port 8080"
