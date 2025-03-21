package handlers

import (
	"cooking/backend/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Chef     bool   `json:"chef"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type AuthHandler struct {
	UserUseCase *usecase.UserUseCase
	ChefUseCase *usecase.ChefUseCase
}

type ChefResponse struct {
	Username       string  `json:"username"`
	FollowersCount int     `json:"followers"`
	Bio            *string `json:"bio,omitempty"`
	Avatar         *string `json:"avatar,omitempty"`
}

type UserResponse struct {
	Username string  `json:"username"`
	Avatar   *string `json:"avatar,omitempty"`
}

func NewAuthHandler(userUseCase *usecase.UserUseCase, chefUseCase *usecase.ChefUseCase) *AuthHandler {
	return &AuthHandler{UserUseCase: userUseCase, ChefUseCase: chefUseCase}
}

func (h *AuthHandler) Register(ctx echo.Context) error {
	var req AuthRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid request data")
	}

	var token string
	var err error
	if req.Chef {
		token, err = h.ChefUseCase.RegisterChef(ctx.Request().Context(), req.Username, req.Password)
	} else {
		token, err = h.UserUseCase.RegisterUser(ctx.Request().Context(), req.Username, req.Password)
	}

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, AuthResponse{Token: token})
}

func (h *AuthHandler) Authenticate(ctx echo.Context) error {
	var req AuthRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid request data")
	}

	var token string
	var err error
	if req.Chef {
		token, err = h.ChefUseCase.AuthenticateChef(ctx.Request().Context(), req.Username, req.Password)
	} else {
		token, err = h.UserUseCase.AuthenticateUser(ctx.Request().Context(), req.Username, req.Password)
	}

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	return ctx.JSON(http.StatusOK, AuthResponse{Token: token})
}

func (h *AuthHandler) ShowChefProfile(ctx echo.Context) error {
	username := ctx.Get("username").(string)
	chefDetails, err := h.ChefUseCase.GetChefInfo(ctx.Request().Context(), username)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "profile not found")
	}
	response := ChefResponse{
		Username:       chefDetails.Name,
		FollowersCount: chefDetails.FollowersCount,
		Bio:            chefDetails.Bio,
		Avatar:         chefDetails.Avatar,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *AuthHandler) ShowUserProfile(ctx echo.Context) error {
	username := ctx.Get("username").(string)
	userDetails, err := h.UserUseCase.GetUserInfo(ctx.Request().Context(), username)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "profile not found")
	}
	response := UserResponse{
		Username: userDetails.Name,
		Avatar:   userDetails.Avatar,
	}

	return ctx.JSON(http.StatusOK, response)
}
