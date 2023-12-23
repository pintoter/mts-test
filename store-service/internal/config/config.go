package config

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

const (
	configPath = "./configs"
	configName = "main"
	envFile    = "./.env"
)

type DB struct {
	User            string
	Password        string
	Host            string
	Port            string
	Name            string
	Sslmode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
}

func (db *DB) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", db.User, db.Password, db.Host, db.Port, db.Name, db.Sslmode)
}

func (db *DB) GetMaxOpenConns() int {
	return db.MaxOpenConns
}

func (db *DB) GetMaxIdleConns() int {
	return db.MaxIdleConns
}

func (db *DB) GetConnMaxIdleTime() time.Duration {
	return db.ConnMaxIdleTime
}

func (db *DB) GetConnMaxLifetime() time.Duration {
	return db.ConnMaxLifetime
}

type Kafka struct {
	Brokers []string
	Topics  []string
}

func (k *Kafka) GetBrokers() []string {
	return k.Brokers
}

func (k *Kafka) GetTopics() []string {
	return k.Topics
}

type Config struct {
	DB    DB
	Kafka Kafka
}

var config = new(Config)
var once sync.Once

func init() {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("loading env file")
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	err = viper.ReadInConfig()
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

		err = envconfig.Process("db", &config.DB)
		if err != nil {
			log.Fatal("error: get env for db")
		}
	})
	return config
}
