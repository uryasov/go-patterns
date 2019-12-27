package builder

import (
	"testing"

	"github.com/uryasov/go-patterns/pkg/models"
)

const (
	basicTest                   = "basicTest"
	expectedBuilderStartAnswer  = "the"
	expectedBuilderMiddleAnswer = "right"
	expectedBuilderEndAnswer    = "order"
)

func Test_builder_BuildPartsBegin(t *testing.T) {
	type fields struct {
		product *models.Product
	}
	type args struct {
		part string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			basicTest,
			fields{new(models.Product)},
			args{expectedBuilderStartAnswer},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder(tt.fields.product)
			b.BuildPartsBegin(tt.args.part)
			if tt.fields.product.Content != expectedBuilderStartAnswer {
				t.Errorf("expected %s in product content, got %s", expectedBuilderStartAnswer, tt.fields.product.Content)
			}
		})
	}
}

func Test_builder_BuildPartsMiddle(t *testing.T) {
	type fields struct {
		product *models.Product
	}
	type args struct {
		part string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			basicTest,
			fields{new(models.Product)},
			args{expectedBuilderMiddleAnswer},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder(tt.fields.product)
			b.BuildPartsMiddle(tt.args.part)
			if tt.fields.product.Content != expectedBuilderMiddleAnswer {
				t.Errorf("expected %s in product content, got %s", expectedBuilderMiddleAnswer, tt.fields.product.Content)
			}
		})
	}
}

func Test_builder_BuildPartsEnd(t *testing.T) {
	type fields struct {
		product *models.Product
	}
	type args struct {
		part string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			basicTest,
			fields{new(models.Product)},
			args{expectedBuilderEndAnswer},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder(tt.fields.product)
			b.BuildPartsEnd(tt.args.part)
			if tt.fields.product.Content != expectedBuilderEndAnswer {
				t.Errorf("expected %s in product content, got %s", expectedBuilderEndAnswer, tt.fields.product.Content)
			}
		})
	}
}
