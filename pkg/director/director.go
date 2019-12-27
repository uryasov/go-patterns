package director

type builder interface {
	BuildPartsBegin(string)
	BuildPartsMiddle(string)
	BuildPartsEnd(string)
}

// Directorer ...
type Directorer interface {
	Construct()
}

type director struct {
	builder1 builder
}

// Construct ...
func (d *director) Construct() {
	d.builder1.BuildPartsBegin("the ")
	d.builder1.BuildPartsMiddle("right ")
	d.builder1.BuildPartsEnd("order")
}

// NewDirector ...
func NewDirector(builder1 builder) Directorer {
	return &director{builder1}
}
