package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"jopanel/backend/models"
	"jopanel/backend/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

type CreateUserInput struct {
	Username string      `json:"username" binding:"required"`
	Email    string      `json:"email" binding:"required,email"`
	Password string      `json:"password" binding:"required,min=6"`
	Role     models.Role `json:"role"`
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Role:     input.Role,
	}
	if user.Role == "" {
		user.Role = models.RoleUser
	}

	createdUser, err := ctrl.userService.CreateUser(user, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ctrl.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) SuspendUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.userService.ChangeStatus(uint(id), "suspended"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to suspend user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User suspended"})
}

func (ctrl *UserController) UnsuspendUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.userService.ChangeStatus(uint(id), "active"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unsuspend user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User activated"})
}
