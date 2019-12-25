package facade

type HardDriver interface {
	SomeHardDriveLogic() error
}

type hardDrive struct {
	someParamInt int
	someParamStr string
}

func (h *hardDrive) SomeHardDriveLogic() error {
	return nil
}

func NewHardDrive(someParamInt int, someParamStr string) HardDriver {
	return &hardDrive{
		someParamInt: someParamInt,
		someParamStr: someParamStr,
	}
}
