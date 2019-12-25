package facade

type Memoryer interface {
	SomeMemoryLogic() error
}

type memory struct {
	someParamInt int
	someParamStr string
}

func (m *memory) SomeMemoryLogic() error {
	return nil
}

func NewMemory(someParamInt int, someParamStr string) Memoryer {
	return &memory{
		someParamInt: someParamInt,
		someParamStr: someParamStr,
	}
}
