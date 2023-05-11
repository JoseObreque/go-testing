package storage

import "github.com/stretchr/testify/mock"

// StorageMock is a mock implementation of Storage interface.
type StorageMock struct {
	Mock mock.Mock
}

// GetValue is a mock implementation of Storage.GetValue
func (s *StorageMock) GetValue(key string) interface{} {
	args := s.Mock.Called(key)
	return args.Get(0)
}

// NewStorageMock returns a new mock instance of Storage.
func NewStorageMock() *StorageMock {
	return &StorageMock{}
}
