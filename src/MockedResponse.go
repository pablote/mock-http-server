package src

import (
	"strconv"
)

type MockedResponse struct {
	Status int
}

func (s *MockedResponse) SetStatus(statusInput []string) error {
	newStatus, err := strconv.ParseInt(statusInput[0], 0, 32)
	if err != nil {
		return err
	}

	s.Status = int(newStatus)
	return nil
}

func NewDefaultMockedResponse() *MockedResponse {
	return &MockedResponse{
		Status:200,
	}
}
