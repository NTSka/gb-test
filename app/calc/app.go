package calc

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"test_gb/src/svc/calc"
)

type App struct {
	svc calc.IService
}

func (a *App) Run(ctx context.Context) {
	address, err := a.svc.GetMaxDiff(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Maximum absolute volume on address: %s\n", address)
}

func NewApp(v *viper.Viper) *App {
	client, err := buildETHClient(v)
	if err != nil {
		log.Fatal(err)
	}

	svc, err := buildCalcSvc(v, client)
	if err != nil {
		log.Fatal(err)
	}

	return &App{
		svc: svc,
	}
}
