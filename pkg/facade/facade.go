package facade

type cpuer interface {
	SomeCPULogic() error
}

type harddriver interface {
	SomeHardDriveLogic() error
}

type memoryer interface {
	SomeMemoryLogic() error
}

// Facader ...
type Facader interface {
	Start() error
}

type facade struct {
	cpuer      cpuer
	harddriver harddriver
	memoryer   memoryer
}

// Start ...
func (f *facade) Start() error {
	if err := f.cpuer.SomeCPULogic(); err != nil {
		return err
	}
	if err := f.harddriver.SomeHardDriveLogic(); err != nil {
		return err
	}
	if err := f.memoryer.SomeMemoryLogic(); err != nil {
		return err
	}
	return nil
}

// NewFacade ...
func NewFacade(
	cpuer cpuer,
	harddriver harddriver,
	memoryer memoryer,
) Facader {
	return &facade{
		cpuer:      cpuer,
		harddriver: harddriver,
		memoryer:   memoryer,
	}
}
