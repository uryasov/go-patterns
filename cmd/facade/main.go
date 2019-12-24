package main

import (
	"log"

	"github.com/uryasov/go-patterns/pkg/facade"
)

func main() {
	cpu := facade.NewCPU(1, "1")
	hardDrive := facade.NewHardDrive(2, "2")
	memory := facade.NewMemory(3, "3")
	facade := facade.NewFacade(cpu, hardDrive, memory)
	err := facade.Start()
	if err != nil {
		log.Panic(err)
	}
}
