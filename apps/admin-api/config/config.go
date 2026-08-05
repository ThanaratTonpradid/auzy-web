package config

import (
	"github.com/dollarsignteam/go-logger"
	"github.com/spf13/viper"

	"auzy-api/helper"
)

var log *logger.Logger

func init() {
	log = logger.NewLogger(logger.LoggerOptions{
		Name:       helper.GetPackageName(),
		HideCaller: true,
	})
}

func AutoReadConfig(cfgFile string) {
	configFile := "local.env"
	explicit := cfgFile != ""
	if explicit {
		configFile = cfgFile
	}

	if err := helper.ReadConfig(configFile); err != nil {
		if explicit {
			log.Fatalf("Config file: %s, %s", configFile, err)
		}
		// Default local.env is optional; Docker / production can rely on env vars.
		log.Infof("No config file loaded (%s); using environment variables", err)
		return
	}
	log.Infof("Using config file: %s", viper.ConfigFileUsed())
}

func AutoLoadConfig(config interface{}) {
	if err := helper.LoadConfig(config); err != nil {
		log.Fatal(err.Error())
	}
	log.Infof("Config loaded: %+v", config)
}
