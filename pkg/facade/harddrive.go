package facade

// HardDriver ...
type HardDriver interface {
	SomeHardDriveLogic() error
}

type hardDrive struct {
	someParamInt int
	someParamStr string
}

// SomeHardDriveLogic ...
func (h *hardDrive) SomeHardDriveLogic() error {
	return nil
}

// NewHardDrive ...
func NewHardDrive(someParamInt int, someParamStr string) HardDriver {
	return &hardDrive{
		someParamInt: someParamInt,
		someParamStr: someParamStr,
	}
}
