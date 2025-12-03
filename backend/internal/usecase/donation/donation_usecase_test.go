package donation

import (
	"testing"

	donationDomain "github.com/madr/backend/internal/domain/donation"
	donationRepo "github.com/madr/backend/internal/repository/donation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDonationRepository is a mock implementation of donation.Repository
type MockDonationRepository struct {
	mock.Mock
}

func (m *MockDonationRepository) Create(don *donationDomain.Donation) error {
	args := m.Called(don)
	return args.Error(0)
}

func (m *MockDonationRepository) GetByID(id uint) (*donationDomain.Donation, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*donationDomain.Donation), args.Error(1)
}

func (m *MockDonationRepository) GetAll(limit, offset int, status *donationDomain.PaymentStatus) ([]donationDomain.Donation, int64, error) {
	args := m.Called(limit, offset, status)
	return args.Get(0).([]donationDomain.Donation), args.Get(1).(int64), args.Error(2)
}

func (m *MockDonationRepository) Update(don *donationDomain.Donation) error {
	args := m.Called(don)
	return args.Error(0)
}

func (m *MockDonationRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockDonationRepository) GetTotalAmount(status *donationDomain.PaymentStatus) (float64, error) {
	args := m.Called(status)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockDonationRepository) GetTotalTransactions(status *donationDomain.PaymentStatus) (int64, error) {
	args := m.Called(status)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDonationRepository) GetAmountPerCategory(status *donationDomain.PaymentStatus) ([]donationRepo.CategoryAmount, error) {
	args := m.Called(status)
	return args.Get(0).([]donationRepo.CategoryAmount), args.Error(1)
}

// TestGetSummary_Success tests successful summary calculation
func TestGetSummary_Success(t *testing.T) {
	// Setup mock
	mockRepo := new(MockDonationRepository)

	successStatus := donationDomain.PaymentStatusSuccess

	// Setup expectations
	mockRepo.On("GetTotalAmount", &successStatus).Return(20000000.0, nil)
	mockRepo.On("GetTotalTransactions", &successStatus).Return(int64(150), nil)
	mockRepo.On("GetAmountPerCategory", &successStatus).Return([]donationRepo.CategoryAmount{
		{
			CategoryID:   1,
			CategoryName: "Pembangunan",
			Amount:       12000000.0,
		},
		{
			CategoryID:   2,
			CategoryName: "Operasional",
			Amount:       6000000.0,
		},
		{
			CategoryID:   3,
			CategoryName: "Sosial",
			Amount:       2000000.0,
		},
	}, nil)

	// Create use case
	useCase := NewUseCase(mockRepo)

	// Test summary
	summary, err := useCase.GetSummary()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, summary)
	assert.Equal(t, 20000000.0, summary.TotalAmount)
	assert.Equal(t, int64(150), summary.TotalTransactions)
	assert.Len(t, summary.PerCategory, 3)
	assert.Equal(t, "Pembangunan", summary.PerCategory[0].CategoryName)
	assert.Equal(t, 12000000.0, summary.PerCategory[0].Amount)
	assert.Equal(t, "Operasional", summary.PerCategory[1].CategoryName)
	assert.Equal(t, 6000000.0, summary.PerCategory[1].Amount)

	// Verify all expectations were met
	mockRepo.AssertExpectations(t)
}

// TestGetSummary_EmptyData tests summary with no donations
func TestGetSummary_EmptyData(t *testing.T) {
	// Setup mock
	mockRepo := new(MockDonationRepository)

	successStatus := donationDomain.PaymentStatusSuccess

	// Setup expectations
	mockRepo.On("GetTotalAmount", &successStatus).Return(0.0, nil)
	mockRepo.On("GetTotalTransactions", &successStatus).Return(int64(0), nil)
	mockRepo.On("GetAmountPerCategory", &successStatus).Return([]donationRepo.CategoryAmount{}, nil)

	// Create use case
	useCase := NewUseCase(mockRepo)

	// Test summary
	summary, err := useCase.GetSummary()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, summary)
	assert.Equal(t, 0.0, summary.TotalAmount)
	assert.Equal(t, int64(0), summary.TotalTransactions)
	assert.Len(t, summary.PerCategory, 0)

	// Verify all expectations were met
	mockRepo.AssertExpectations(t)
}

// TestGetSummary_ErrorHandling tests error handling in summary
func TestGetSummary_ErrorHandling(t *testing.T) {
	// Setup mock
	mockRepo := new(MockDonationRepository)

	successStatus := donationDomain.PaymentStatusSuccess

	// Setup expectations - simulate error
	mockRepo.On("GetTotalAmount", &successStatus).Return(0.0, assert.AnError)

	// Create use case
	useCase := NewUseCase(mockRepo)

	// Test summary
	summary, err := useCase.GetSummary()

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, summary)
	assert.Equal(t, "failed to calculate total amount", err.Error())

	// Verify expectations
	mockRepo.AssertExpectations(t)
}

