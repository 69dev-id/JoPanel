package models

import "time"

// Schema Version Constant
const CurrentSchemaVersion = 1

// Account Config - /etc/jopanel/users/<user>/account.json
type AccountConfig struct {
	SchemaVersion int       `json:"schema_version"`
	Username      string    `json:"username"`
	UID           int       `json:"uid"`
	GID           int       `json:"gid"`
	PrimaryDomain string    `json:"primary_domain"`
	HomeDir       string    `json:"home_dir"`
	Shell         string    `json:"shell"`
	Status        string    `json:"status"` // active, suspended
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Package       string    `json:"package"`
	ContactEmail  string    `json:"contact_email"`
	LastLoginIP   string    `json:"last_login_ip"`
}

// Limits Config - /etc/jopanel/users/<user>/limits.json
type LimitsConfig struct {
	SchemaVersion    int `json:"schema_version"`
	DiskQuotaMB      int `json:"disk_quota_mb"`
	BandwidthQuotaMB int `json:"bandwidth_quota_mb"`
	MaxAddonDomains  int `json:"max_addon_domains"`
	MaxSubdomains    int `json:"max_subdomains"`
	MaxAliasDomains  int `json:"max_alias_domains"`
	MaxDatabases     int `json:"max_databases"`
	MaxDBUsers       int `json:"max_db_users"`
	MaxFTPAccounts   int `json:"max_ftp_accounts"`
	MaxCronJobs      int `json:"max_cron_jobs"`
	CPULimitPercent  int `json:"cpu_limit_percent"`
	MemoryLimitMB    int `json:"memory_limit_mb"`
	IOLimitMBs       int `json:"io_limit_mb_s"`
}

// Domains Config - /etc/jopanel/users/<user>/domains.json
type DomainInfo struct {
	Domain   string `json:"domain"`
	DocRoot  string `json:"docroot"`
	Parent   string `json:"parent,omitempty"` // for subdomain
}

type DomainsConfig struct {
	SchemaVersion int          `json:"schema_version"`
	Primary       string       `json:"primary"`
	AddonDomains  []DomainInfo `json:"addon_domains"`
	Subdomains    []DomainInfo `json:"subdomains"`
	AliasDomains  []string     `json:"alias_domains"`
}

// Services Config - /etc/jopanel/users/<user>/services.json
type PHPConfig struct {
	Enabled     bool     `json:"enabled"`
	Version     string   `json:"version"`
	FPMPool     string   `json:"fpm_pool"`
	OpenBasedir []string `json:"open_basedir"`
}

type NodeJSApp struct {
	Name         string            `json:"name"`
	Root         string            `json:"root"`
	StartCommand string            `json:"start_command"`
	Port         int               `json:"port"`
	Env          map[string]string `json:"env"`
	Status       string            `json:"status"`
}

type NodeJSConfig struct {
	Enabled bool        `json:"enabled"`
	Apps    []NodeJSApp `json:"apps"`
}

type ServicesConfig struct {
	SchemaVersion int           `json:"schema_version"`
	WebServer     string        `json:"web_server"`
	PHP           PHPConfig     `json:"php"`
	NodeJS        NodeJSConfig  `json:"nodejs"`
	FTP           struct {
		Enabled bool `json:"enabled"`
	} `json:"ftp"`
	Cron struct {
		Enabled bool `json:"enabled"`
	} `json:"cron"`
}

// SSH Config - /etc/jopanel/users/<user>/ssh.json
type SSHConfig struct {
	SchemaVersion        int      `json:"schema_version"`
	Enabled              bool     `json:"enabled"`
	Mode                 string   `json:"mode"`
	AuthorizedKeys       []string `json:"authorized_keys"`
	ForceCommand         string   `json:"force_command"`
	ChrootDir            string   `json:"chroot_dir"`
	AllowPortForwarding  bool     `json:"allow_port_forwarding"`
	AllowAgentForwarding bool     `json:"allow_agent_forwarding"`
	AllowX11Forwarding   bool     `json:"allow_x11_forwarding"`
}

// Meta Config - /etc/jopanel/users/<user>/meta.json
type AuditInfo struct {
	LastAction string    `json:"last_action"`
	LastActor  string    `json:"last_actor"`
	LastTime   time.Time `json:"last_time"`
}

type MetaConfig struct {
	SchemaVersion       int       `json:"schema_version"`
	Notes               string    `json:"notes"`
	LastAppliedRevision string    `json:"last_applied_revision"`
	Audit               AuditInfo `json:"audit"`
}
