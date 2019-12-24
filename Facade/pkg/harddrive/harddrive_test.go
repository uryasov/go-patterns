package harddrive

import "testing"

func Test_hardDrive_SomeHardDriveLogic(t *testing.T) {
	type fields struct {
		someParamInt int
		someParamStr string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"basicTest",
			fields{
				someParamInt: 1,
				someParamStr: "1",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hardDrive{
				someParamInt: tt.fields.someParamInt,
				someParamStr: tt.fields.someParamStr,
			}
			if err := h.SomeHardDriveLogic(); (err != nil) != tt.wantErr {
				t.Errorf("hardDrive.SomeHardDriveLogic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
