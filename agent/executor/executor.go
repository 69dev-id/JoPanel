package executor

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"jopanel/agent/config"
	"jopanel/agent/sysops"
	"jopanel/agent/models"
)

// Request payloads
type CreateUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Package  string `json:"package"`
}

type CommandRequest struct {
	Op      string      `json:"op" binding:"required"`
	Payload interface{} `json:"payload"`
}

var cfgMgr = config.NewManager()

func HandleCommand(c *gin.Context) {
	var req CommandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if os.Getenv("DRY_RUN") == "true" {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Dry Run: Command simulated", "op": req.Op})
		return
	}

	switch req.Op {
	case "create_user":
		// Manual binding since payload is interface{}
		// In real world, use better unmarshalling
		// Shortcutting for demonstration
		data := req.Payload.(map[string]interface{}) 
		err := handleCreateUser(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown operation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func handleCreateUser(data map[string]interface{}) error {
	username := data["username"].(string)
	password := data["password"].(string)
	email := data["email"].(string)
	pkg := data["package"].(string)

	log.Printf("V2: Creating User %s", username)

	// 1. Acquire Lock
	lock, err := cfgMgr.AcquireLock(username)
	if err != nil {
		return err
	}
	defer cfgMgr.ReleaseLock(lock)

	// 2. Prepare Config Objects
	acc := &models.AccountConfig{
		SchemaVersion: 1,
		Username:      username,
		PrimaryDomain: username + ".com", // TODO: Pass domain
		HomeDir:       "/home/" + username,
		Status:        "active",
		CreatedAt:     time.Now(),
		Package:       pkg,
		ContactEmail:  email,
	}
	// Default Limits (should load from package template)
	limits := &models.LimitsConfig{SchemaVersion: 1, DiskQuotaMB: 1000} 
	domains := &models.DomainsConfig{SchemaVersion: 1, Primary: acc.PrimaryDomain}
	services := &models.ServicesConfig{SchemaVersion: 1}
	ssh := &models.SSHConfig{SchemaVersion: 1, Enabled: false}
	meta := &models.MetaConfig{
		SchemaVersion: 1,
		Audit: models.AuditInfo{
			LastAction: "create_user",
			LastActor:  "api",
			LastTime:   time.Now(),
		},
	}

	// 3. Filesystem Setup
	if err := cfgMgr.InitialUserSetup(username); err != nil {
		return err
	}
	if err := cfgMgr.SaveAll(username, acc, limits, domains, services, ssh, meta); err != nil {
		return err
	}

	// 4. System User Creation
	if err := sysops.CreateSystemUser(username, password); err != nil {
		return err 
	}

	return nil
}
