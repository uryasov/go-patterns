package harddrive

type HardDrive interface {
	SomeHardDriveLogic() error
}

type hardDrive struct {
	someParamInt int
	someParamStr string
}

func (h *hardDrive) SomeHardDriveLogic() error {
	return nil
}

func NewHardDrive(
	someParamInt int,
	someParamStr string,
) HardDrive {
	return &hardDrive{
		someParamInt: someParamInt,
		someParamStr: someParamStr,
	}
}
