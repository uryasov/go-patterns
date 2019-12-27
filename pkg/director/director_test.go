package director

import (
	"testing"

	"github.com/uryasov/go-patterns/pkg/models"
)

const (
	basicTest           = "basicTest"
	expectedBasicAnswer = "the right order"
)

type buildertest struct {
	product *models.Product
}

func (b *buildertest) BuildPartsBegin(part string) {
	b.product.Content += part
}
func (b *buildertest) BuildPartsMiddle(part string) {
	b.product.Content += part
}
func (b *buildertest) BuildPartsEnd(part string) {
	b.product.Content += part
}

func NewBuilderTest(product *models.Product) builder {
	return &buildertest{product}
}

func Test_director_Construct(t *testing.T) {
	var product models.Product
	bb := NewBuilderTest(&product)
	type fields struct {
		builder1 builder
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			basicTest,
			fields{bb},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDirector(tt.fields.builder1)
			d.Construct()
			if product.Content != expectedBasicAnswer {
				t.Errorf("expected %s in product content, got %s", expectedBasicAnswer, product.Content)
			}
		})
	}
}
