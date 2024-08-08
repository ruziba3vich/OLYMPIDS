package auth

import (
	"encoding/json"
	"log/slog"

	"github.com/gin-gonic/gin"
	pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/auth"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/msgbroker/auth"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/redisservice"
)

type (
	AuthHandler struct {
		auth      pb.AuthServiceClient
		logger    *slog.Logger
		redis     *redisservice.RedisService
		msgbroker *auth.AuthMsgBroker
	}
)

func NewAthleteHandler(logger *slog.Logger, auth pb.AuthServiceClient, redis *redisservice.RedisService, msgbroker *auth.AuthMsgBroker) *AuthHandler {
	return &AuthHandler{
		auth:      auth,
		logger:    logger,
		redis:     redis,
		msgbroker: msgbroker,
	}
}

// RegisterHandler godoc
// @Summary Register a new user
// @Description Register a new user with an email and password
// @Tags User Auth
// @Accept json
// @Produce json
// @Param request body pb.RegisterRequest true "Register Request"
// @Success 201 {object} pb.RegisterResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/user/register [post]
func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	h.logger.Info("RegisterHandler called")
	var req pb.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	body, err := json.Marshal(&req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := h.msgbroker.Register(body); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.auth.GetUserByEmail(c.Request.Context(), &pb.GetUserByEmailRequest{Email: req.Email})
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(201, resp)
}

// LoginHandler godoc
// @Summary User login
// @Description Log in a user with email and password
// @Tags User Auth
// @Accept json
// @Produce json
// @Param request body pb.LoginRequest true "Login Request"
// @Success 200 {object} pb.LoginResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/user/login [post]
func (h *AuthHandler) LoginHandler(c *gin.Context) {
	h.logger.Info("LoginHandler called")
	var req pb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.auth.Login(c.Request.Context(), &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, resp)
}

// LogoutHandler godoc
// @Summary User logout
// @Description Log out a user by their ID
// @Tags User Auth
// @Accept json
// @Produce json
// @Param request body pb.LogoutRequest true "Logout Request"
// @Success 200 {object} pb.LogoutResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/user/logout [post]
func (h *AuthHandler) LogoutHandler(c *gin.Context) {
	h.logger.Info("LogoutHandler called")
	var req pb.LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.auth.Logout(c.Request.Context(), &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, resp)
}

// AdminLoginHandler godoc
// @Summary Admin login
// @Description Log in an admin user with email and password
// @Tags Admin Auth
// @Accept json
// @Produce json
// @Param request body pb.LoginRequest true "Login Request"
// @Success 200 {object} pb.LoginResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/admin/login [post]
func (h *AuthHandler) AdminLoginHandler(c *gin.Context) {
	h.logger.Info("AdminLoginHandler called")
	var req pb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.auth.Login(c.Request.Context(), &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, resp)
}

// AdminLogoutHandler godoc
// @Summary Admin logout
// @Description Log out an admin user by their ID
// @Tags Admin Auth
// @Accept json
// @Produce json
// @Param request body pb.LogoutRequest true "Logout Request"
// @Success 200 {object} pb.LogoutResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/admin/logout [post]
func (h *AuthHandler) AdminLogoutHandler(c *gin.Context) {
	h.logger.Info("AdminLogoutHandler called")
	var req pb.LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.auth.Logout(c.Request.Context(), &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, resp)
}

// UpdateUserHandler godoc
// @Summary Update user
// @Description Update a user's information by admin
// @Tags Admin Auth
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body pb.UpdateUserRequest true "Update User Request"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 403 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /admin/auth/update/{id} [put]
func (h *AuthHandler) UpdateUserHandler(c *gin.Context) {
	h.logger.Info("UpdateUserHandler called")
	var req pb.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	body, err := json.Marshal(&req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := h.msgbroker.UpdateUser(body); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "User updated successfully"})
}

// DeleteUserHandler godoc
// @Summary Delete user
// @Description Soft delete a user by admin
// @Tags Admin Auth
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body pb.DeleteUserRequest true "Delete User Request"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 403 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /admin/auth/delete/{id} [delete]
func (h *AuthHandler) DeleteUserHandler(c *gin.Context) {
	h.logger.Info("DeleteUserHandler called")
	var req pb.DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	body, err := json.Marshal(&req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := h.msgbroker.DeleteUser(body); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "User deleted successfully"})
}

// @Summary Super Admin Login
// @Description Login as a super admin
// @Tags Super Admin
// @Accept json
// @Produce json
// @Param loginRequest body pb.LoginRequest true "Login Request"
// @Success 200 {object} pb.LoginResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /superadmin/login [post]
func (h *AuthHandler) SuperAdminLoginHandler(c *gin.Context) {
	h.logger.Info("SuperAdminLoginHandler called")
	var req pb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.auth.Login(c.Request.Context(), &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary Super Admin Logout
// @Description Logout from a super admin session
// @Tags Super Admin
// @Accept json
// @Produce json
// @Param logoutRequest body pb.LogoutRequest true "Logout Request"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /superadmin/logout [post]
func (h *AuthHandler) SuperAdminLogoutHandler(c *gin.Context) {
	h.logger.Info("SuperAdminLogoutHandler called")
	var req pb.LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.auth.Logout(c.Request.Context(), &req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary Create Admin
// @Description Create a new admin user by a super admin
// @Tags Super Admin
// @Accept json
// @Produce json
// @Param createAdminRequest body pb.CreateAdminRequest true "Create Admin Request"
// @Success 201 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /superadmin/createadmin [post]
func (h *AuthHandler) SuperAdminCreateAdminHandler(c *gin.Context) {
	h.logger.Info("SuperAdminCreateAdminHandler called")
	var req pb.CreateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	body, err := json.Marshal(&req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := h.msgbroker.CreateAdmin(body); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(201, gin.H{"message": "Admin created successfully"})
}
