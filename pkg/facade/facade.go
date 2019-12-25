package facade

type Facade interface {
	Start() error
}

type cpuer interface {
	SomeCPULogic() error
}

type harddriver interface {
	SomeHardDriveLogic() error
}

type memoryer interface {
	SomeMemoryLogic() error
}

type facade struct {
	cpuer      cpuer
	harddriver harddriver
	memoryer   memoryer
}

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

func NewFacade(
	cpuer cpuer,
	harddriver harddriver,
	memoryer memoryer,
) Facade {
	return &facade{
		cpuer:      cpuer,
		harddriver: harddriver,
		memoryer:   memoryer,
	}
}
