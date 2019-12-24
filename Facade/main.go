package main

import (
	"log"

	"github.com/uryasov/go-patterns/Facade/pkg/cpu"
	"github.com/uryasov/go-patterns/Facade/pkg/facade"
	"github.com/uryasov/go-patterns/Facade/pkg/harddrive"
	"github.com/uryasov/go-patterns/Facade/pkg/memory"
)

func main() {
	cpu := cpu.NewCPU(1, "1")
	hardDrive := harddrive.NewHardDrive(2, "2")
	memory := memory.NewMemory(3, "3")
	facade := facade.NewFacade(cpu, hardDrive, memory)
	err := facade.Start()
	if err != nil {
		log.Panic(err)
	}
}
