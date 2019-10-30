package configs

import (
	"errors"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/constants"
	"github.com/spf13/viper"
	"strings"
)

var cfg *config

func init() {
	m := strings.Split(args.Config, ",")
	switch m[0] {
	case "consul":
		break
	default:
		cfg = loadFromToml(m[0])
	}
}
func Env(key string, value ...interface{}) interface{} {
	mode := args.Mode
	modeKey := strings.Join([]string{mode, key}, ".")
	commKey := strings.Join([]string{"common", key}, ".")
	if cfg.IsSet(modeKey) {
		return cfg.Get(modeKey)
	} else if cfg.IsSet(commKey) {
		return cfg.Get(commKey)
	} else {
		return value[0]
	}
}

func EnvString(key string, value ...interface{}) string {
	mode := args.Mode
	modeKey := strings.Join([]string{mode, key}, ".")
	commKey := strings.Join([]string{"common", key}, ".")
	var ret string
	if cfg.IsSet(modeKey) {
		ret = cfg.GetString(modeKey)
	} else if cfg.IsSet(commKey) {
		ret = cfg.GetString(commKey)
	} else {
		ret = value[0].(string)
	}
	return ret
}
func EnvInt(key string, value ...interface{}) int {
	mode := args.Mode
	modeKey := strings.Join([]string{mode, key}, ".")
	commKey := strings.Join([]string{"common", key}, ".")
	var ret int
	if cfg.IsSet(modeKey) {
		ret = cfg.GetInt(modeKey)
	} else if cfg.IsSet(commKey) {
		ret = cfg.GetInt(commKey)
	} else {
		ret = value[0].(int)
	}
	return ret
}
func EnvBool(key string, value ...interface{}) bool {
	mode := args.Mode
	modeKey := strings.Join([]string{mode, key}, ".")
	commKey := strings.Join([]string{"common", key}, ".")
	var ret bool
	if cfg.IsSet(modeKey) {
		ret = cfg.GetBool(modeKey)
	} else if cfg.IsSet(commKey) {
		ret = cfg.GetBool(commKey)
	} else {
		ret = value[0].(bool)
	}
	return ret
}
func EnvStringSlice(key string, value ...interface{}) []string {
	mode := args.Mode
	modeKey := strings.Join([]string{mode, key}, ".")
	commKey := strings.Join([]string{"common", key}, ".")
	var ret []string
	if cfg.IsSet(modeKey) {
		ret = cfg.GetStringSlice(modeKey)
	} else if cfg.IsSet(commKey) {
		ret = cfg.GetStringSlice(commKey)
	} else {
		ret = value[0].([]string)
	}
	return ret
}

type config struct {
	*viper.Viper
}

func loadFromToml(fileName string) *config {
	c := &config{}
	c.Viper = viper.New()
	c.SetConfigName(fileName)
	c.AddConfigPath("./")
	c.AddConfigPath("../")
	c.AddConfigPath("../../")
	c.AddConfigPath("./env/")
	c.AddConfigPath("../env/")
	c.SetConfigType("toml")
	if err := c.ReadInConfig(); err != nil {
		panic(errors.New(constants.ErrLoadEnv))
	}
	return c
}
