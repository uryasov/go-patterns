package main

import (
	"fmt"

	"github.com/uryasov/go-patterns/pkg/builder"
	"github.com/uryasov/go-patterns/pkg/director"
	"github.com/uryasov/go-patterns/pkg/models"
)

func main() {
	var product models.Product
	newBuilder := builder.NewBuilder(&product)
	newDirector := director.NewDirector(newBuilder)
	newDirector.Construct()
	fmt.Println(product.Content)
}
