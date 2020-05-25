package src

import (
	"github.com/pablote/mock-server/src/providers"
	"strconv"
)

type MockedResponse struct {
	Status providers.IntProvider
}

func (s *MockedResponse) SetStatus(statusInput []string) error {
	newStatus, err := strconv.ParseInt(statusInput[0], 0, 32)
	if err != nil {
		return err
	}

	s.Status = providers.FixedIntProvider{Value: int(newStatus)}
	return nil
}

func NewDefaultMockedResponse() *MockedResponse {
	return &MockedResponse{
		Status: providers.FixedIntProvider{Value: 200},
	}
}
