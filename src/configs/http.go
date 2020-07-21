package configs

import "github.com/spf13/viper"

// HTTPConfig 與 HTTP 相關的設定
type HTTPConfig struct {
	Port   string
	Enable bool
}

func init() {
	viper.SetDefault("http.port", "12345")
	viper.SetDefault("http.enable", true)
}
