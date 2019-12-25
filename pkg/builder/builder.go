package builder

// Builderer ...
type Builderer interface {
	buildPartsBegin() string
	buildPartsMiddle() string
	buildPartsEnd() string
}

type builder struct {
}

func (b *builder) buildPartsBegin() string {
	return "the"
}
func (b *builder) buildPartsMiddle() string {
	return "right"
}
func (b *builder) buildPartsEnd() string {
	return "order"
}

// NewBuilder ...
func NewBuilder() Builderer {
	return &builder{}
}
