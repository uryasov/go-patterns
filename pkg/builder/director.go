package builder

type builderer interface {
	buildPartsBegin() string
	buildPartsMiddle() string
	buildPartsEnd() string
}

type Directorer interface {
	Construct() string
}

type director struct {
	builder1 builderer
}

func (d *director) Construct() string {
	finishedProduct := d.builder1.buildPartsBegin() + " " + d.builder1.buildPartsMiddle() + " " + d.builder1.buildPartsEnd()
	return finishedProduct
}

func NewDirector(builder1 builderer) Directorer {
	return &director{builder1}
}
