package providers

type FixedIntProvider struct {
	Value int
}

func (p FixedIntProvider) Get() int {
	return p.Value
}

