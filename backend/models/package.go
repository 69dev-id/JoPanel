package models

import "gorm.io/gorm"

type Package struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"not null" json:"name"`
	DiskQuotaMB   int            `json:"disk_quota_mb"` // 0 for unlimited
	BandwidthMB   int            `json:"bandwidth_mb"`
	MaxDomains    int            `json:"max_domains"`
	MaxDatabases  int            `json:"max_databases"`
	MaxEmails     int            `json:"max_emails"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
