package auth

import (
	"errors"
	"time"

	refreshTokenDomain "github.com/madr/backend/internal/domain/refreshtoken"
	userDomain "github.com/madr/backend/internal/domain/user"
	refreshTokenRepo "github.com/madr/backend/internal/repository/refreshtoken"
	userRepo "github.com/madr/backend/internal/repository/user"
	"github.com/madr/backend/pkg/bcrypt"
	"github.com/madr/backend/pkg/jwt"
	"github.com/madr/backend/pkg/logger"
)

// UseCase defines the interface for auth use case
type UseCase interface {
	Register(req *RegisterRequest) (*RegisterResponse, error)
	Login(req *LoginRequest, userAgent, ipAddress string) (*LoginResponse, error)
	RefreshToken(refreshToken string) (*RefreshTokenResponse, error)
	GetMe(userID uint) (*MeResponse, error)
	Logout(refreshToken string) error
	LogoutAll(userID uint) error
}

// RegisterRequest represents the request to register a new user
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"max=255"`
	Role     string `json:"role,omitempty"` // Optional, defaults to "user"
}

// RegisterResponse represents the response after registration
type RegisterResponse struct {
	User *userDomain.User `json:"user"`
}

// LoginRequest represents the request to login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the response after login
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	User         *userDomain.User `json:"user"`
}

// RefreshTokenRequest represents the request to refresh token
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// RefreshTokenResponse represents the response after refreshing token
type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

// MeResponse represents the current user information
type MeResponse struct {
	User *userDomain.User `json:"user"`
}

type useCase struct {
	userRepo         userRepo.Repository
	refreshTokenRepo refreshTokenRepo.Repository
}

// NewUseCase creates a new auth use case
func NewUseCase(userRepoInstance userRepo.Repository, refreshTokenRepoInstance refreshTokenRepo.Repository) UseCase {
	return &useCase{
		userRepo:         userRepoInstance,
		refreshTokenRepo: refreshTokenRepoInstance,
	}
}

// Register registers a new user
func (uc *useCase) Register(req *RegisterRequest) (*RegisterResponse, error) {
	// Check if username already exists
	exists, err := uc.userRepo.ExistsByUsername(req.Username)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to check username existence")
		return nil, errors.New("failed to check username")
	}
	if exists {
		return nil, errors.New("username already exists")
	}

	// Check if email already exists
	exists, err = uc.userRepo.ExistsByEmail(req.Email)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to check email existence")
		return nil, errors.New("failed to check email")
	}
	if exists {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to hash password")
		return nil, errors.New("failed to process password")
	}

	// Determine role
	userRole := userDomain.RoleUser
	if req.Role == "admin" {
		userRole = userDomain.RoleAdmin
	}

	// Create user
	newUser := &userDomain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
		Role:     userRole,
		IsActive: true,
	}

	if err := uc.userRepo.Create(newUser); err != nil {
		logger.Error().Err(err).Msg("Failed to create user")
		return nil, errors.New("failed to create user")
	}

	logger.Info().
		Uint("user_id", newUser.ID).
		Str("username", newUser.Username).
		Msg("User registered successfully")

	return &RegisterResponse{
		User: newUser,
	}, nil
}

// Login authenticates a user and returns tokens
func (uc *useCase) Login(req *LoginRequest, userAgent, ipAddress string) (*LoginResponse, error) {
	// Get user by username
	usr, err := uc.userRepo.GetByUsername(req.Username)
	if err != nil {
		logger.Warn().Str("username", req.Username).Msg("Login attempt with invalid username")
		return nil, errors.New("invalid credentials")
	}

	// Check if user is active
	if !usr.IsActive {
		logger.Warn().Uint("user_id", usr.ID).Msg("Login attempt for inactive user")
		return nil, errors.New("account is inactive")
	}

	// Verify password
	if !bcrypt.CheckPasswordHash(req.Password, usr.Password) {
		logger.Warn().Uint("user_id", usr.ID).Msg("Login attempt with invalid password")
		return nil, errors.New("invalid credentials")
	}

	// Generate access token
	accessToken, err := jwt.GenerateAccessToken(usr.ID, usr.Username, string(usr.Role))
	if err != nil {
		logger.Error().Err(err).Uint("user_id", usr.ID).Msg("Failed to generate access token")
		return nil, errors.New("failed to generate token")
	}

	// Generate refresh token
	refreshTokenString, err := jwt.GenerateRefreshToken()
	if err != nil {
		logger.Error().Err(err).Uint("user_id", usr.ID).Msg("Failed to generate refresh token")
		return nil, errors.New("failed to generate refresh token")
	}

	// Store refresh token in database
	refreshToken := &refreshTokenDomain.RefreshToken{
		Token:     refreshTokenString,
		UserID:    usr.ID,
		ExpiresAt: time.Now().Add(jwt.GetRefreshTokenExpiry()),
		IsRevoked: false,
		UserAgent: userAgent,
		IPAddress: ipAddress,
	}

	if err := uc.refreshTokenRepo.Create(refreshToken); err != nil {
		logger.Error().Err(err).Uint("user_id", usr.ID).Msg("Failed to store refresh token")
		return nil, errors.New("failed to store refresh token")
	}

	// Update last login
	if err := uc.userRepo.UpdateLastLogin(usr.ID); err != nil {
		logger.Warn().Err(err).Uint("user_id", usr.ID).Msg("Failed to update last login")
		// Don't fail the login if this fails
	}

	logger.Info().
		Uint("user_id", usr.ID).
		Str("username", usr.Username).
		Msg("User logged in successfully")

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenString,
		TokenType:    "Bearer",
		ExpiresIn:    int64(jwt.GetAccessTokenExpiry().Seconds()),
		User:         usr,
	}, nil
}

// RefreshToken refreshes an access token using a refresh token
func (uc *useCase) RefreshToken(refreshTokenString string) (*RefreshTokenResponse, error) {
	// Get refresh token from database
	rt, err := uc.refreshTokenRepo.GetByToken(refreshTokenString)
	if err != nil {
		logger.Warn().Msg("Invalid refresh token")
		return nil, errors.New("invalid refresh token")
	}

	// Check if token is valid
	if !rt.IsValid() {
		logger.Warn().Uint("token_id", rt.ID).Msg("Refresh token is expired or revoked")
		return nil, errors.New("refresh token is expired or revoked")
	}

	// Get user
	usr, err := uc.userRepo.GetByID(rt.UserID)
	if err != nil {
		logger.Error().Err(err).Uint("user_id", rt.UserID).Msg("Failed to get user for refresh token")
		return nil, errors.New("user not found")
	}

	// Check if user is active
	if !usr.IsActive {
		logger.Warn().Uint("user_id", usr.ID).Msg("Refresh token attempt for inactive user")
		return nil, errors.New("account is inactive")
	}

	// Generate new access token
	accessToken, err := jwt.GenerateAccessToken(usr.ID, usr.Username, string(usr.Role))
	if err != nil {
		logger.Error().Err(err).Uint("user_id", usr.ID).Msg("Failed to generate access token")
		return nil, errors.New("failed to generate token")
	}

	// Optionally generate new refresh token (rotate refresh token)
	newRefreshTokenString, err := jwt.GenerateRefreshToken()
	if err != nil {
		logger.Error().Err(err).Uint("user_id", usr.ID).Msg("Failed to generate new refresh token")
		return nil, errors.New("failed to generate refresh token")
	}

	// Revoke old refresh token
	if err := uc.refreshTokenRepo.Revoke(rt.Token); err != nil {
		logger.Warn().Err(err).Uint("token_id", rt.ID).Msg("Failed to revoke old refresh token")
		// Continue anyway
	}

	// Store new refresh token
	newRefreshToken := &refreshTokenDomain.RefreshToken{
		Token:     newRefreshTokenString,
		UserID:    usr.ID,
		ExpiresAt: time.Now().Add(jwt.GetRefreshTokenExpiry()),
		IsRevoked: false,
		UserAgent: rt.UserAgent,
		IPAddress: rt.IPAddress,
	}

	if err := uc.refreshTokenRepo.Create(newRefreshToken); err != nil {
		logger.Error().Err(err).Uint("user_id", usr.ID).Msg("Failed to store new refresh token")
		return nil, errors.New("failed to store refresh token")
	}

	logger.Info().
		Uint("user_id", usr.ID).
		Msg("Token refreshed successfully")

	return &RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshTokenString,
		TokenType:    "Bearer",
		ExpiresIn:    int64(jwt.GetAccessTokenExpiry().Seconds()),
	}, nil
}

// GetMe returns the current user information
func (uc *useCase) GetMe(userID uint) (*MeResponse, error) {
	usr, err := uc.userRepo.GetByID(userID)
	if err != nil {
		logger.Error().Err(err).Uint("user_id", userID).Msg("Failed to get user")
		return nil, errors.New("user not found")
	}

	return &MeResponse{
		User: usr,
	}, nil
}

// Logout revokes a refresh token
func (uc *useCase) Logout(refreshToken string) error {
	if err := uc.refreshTokenRepo.Revoke(refreshToken); err != nil {
		logger.Error().Err(err).Msg("Failed to revoke refresh token")
		return errors.New("failed to logout")
	}

	logger.Info().Msg("User logged out successfully")
	return nil
}

// LogoutAll revokes all refresh tokens for a user
func (uc *useCase) LogoutAll(userID uint) error {
	if err := uc.refreshTokenRepo.RevokeAllByUserID(userID); err != nil {
		logger.Error().Err(err).Uint("user_id", userID).Msg("Failed to revoke all refresh tokens")
		return errors.New("failed to logout from all devices")
	}

	logger.Info().Uint("user_id", userID).Msg("User logged out from all devices")
	return nil
}

