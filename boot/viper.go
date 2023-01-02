package boot

import (
	"fmt"
	"github.com/spf13/viper"
	g "zhihu/global"
)

const (
	configPath = "conf/config.yaml"
)

func ViperSetup() {
	v := viper.New()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed err:%v", err))
	}
	if err := v.Unmarshal(&g.Config); err != nil {
		panic(fmt.Errorf("unmarshal config failed err:%v", err))
	}
}
