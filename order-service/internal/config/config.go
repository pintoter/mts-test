package config

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Grpc struct {
	Host              string
	Port              string
	MaxConnectionAge  time.Duration
	MaxConnectionIdle time.Duration
	Time              time.Duration
	Timeout           time.Duration
}

func (g *Grpc) GetGrpcAddr() string {
	return fmt.Sprintf("%s:%s", g.Host, g.Port)
}

func (g *Grpc) GetMaxConnAge() time.Duration {
	return g.MaxConnectionAge
}

func (g *Grpc) GetMaxConnIdle() time.Duration {
	return g.MaxConnectionIdle
}

func (g *Grpc) GetTime() time.Duration {
	return g.Time
}

func (g *Grpc) GetTimeout() time.Duration {
	return g.Timeout
}

type Config struct {
	Grpc Grpc
}

var config = new(Config)
var once sync.Once

func init() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("main")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("reading config err")
	}
}

func Read() *Config {
	once.Do(func() {
		var err error

		err = viper.Unmarshal(config)
		if err != nil {
			log.Fatal("reading config")
		}
	})
	return config
}
