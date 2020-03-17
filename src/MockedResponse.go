package src

type MockedResponse struct {
	Status int
}

func (s *MockedResponse) SetStatus(status int) {
	s.Status = status
}

func NewDefaultMockedResponse() *MockedResponse {
	return &MockedResponse{
		Status:200,
	}
}
