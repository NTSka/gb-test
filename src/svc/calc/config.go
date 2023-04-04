package calc

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

type Config struct {
	Parallel int
	Count    int64
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Parallel, validation.Required),
		validation.Field(&c.Count, validation.Required))
}

func NewConfig(v *viper.Viper) (*Config, error) {
	c := &Config{
		Parallel: v.GetInt("calc.parallel"),
		Count:    v.GetInt64("calc.count"),
	}

	return c, c.Validate()
}
