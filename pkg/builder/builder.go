package builder

import "github.com/uryasov/go-patterns/pkg/models"

// Builder ...
type Builder interface {
	BuildPartsBegin(string)
	BuildPartsMiddle(string)
	BuildPartsEnd(string)
}

type builder struct {
	product *models.Product
}

// BuildPartsBegin ...
func (b *builder) BuildPartsBegin(part string) {
	b.product.Content += part
} 
// BuildPartsMiddle ...
func (b *builder) BuildPartsMiddle(part string) {
	b.product.Content += part
} 
// BuildPartsEnd ...
func (b *builder) BuildPartsEnd(part string) {
	b.product.Content += part
}

// NewBuilder ...
func NewBuilder(product *models.Product) Builder {
	return &builder{product}
}
