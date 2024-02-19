package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	Conf *Config
)

func Init(confPath string) error {
	err := initConfig(confPath)
	if err != nil {
		return err
	}
	return nil
}

func initConfig(confPath string) error {
	if confPath != "" {
		viper.SetConfigFile(confPath)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("lapp")
	if err := viper.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}
	err := viper.Unmarshal(&Conf)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

type Config struct {
	MySQLConfig MySQLConfig
}

type MySQLConfig struct {
	Name            string
	Addr            string
	UserName        string
	Password        string
	ShowLog         bool
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime int
}
