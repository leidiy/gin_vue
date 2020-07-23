package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type App struct {
	Title   string `toml:"title"`
	Explain string `toml:"explain"`
	Mod     string `toml:"mod"`
	Addr    string `toml:"addr"`
	Srv     string `toml:"srv"`
	JwtKey  string `toml:"jwtKey"`
	JwtExp  string `toml:"jwtExp"`
	Author  struct {
		Name    string `toml:"name"`
		Website string `toml:"website"`
	}
	WeChat struct {
		AppId     string `toml:"appId"`
		AppSecret string `toml:"appSecret"`
	}
	Database struct {
		User     string `toml:"user"`
		Password string `toml:"password"`
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		DbName   string `toml:"dbName"`
		Params   string `toml:"params"`
	}
	Gorm struct {
		Idle        int16 `toml:"idle"`
		Open        int16 `toml:"open"`
		Show        bool  `toml:"show"`
		Sync        bool  `toml:"sync"`
		CacheEnable bool  `toml:"cache_enable"`
		CacheCount  int16 `toml:"cache_count"`
	}
}

func LoadConf() {
	var app App
	file, err := toml.DecodeFile("conf.toml", &app)
	if err != nil {
		log.Fatalf("解析配置文件出错")
		return
	}
	fmt.Println(file)

}
