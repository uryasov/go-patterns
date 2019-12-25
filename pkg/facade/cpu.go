package facade

type CPU interface {
	SomeCPULogic() error
}

type cpu struct {
	someParamInt int
	someParamStr string
}

func (c *cpu) SomeCPULogic() error {
	return nil
}

func NewCPU(someParamInt int, someParamStr string) CPU {
	return &cpu{
		someParamInt: someParamInt,
		someParamStr: someParamStr,
	}
}
