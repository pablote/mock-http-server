package providers

import "math/rand"

type ProbabilityIntProvider struct {
	Probability int
	FirstStatus int
	SecondStatus int
}

func (p ProbabilityIntProvider) Get() int {
	r := rand.Intn(100)
	if r <= p.Probability {
		return p.FirstStatus
	} else {
		return p.SecondStatus
	}
}

