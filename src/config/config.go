package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var config struct {
	HTTP HTTPConfig
}

// Init 初始化設定
func Init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.Unmarshal(&config); err != nil {
		panic("viper.Unmarshal failed: " + err.Error())
	}
}

// Display 印出目前設定檔的資訊
func Display() {
	fmt.Printf("%#v\n", config)
}

// HTTP 回傳複製的 Http 設定檔
func HTTP() HTTPConfig { return config.HTTP }
