package facade

// CPUer cpu interface
type CPUer interface {
	SomeCPULogic() error
}

type cpu struct {
	someParamInt int
	someParamStr string
}

// SomeCPULogic does something
func (c *cpu) SomeCPULogic() error {
	return nil
}

// NewCPU new cpu
func NewCPU(someParamInt int, someParamStr string) CPUer {
	return &cpu{
		someParamInt: someParamInt,
		someParamStr: someParamStr,
	}
}
