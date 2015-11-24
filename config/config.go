package config

import (
	"os"
	"path/filepath"

	"github.com/miyabayt/gin-rest/logger"
	conf "github.com/robfig/config"
)

var (
	mergedConfig *MergedConfig
)

type MergedConfig struct {
	c   *conf.Config
	env string
}

func (mc *MergedConfig) Int(key string) (int, error) {
	value, err := mc.c.Int(mc.env, key)
	if err != nil {
		logger.Errorf("failed to get int value. key=%s", key)
		return -9999, err
	}

	return value, nil
}

func (mc *MergedConfig) String(key string) (string, error) {
	value, err := mc.c.String(mc.env, key)
	if err != nil {
		logger.Errorf("failed to get string value. key=%s", key)
		return "", err
	}

	return value, nil
}

func Int(key string) (int, error) {
	return mergedConfig.Int(key)
}

func String(key string) (string, error) {
	return mergedConfig.String(key)
}

func init() {
	confPaths, _ := filepath.Glob("./conf/*.conf")

	for _, configPath := range confPaths {
		conf, err := conf.ReadDefault(configPath)
		if err != nil {
			panic("failed to read config file.")
		}

		if mergedConfig == nil {
			env := os.Getenv("ENV")
			if len(env) == 0 {
				env = "dev"
			}
			mergedConfig = &MergedConfig{conf, env}
		} else {
			mergedConfig.c.Merge(conf)
		}
	}
}
