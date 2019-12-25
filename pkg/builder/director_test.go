package builder

import "testing"

const (
	basicTest                   = "basicTest"
	expectedBuilderStartAnswer  = "the"
	expectedBuilderMiddleAnswer = "right"
	expectedBuilderEndAnswer    = "order"
	expectedBasicAnswer         = "the right order"
)

// using builder from builder.go because it isn't import,
// if it was in different package would create mock builderer implementation
func makeBuilder() builderer {
	return &builder{}
}

func Test_director_Construct(t *testing.T) {
	type fields struct {
		builder1 builderer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			basicTest,
			fields{
				makeBuilder(),
			},
			expectedBasicAnswer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &director{
				builder1: tt.fields.builder1,
			}
			if got := d.Construct(); got != tt.want {
				t.Errorf("director.Construct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builder_buildPartsBegin(t *testing.T) {
	tests := []struct {
		name string
		b    *builder
		want string
	}{
		{
			basicTest,
			&builder{},
			expectedBuilderStartAnswer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &builder{}
			if got := b.buildPartsBegin(); got != tt.want {
				t.Errorf("builder.buildPartsBegin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builder_buildPartsMiddle(t *testing.T) {
	tests := []struct {
		name string
		b    *builder
		want string
	}{
		{
			basicTest,
			&builder{},
			expectedBuilderMiddleAnswer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &builder{}
			if got := b.buildPartsMiddle(); got != tt.want {
				t.Errorf("builder.buildPartsMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builder_buildPartsEnd(t *testing.T) {
	tests := []struct {
		name string
		b    *builder
		want string
	}{
		{
			basicTest,
			&builder{},
			expectedBuilderEndAnswer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &builder{}
			if got := b.buildPartsEnd(); got != tt.want {
				t.Errorf("builder.buildPartsEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
