package main

import (
	"fmt"

	"github.com/uryasov/go-patterns/pkg/builder"
)

func main() {
	newBuilder := builder.NewBuilder()
	newDirector := builder.NewDirector(newBuilder)
	product := newDirector.Construct()
	fmt.Println(product)
}
