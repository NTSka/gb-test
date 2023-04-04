package main

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"os"
	"test_gb/app/calc"
)

func main() {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetEnvPrefix("development")

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	v.AddConfigPath(dir)

	v.SetConfigName("config")
	err = v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := calc.NewApp(v)

	ctx := context.Background()

	app.Run(ctx)
}
