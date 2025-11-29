package auth

import (
	"testing"
	"time"

	"github.com/madr/backend/internal/domain/models"
	refreshTokenDomain "github.com/madr/backend/internal/domain/refreshtoken"
	userDomain "github.com/madr/backend/internal/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of user.Repository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(usr *userDomain.User) error {
	args := m.Called(usr)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(id uint) (*userDomain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*userDomain.User), args.Error(1)
}

func (m *MockUserRepository) GetByUsername(username string) (*userDomain.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*userDomain.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (*userDomain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*userDomain.User), args.Error(1)
}

func (m *MockUserRepository) Update(usr *userDomain.User) error {
	args := m.Called(usr)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateLastLogin(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) ExistsByUsername(username string) (bool, error) {
	args := m.Called(username)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) ExistsByEmail(email string) (bool, error) {
	args := m.Called(email)
	return args.Bool(0), args.Error(1)
}

// MockRefreshTokenRepository is a mock implementation of refreshtoken.Repository
type MockRefreshTokenRepository struct {
	mock.Mock
}

func (m *MockRefreshTokenRepository) Create(rt *refreshTokenDomain.RefreshToken) error {
	args := m.Called(rt)
	return args.Error(0)
}

func (m *MockRefreshTokenRepository) GetByToken(token string) (*refreshTokenDomain.RefreshToken, error) {
	args := m.Called(token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*refreshTokenDomain.RefreshToken), args.Error(1)
}

func (m *MockRefreshTokenRepository) GetByUserID(userID uint) ([]refreshTokenDomain.RefreshToken, error) {
	args := m.Called(userID)
	return args.Get(0).([]refreshTokenDomain.RefreshToken), args.Error(1)
}

func (m *MockRefreshTokenRepository) Revoke(token string) error {
	args := m.Called(token)
	return args.Error(0)
}

func (m *MockRefreshTokenRepository) RevokeAllByUserID(userID uint) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockRefreshTokenRepository) DeleteExpired() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRefreshTokenRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// TestLogin_Success tests successful login
func TestLogin_Success(t *testing.T) {
	// Setup mocks
	mockUserRepo := new(MockUserRepository)
	mockRefreshTokenRepo := new(MockRefreshTokenRepository)

	// Create test user
	testUser := &userDomain.User{
		BaseModel: models.BaseModel{ID: 1},
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", // bcrypt hash for "password123"
		Name:      "Test User",
		Role:      userDomain.RoleUser,
		IsActive:  true,
	}

	// Setup expectations
	mockUserRepo.On("GetByUsername", "testuser").Return(testUser, nil)
	mockUserRepo.On("UpdateLastLogin", uint(1)).Return(nil)
	mockRefreshTokenRepo.On("Create", mock.Anything).Return(nil)

	// Create use case
	useCase := NewUseCase(mockUserRepo, mockRefreshTokenRepo)

	// Test login
	req := &LoginRequest{
		Username: "testuser",
		Password: "password123",
	}

	response, err := useCase.Login(req, "test-agent", "127.0.0.1")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.AccessToken)
	assert.NotEmpty(t, response.RefreshToken)
	assert.Equal(t, "Bearer", response.TokenType)
	assert.Equal(t, testUser.ID, response.User.ID)
	assert.Equal(t, testUser.Username, response.User.Username)

	// Verify all expectations were met
	mockUserRepo.AssertExpectations(t)
	mockRefreshTokenRepo.AssertExpectations(t)
}

// TestLogin_InvalidCredentials tests login with invalid credentials
func TestLogin_InvalidCredentials(t *testing.T) {
	// Setup mocks
	mockUserRepo := new(MockUserRepository)
	mockRefreshTokenRepo := new(MockRefreshTokenRepository)

	// Create test user
	testUser := &userDomain.User{
		BaseModel: models.BaseModel{ID: 1},
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", // bcrypt hash for "password123"
		IsActive:  true,
	}

	// Setup expectations
	mockUserRepo.On("GetByUsername", "testuser").Return(testUser, nil)

	// Create use case
	useCase := NewUseCase(mockUserRepo, mockRefreshTokenRepo)

	// Test login with wrong password
	req := &LoginRequest{
		Username: "testuser",
		Password: "wrongpassword",
	}

	response, err := useCase.Login(req, "test-agent", "127.0.0.1")

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "invalid credentials", err.Error())

	// Verify expectations
	mockUserRepo.AssertExpectations(t)
	mockRefreshTokenRepo.AssertExpectations(t)
}

// TestLogin_InactiveUser tests login with inactive user
func TestLogin_InactiveUser(t *testing.T) {
	// Setup mocks
	mockUserRepo := new(MockUserRepository)
	mockRefreshTokenRepo := new(MockRefreshTokenRepository)

	// Create inactive test user
	testUser := &userDomain.User{
		BaseModel: models.BaseModel{ID: 1},
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy",
		IsActive:  false,
	}

	// Setup expectations
	mockUserRepo.On("GetByUsername", "testuser").Return(testUser, nil)

	// Create use case
	useCase := NewUseCase(mockUserRepo, mockRefreshTokenRepo)

	// Test login
	req := &LoginRequest{
		Username: "testuser",
		Password: "password123",
	}

	response, err := useCase.Login(req, "test-agent", "127.0.0.1")

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "account is inactive", err.Error())

	// Verify expectations
	mockUserRepo.AssertExpectations(t)
	mockRefreshTokenRepo.AssertExpectations(t)
}

// TestRefreshToken_Success tests successful token refresh
func TestRefreshToken_Success(t *testing.T) {
	// Setup mocks
	mockUserRepo := new(MockUserRepository)
	mockRefreshTokenRepo := new(MockRefreshTokenRepository)

	// Create test user
	testUser := &userDomain.User{
		BaseModel: models.BaseModel{ID: 1},
		Username:  "testuser",
		Email:     "test@example.com",
		Role:      userDomain.RoleUser,
		IsActive:  true,
	}

	// Create valid refresh token
	validToken := &refreshTokenDomain.RefreshToken{
		BaseModel: models.BaseModel{ID: 1},
		Token:     "valid-refresh-token",
		UserID:    1,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		IsRevoked: false,
	}

	// Setup expectations
	mockRefreshTokenRepo.On("GetByToken", "valid-refresh-token").Return(validToken, nil)
	mockUserRepo.On("GetByID", uint(1)).Return(testUser, nil)
	mockRefreshTokenRepo.On("Revoke", "valid-refresh-token").Return(nil)
	mockRefreshTokenRepo.On("Create", mock.Anything).Return(nil)

	// Create use case
	useCase := NewUseCase(mockUserRepo, mockRefreshTokenRepo)

	// Test refresh token
	response, err := useCase.RefreshToken("valid-refresh-token")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.AccessToken)
	assert.NotEmpty(t, response.RefreshToken)
	assert.Equal(t, "Bearer", response.TokenType)

	// Verify expectations
	mockUserRepo.AssertExpectations(t)
	mockRefreshTokenRepo.AssertExpectations(t)
}

