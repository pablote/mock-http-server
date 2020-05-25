package src

import (
	"errors"
	"github.com/pablote/mock-server/src/providers"
	"strconv"
	"strings"
)

type MockedResponse struct {
	Status providers.IntProvider
	Body   providers.StringProvider
}

func (s *MockedResponse) SetStatus(statusInput []string) error {
	if strings.Count(statusInput[0], ":") == 2 {
		inputs := strings.Split(statusInput[0], ":")

		probability, err := strconv.Atoi(inputs[0])
		if err != nil {
			return err
		}

		firstStatus, err := strconv.Atoi(inputs[1])
		if err != nil {
			return err
		}

		secondStatus, err := strconv.Atoi(inputs[2])
		if err != nil {
			return err
		}

		s.Status = providers.ProbabilityIntProvider{
			Probability:  probability,
			FirstStatus:  firstStatus,
			SecondStatus: secondStatus,
		}
		return nil
	}

	newStatus, err := strconv.Atoi(statusInput[0])
	if err != nil {
		return err
	}

	s.Status = providers.FixedIntProvider{Value: int(newStatus)}
	return nil
}

func (s *MockedResponse) SetBody(bodyInput []string) error {
	if len(bodyInput[0]) > 1024 {
		return errors.New("body is too large")
	}

	s.Body = providers.FixedStringProvider{Value: bodyInput[0]}
	return nil
}

func NewDefaultMockedResponse() *MockedResponse {
	return &MockedResponse{
		Status: providers.FixedIntProvider{Value: 200},
		Body:   providers.FixedStringProvider{Value: ""},
	}
}
