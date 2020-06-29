package mocks

import (
	"github.com/gh0stl1m/goblog/accountservice/domains"
	"github.com/stretchr/testify/mock"
)

// MockAccountUseCases is a genertated mock type
type MockAccountUseCases struct {
	mock.Mock
}

// GetByID it's a mock
func (m *MockAccountUseCases) GetByID(accountID string) (domains.Account, error) {
	args := m.Called(accountID)
	return args.Get(0).(domains.Account), args.Error(1)
}
