-- JoPanel Database Schema (MariaDB)

CREATE DATABASE IF NOT EXISTS jopanel;
USE jopanel;

-- Users (Admins and Hosting Users)
CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(64) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role ENUM('admin', 'reseller', 'user') NOT NULL DEFAULT 'user',
    package_id INT UNSIGNED,
    status ENUM('active', 'suspended', 'pending') NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Hosting Packages
CREATE TABLE packages (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    disk_quota_mb INT NOT NULL DEFAULT 1000, -- 0 for unlimited
    bandwidth_mb INT NOT NULL DEFAULT 10000,
    max_domains INT NOT NULL DEFAULT 1,
    max_databases INT NOT NULL DEFAULT 1,
    max_ftp INT NOT NULL DEFAULT 1,
    max_emails INT NOT NULL DEFAULT 1
);

-- Domains
CREATE TABLE domains (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    domain_name VARCHAR(255) NOT NULL UNIQUE,
    document_root VARCHAR(255) NOT NULL,
    type ENUM('primary', 'addon', 'subdomain', 'alias') NOT NULL DEFAULT 'primary',
    is_ssl_active BOOLEAN DEFAULT FALSE,
    php_version VARCHAR(10) DEFAULT '8.2',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Databases
CREATE TABLE databases (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    db_name VARCHAR(64) NOT NULL,
    db_user VARCHAR(64) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- FTP Accounts
CREATE TABLE ftp_accounts (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    username VARCHAR(64) NOT NULL UNIQUE,
    home_dir VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Settings (Global Config)
CREATE TABLE settings (
    `key` VARCHAR(100) PRIMARY KEY,
    `value` TEXT
);
