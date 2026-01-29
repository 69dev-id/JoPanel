package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"jopanel/agent/models"
)

const BaseDir = "/etc/jopanel/users"
const LockDir = "/run/jopanel/locks"

type Manager struct{}

func NewManager() *Manager {
	// Ensure directories exist
	os.MkdirAll(BaseDir, 0700)
	os.MkdirAll(LockDir, 0755)
	return &Manager{}
}

// AcquireLock uses syscall.Flock for exclusive locking
func (m *Manager) AcquireLock(username string) (*os.File, error) {
	lockPath := filepath.Join(LockDir, fmt.Sprintf("user-%s.lock", username))
	file, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	// Exclusive lock, blocking
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	if err != nil {
		file.Close()
		return nil, err
	}
	return file, nil
}

func (m *Manager) ReleaseLock(file *os.File) {
	if file != nil {
		syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
		file.Close()
	}
}

// AtomicWrite writes JSON to a temp file and renames it
func (m *Manager) AtomicWrite(path string, data interface{}) error {
	tmpPath := path + ".tmp"
	
	// Open/Create temp file
	file, err := os.OpenFile(tmpPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	
	// Encode JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		file.Close()
		return err
	}

	// Sync to disk
	if err := file.Sync(); err != nil {
		file.Close()
		return err
	}
	file.Close()

	// Atomic Rename
	return os.Rename(tmpPath, path)
}

func (m *Manager) LoadAccount(username string) (*models.AccountConfig, error) {
	path := filepath.Join(BaseDir, username, "account.json")
	var cfg models.AccountConfig
	err := m.readJSON(path, &cfg)
	return &cfg, err
}

func (m *Manager) SaveAccount(username string, cfg *models.AccountConfig) error {
	path := filepath.Join(BaseDir, username, "account.json")
	return m.AtomicWrite(path, cfg)
}

// Helper to read JSON
func (m *Manager) readJSON(path string, dest interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(dest)
}

// InitialUserSetup creates the directory structure for a new user
func (m *Manager) InitialUserSetup(username string) error {
	userDir := filepath.Join(BaseDir, username)
	if err := os.MkdirAll(userDir, 0700); err != nil {
		return err
	}
	// TODO: Chown to root:root (already default for new dirs by root process)
	return nil
}

func (m *Manager) SaveAll(username string, acc *models.AccountConfig, limits *models.LimitsConfig, domains *models.DomainsConfig, services *models.ServicesConfig, ssh *models.SSHConfig, meta *models.MetaConfig) error {
	dir := filepath.Join(BaseDir, username)
	if err := m.AtomicWrite(filepath.Join(dir, "account.json"), acc); err != nil { return err }
	if err := m.AtomicWrite(filepath.Join(dir, "limits.json"), limits); err != nil { return err }
	if err := m.AtomicWrite(filepath.Join(dir, "domains.json"), domains); err != nil { return err }
	if err := m.AtomicWrite(filepath.Join(dir, "services.json"), services); err != nil { return err }
	if err := m.AtomicWrite(filepath.Join(dir, "ssh.json"), ssh); err != nil { return err }
	if err := m.AtomicWrite(filepath.Join(dir, "meta.json"), meta); err != nil { return err }
	return nil
}
