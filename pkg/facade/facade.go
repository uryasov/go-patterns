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
	err := f.cpuer.SomeCPULogic()
	if err != nil {
		return err
	}
	err = f.harddriver.SomeHardDriveLogic()
	if err != nil {
		return err
	}
	err = f.memoryer.SomeMemoryLogic()
	if err != nil {
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
