package calc

import (
	"emperror.dev/errors"
	"github.com/spf13/viper"
	"test_gb/src/clients/eth"
	"test_gb/src/svc/calc"
)

func buildETHClient(v *viper.Viper) (*eth.Client, error) {
	config, err := eth.NewConfig(v)
	if err != nil {
		return nil, errors.Wrap(err, "eth.NewConfig")
	}

	return eth.NewClient(config)
}

func buildCalcSvc(v *viper.Viper, client calc.IClient) (calc.IService, error) {
	config, err := calc.NewConfig(v)
	if err != nil {
		return nil, errors.Wrap(err, "calc.NewConfig")
	}

	return calc.NewService(config, client), nil
}
