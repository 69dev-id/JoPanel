package sysops

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// Regex for valid linux username: starts with letter, 3-16 chars, lowercase system-safe
var usernameRegex = regexp.MustCompile(`^[a-z][a-z0-9]{2,15}$`)

type UserOps struct{}

func ValidateUsername(username string) error {
	if !usernameRegex.MatchString(username) {
		return fmt.Errorf("invalid username format (must be 3-16 chars, lowercase alphanumeric, start with letter)")
	}
	return nil
}

func CreateSystemUser(username, password string) error {
	if err := ValidateUsername(username); err != nil {
		return err
	}

	// 1. Create Group
	// groupadd --force <username>
	if err := exec.Command("groupadd", "--force", username).Run(); err != nil {
		return fmt.Errorf("failed to create group: %v", err)
	}

	// 2. Create User
	// useradd --create-home --home-dir /home/<username> --gid <username> --shell /bin/bash --comment "JoPanel User" <username>
	homeDir := fmt.Sprintf("/home/%s", username)
	cmd := exec.Command("useradd", 
		"--create-home", 
		"--home-dir", homeDir, 
		"--gid", username, 
		"--shell", "/bin/bash", 
		"--comment", "JoPanel User", 
		username,
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to create user: %s", string(out))
	}

	// 3. Set Permissions (750)
	// chmod 750 /home/<username>
	if err := exec.Command("chmod", "750", homeDir).Run(); err != nil {
		return fmt.Errorf("failed to chmod home: %v", err)
	}

	// 4. Set Password
	if err := SetUserPassword(username, password); err != nil {
		return err
	}

	return nil
}

func SetUserPassword(username, password string) error {
	// Usage: printf 'user:pass' | chpasswd
	cmd := exec.Command("chpasswd")
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%s:%s", username, password))
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to set password: %s", string(out))
	}
	return nil
}

func LockUser(username string) error {
	// passwd -l <user>
	if err := exec.Command("passwd", "-l", username).Run(); err != nil {
		return err
	}
	// usermod -s /usr/sbin/nologin <user>
	return exec.Command("usermod", "-s", "/usr/sbin/nologin", username).Run()
}

func UnlockUser(username string) error {
	// passwd -u <user>
	if err := exec.Command("passwd", "-u", username).Run(); err != nil {
		return err
	}
	// usermod -s /bin/bash <user>
	return exec.Command("usermod", "-s", "/bin/bash", username).Run()
}
