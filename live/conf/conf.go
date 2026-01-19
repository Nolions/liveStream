package conf

import (
	"strings"

	"github.com/spf13/viper"
)

type Conf struct {
	App App `mapstructure:"app"`
}

type App struct {
	Name         string `mapstructure:"name"`
	Addr         string `mapstructure:"addr"`
	Debug        bool   `mapstructure:"debug"`
	ReadTimeout  string `mapstructure:"read_timeout"`
	WriteTimeout string `mapstructure:"write_timeout"`
	Live         Live   `mapstructure:"live"`
}

type Live struct {
	RtmpUrl string `mapstructure:"rtmp_url"`
	AppName string `mapstructure:"app_name"`
	Secret  string `mapstructure:"secret"`
}

func init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func New(file string) (conf *Conf, err error) {
	conf = &Conf{}
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
