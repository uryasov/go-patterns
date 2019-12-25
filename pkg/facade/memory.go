package facade

// Memoryer ...
type Memoryer interface {
	SomeMemoryLogic() error
}

type memory struct {
	someParamInt int
	someParamStr string
}

// SomeMemoryLogic ...
func (m *memory) SomeMemoryLogic() error {
	return nil
}

// NewMemory ...
func NewMemory(someParamInt int, someParamStr string) Memoryer {
	return &memory{
		someParamInt: someParamInt,
		someParamStr: someParamStr,
	}
}
