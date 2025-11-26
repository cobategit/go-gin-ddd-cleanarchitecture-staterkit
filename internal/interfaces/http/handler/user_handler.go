package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/domain/d_user"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/usecase/uc_user"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/pkg/response"
)

// UserHandler berisi handler HTTP untuk user.
type UserHandler struct {
	uc *uc_user.UserUseCase
}

func NewUserHandler(uc *uc_user.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

// @Summary Register user
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body d_user.RegisterRequest true "Register payload"
// @Success 201 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Router /api/v1/auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req d_user.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ToJson(c, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	user, err := h.uc.Register(req.Name, req.Email, req.Password)
	if err != nil {
		response.ToJson(c, http.StatusBadRequest, false, err.Error(), nil)
		return
	}
	response.ToJson(c, http.StatusCreated, true, "user created", user)
}

// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body d_user.LoginRequest true "Login payload"
// @Success 200 {object} response.ApiResponse
// @Failure 401 {object} response.ApiResponse
// @Router /api/v1/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req d_user.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ToJson(c, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	token, err := h.uc.Login(req.Email, req.Password)
	if err != nil {
		if err == uc_user.ErrInvalidCredentials {
			response.ToJson(c, http.StatusUnauthorized, false, "invalid credentials", nil)
		} else {
			response.ToJson(c, http.StatusInternalServerError, false, err.Error(), nil)
		}
		return
	}
	response.ToJson(c, http.StatusOK, true, "login success", gin.H{"token": token})
}

// @Summary Get profile
// @Tags User
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 401 {object} response.ApiResponse
// @Router /api/v1/users/me [get]
// @Security BearerAuth
func (h *UserHandler) Me(c *gin.Context) {
	idVal, exists := c.Get("userID")
	if !exists {
		response.ToJson(c, http.StatusUnauthorized, false, "unauthorized", nil)
		return
	}

	userID, ok := idVal.(int64)
	if !ok {
		response.ToJson(c, http.StatusUnauthorized, false, "invalid user id", nil)
		return
	}

	user, err := h.uc.GetProfile(userID)
	if err != nil {
		response.ToJson(c, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	response.ToJson(c, http.StatusOK, true, "profile", user)
}
