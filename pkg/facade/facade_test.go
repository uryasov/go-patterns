package facade

import "testing"

func Test_cpu_SomeCPULogic(t *testing.T) {
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
			c := &cpu{
				someParamInt: tt.fields.someParamInt,
				someParamStr: tt.fields.someParamStr,
			}
			if err := c.SomeCPULogic(); (err != nil) != tt.wantErr {
				t.Errorf("cpu.SomeCPULogic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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

func Test_memory_SomeMemoryLogic(t *testing.T) {
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
			m := &memory{
				someParamInt: tt.fields.someParamInt,
				someParamStr: tt.fields.someParamStr,
			}
			if err := m.SomeMemoryLogic(); (err != nil) != tt.wantErr {
				t.Errorf("memory.SomeMemoryLogic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
