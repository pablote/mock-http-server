package providers

type FixedStringProvider struct {
	Value string
}

func (p FixedStringProvider) Get() string {
	return p.Value
}

