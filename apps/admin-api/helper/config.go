package helper

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func ReadConfig(cfgFile string) error {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(cfgFile)
	return viper.ReadInConfig()
}

func LoadConfig(config interface{}) error {
	if err := bindEnvs(config); err != nil {
		return fmt.Errorf("bind env config: %s", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("unmarshal config: %s", err)
	}

	if err := defaults.Set(config); err != nil {
		return fmt.Errorf("defaults config: %s", err)
	}

	return validator.New().Struct(config)
}

// bindEnvs registers mapstructure keys with Viper so AutomaticEnv / Unmarshal
// pick up environment variables even when no config file was loaded.
func bindEnvs(config interface{}) error {
	v := reflect.ValueOf(config)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		key := field.Tag.Get("mapstructure")
		if key == "" || key == "-" {
			continue
		}
		if err := viper.BindEnv(key); err != nil {
			return err
		}
	}
	return nil
}
