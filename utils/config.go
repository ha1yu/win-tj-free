package utils

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var Configs *Config

func init() {
	//InitConfig()
	log.Println(123)
}

type Config struct {
	Debug         bool
	Name          string // 当前服务器名称
	WebServerPort string // 当前服务器主机监听端口号
	WebUserName   string // 当前服务器主机监听端口号
	WebUserPasswd string // 当前服务器主机监听端口号
	ServerVersion string // 当前服务端版本号
	MessageAesKey string // AES加密密钥,服务端客户端需要一致,服务端为了性能直接写死,客户端不写死,每次计算出来
}

func InitConfig() {
	viper.SetConfigName("config")                          // 配置文件前缀
	viper.SetConfigType("yml")                             // 配置文件后缀
	viper.AddConfigPath("./")                              // 绑定配置路径
	viper.AutomaticEnv()                                   // 绑定全部环境变量
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // 字符串替换

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置文件出错", err)
	}
	viper.WatchConfig() // 开启热加载，如果配置文件发生改变则热加载

	viper.OnConfigChange(func(e fsnotify.Event) { // 热加载的回调
		Configs = GetConfig()
		log.Println("==> 重新加载配置文件", e.Name, e.Op)
	})

	Configs = GetConfig()
}

func GetConfig() *Config {
	var config = &Config{
		Debug:         viper.GetBool("Debug"),
		Name:          viper.GetString("name"),
		ServerVersion: viper.GetString("ServerVersion"),
		WebUserName:   viper.GetString("WebUserName"),
		WebUserPasswd: viper.GetString("WebUserPasswd"),
		MessageAesKey: viper.GetString("MessageAesKey"),
	}
	return config
}

func (c Config) String() string {
	str, _ := json.Marshal(c)
	return string(str)
}
