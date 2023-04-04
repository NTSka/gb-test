package eth

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

type Config struct {
	URI string
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.URI, validation.Required))
}

func NewConfig(v *viper.Viper) (*Config, error) {
	c := &Config{
		URI: v.GetString("client.eth.URI"),
	}

	return c, c.Validate()
}
