package environment

import (
	"github.com/caarlos0/env/v6"
)

type appConfig struct {
	DSN string `env:"DSN" envDefault:"root:@/cloud?parseTime=true"`
}

var Config appConfig

func Initialize() {
	if err := env.Parse(&Config); err != nil {
		panic(err.Error())
	}
}
