package conf

import (
	"github.com/spf13/viper"
	"log"
)

var cfg config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.SetDefault("echo.listen", "0.0.0.0:8080")
	viper.SetDefault("mysql.port", 3306)
	viper.SetDefault("mysql.max_idle_conn", 10)
	viper.SetDefault("mysql.max_open_conn", 5)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
}
func GetCfg() *config {
	return &cfg
}

type config struct {
	Echo struct {
		Listen string `toml:"listen"`
	} `toml:"echo"`
	Jwt struct {
		Exp    int    `toml:"exp"`
		Secret string `toml:"secret"`
	} `toml:"jwt"`
	Admin struct {
		User     string `toml:"user"`
		Password string `toml:"password"`
	} `toml:"admin"`
	Mysql struct {
		Host        string `toml:"host"`
		Port        int    `toml:"port"`
		User        string `toml:"user"`
		Password    string `toml:"password"`
		Database    string `toml:"database"`
		MaxIdleConn int    `toml:"max_idle_conn"`
		MaxOpenConn int    `toml:"max_open_conn"`
	} `toml:"mysql"`
}
