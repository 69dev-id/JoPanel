package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"jopanel/backend/services"
)

type FileController struct {
	fileService services.FileService
}

func NewFileController(fileService services.FileService) *FileController {
	return &FileController{fileService: fileService}
}

// Helper to get user home directory (mocked for now)
func getUserHome(c *gin.Context) string {
	// In real app, fetch user from DB or JWT claims and return /home/username
	// For dev, we use a temp dir or local path
	userHome := os.Getenv("DEV_USER_HOME")
	if userHome == "" {
		return "./tmp_home"
	}
	return userHome
}

func (ctrl *FileController) ListFiles(c *gin.Context) {
	path := c.DefaultQuery("path", "/")
	files, err := ctrl.fileService.ListFiles(getUserHome(c), path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}

func (ctrl *FileController) GetContent(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path required"})
		return
	}
	content, err := ctrl.fileService.ReadFile(getUserHome(c), path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"content": content})
}

func (ctrl *FileController) Upload(c *gin.Context) {
	path := c.PostForm("path") // Directory to upload to
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if err := ctrl.fileService.UploadFile(getUserHome(c), path, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded"})
}

func (ctrl *FileController) Mkdir(c *gin.Context) {
	var input struct {
		Path string `json:"path" binding:"required"` // Full relative path of new folder
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.fileService.CreateDirectory(getUserHome(c), input.Path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Directory created"})
}

func (ctrl *FileController) Delete(c *gin.Context) {
	path := c.Query("path")
	if err := ctrl.fileService.DeletePath(getUserHome(c), path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
