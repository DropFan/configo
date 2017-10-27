package main

import (
	"fmt"
	"time"

	"github.com/DropFan/configo"
)

// DemoConfig demo config struct
type DemoConfig struct {
	DemoName string     `toml:"name" json:"name"`
	RunMode  string     `toml:"mode" json:"mode" yaml:"mode"`
	Server   ServerConf `toml:"server" json:"server"`
	Redis    RedisConf
	Mysql    MysqlConf
}

// ServerConf demo server config
type ServerConf struct {
	Type string `toml:"type" json:"type"` // "thrift" or "http"
	Host string `toml:"host" json:"host"`
	Port int32  `toml:"port" json:"port"`
	Addr string `toml:"addr" json:"addr"` // "host:port"
}

// RedisConf demo redis config
type RedisConf struct {
	Host string `toml:"host" json:"host"`
	Port int    `toml:"port" json:"port"`
}

// MysqlConf demo mysql config
type MysqlConf struct {
	Host   string
	Port   int
	User   string
	Pass   string
	DBName string //`toml:"dbname" json:"dbname"`
}

var (
	demoConf     *DemoConfig
	configLoader *configo.Loader
	err          error
)

func main() {
	demoConf = new(DemoConfig)
	f := "demo.conf"

	configLoader, err = configo.NewFromFile(f, demoConf)
	if err != nil {
		panic(err)
	}
	fmt.Println("load config done.")
	fmt.Printf("demoConf: %#v\n", demoConf)

	for i := 10; i > 0; i-- {
		fmt.Printf("sleep %d sec...You can modify config file to test reload method\n", i)
		time.Sleep(time.Second)
	}

	fmt.Println("reload config from file...")

	err = configLoader.Reload() // or configLoader.LoadFromFile(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("demoConf: %#v\n", demoConf)
}

func demo2() {
	demoConf = new(DemoConfig)
	f := "demo.conf"
	configLoader, err = configo.New(demoConf)
	if err != nil {
		panic(err)
	}

	err = configLoader.LoadFromFile(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("load config done.")
	fmt.Printf("demoConf: %#v\n", demoConf)
}
