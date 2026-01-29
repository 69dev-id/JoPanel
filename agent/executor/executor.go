package executor

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type CommandRequest struct {
	Op   string            `json:"op" binding:"required"`
	Args map[string]string `json:"args"`
}

func HandleCommand(c *gin.Context) {
	var req CommandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received Op: %s Args: %v", req.Op, req.Args)

	if os.Getenv("DRY_RUN") == "true" {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Dry Run: Command simulated", "op": req.Op})
		return
	}

	switch req.Op {
	case "create_user":
		err := createUser(req.Args)
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

func createUser(args map[string]string) error {
	username := args["username"]
	// Validate username (simple regex check needed in real app)
	if username == "" || strings.Contains(username, ";") {
		return log.Output(1, "Invalid username") // Simple error
	}

	cmd := exec.Command("useradd", "-m", "-s", "/bin/bash", username)
	return cmd.Run()
}
