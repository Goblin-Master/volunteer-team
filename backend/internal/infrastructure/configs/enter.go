package configs

import (
	"fmt"
)

var Conf = new(Config)

type Config struct {
	App   App   `mapstructure:"app"`
	DB    DB    `mapstructure:"db"`
	Redis Redis `mapstructure:"redis"`
	Auth  Auth  `mapstructure:"auth"`
	Email Email `mapstructure:"email"`
}

type App struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	Env     string `mapstructure:"env"`
	Log     string `mapstructure:"log"`
	Uploads string `mapstructure:"uploads"`
}

func (app *App) Link() string {
	return fmt.Sprintf("%s:%d", app.Host, app.Port)
}

type DB struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

func (db *DB) DSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", db.Username, db.Password, db.Host, db.Port, db.DBName, "5s")
	return dsn
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	Enable   bool   `mapstructure:"enable"`
}

func (redis *Redis) DSN() string {
	dsn := fmt.Sprintf("%s:%d", redis.Host, redis.Port)
	return dsn
}

// JWT 认证需要的密钥和过期时间配置
type Auth struct {
	AccessSecret string `json:"accessSecret"`
	AccessExpire int64  `json:"accessExpire"`
}

type Email struct {
	Domain       string `mapstructure:"domain"`
	Port         int    `mapstructure:"port"`
	SendEmail    string `mapstructure:"sendEmail"`
	AuthCode     string `mapstructure:"authCode"`
	SendNickname string `mapstructure:"sendNickname"`
	Subject      string `mapstructure:"subject"`
	SSL          bool   `mapstructure:"ssl"`
	TLS          bool   `mapstructure:"tls"`
}
