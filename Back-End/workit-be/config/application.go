package config

import (
	"github.com/DeniesKresna/gohelper/utint"
	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/ricnah/workit-be/types/constants"
)

type App struct {
	Name     string
	Host     string
	Port     string
	Env      string
	Secret   string
	DataRows int64
}

func (cfg *Config) SetConfigApplication() (err error) {
	app := App{
		Name:     utstring.GetEnv(constants.ENV_APP_NAME),
		Host:     utstring.GetEnv(constants.ENV_APP_HOST),
		Port:     utstring.GetEnv(constants.ENV_APP_PORT),
		Env:      utstring.GetEnv(constants.ENV_APP_ENV),
		Secret:   utstring.GetEnv(constants.ENV_APP_SECRET),
		DataRows: utint.Convert64FromString(utstring.GetEnv(constants.ENV_APP_MAX_ROWS), 20),
	}

	cfg.App = app
	return
}
