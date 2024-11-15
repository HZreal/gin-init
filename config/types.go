package config

import "fmt"

/**
 * @Author nico
 * @Date 2024-11-15
 * @File: types.go
 * @Description:
 */

type Config struct {
	Gin        GinConfig         `yaml:"gin"`
	Mysql      *MysqlConfig      `yaml:"mysql"`
	Postgresql *PostgresqlConfig `yaml:"postgresql"`
	Redis      *RedisConfig      `yaml:"redis"`
	RabbitMQ   *RabbitMQConfig   `yaml:"rabbitmq"`
	GRPC       *GRPCConfig       `yaml:"grpc"`
}

// GinConfig 配置
type GinConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

func (g *GinConfig) GetAddr() string {
	return g.Host + ":" + g.Port
}

// MysqlConfig 配置
type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Timeout  string `yaml:"timeout"`
	MaxConn  int    `yaml:"maxConn"`
	MaxOpen  int    `yaml:"maxOpen"`
}

func (m *MysqlConfig) GetDsn() string {
	//
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local&timeout=" + m.Timeout
}

// RedisConfig 配置
type RedisConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"DB"`
}

func (r *RedisConfig) GetAddr() string {
	return r.Host + ":" + r.Port
}

// RabbitMQConfig 配置
type RabbitMQConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

func (q *RabbitMQConfig) GetUrl() string {
	// amqp://admin:root123456@localhost:5672/%2Flocal

	return "amqp://" + q.Username + ":" + q.Password + "@" + q.Host + ":" + q.Port + "/%2F" + q.Vhost
}

// GRPCConfig 配置
type GRPCConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (g *GRPCConfig) GetAddr() string {
	return g.Host + ":" + g.Port
}

// PostgresqlConfig 配置
type PostgresqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (m *PostgresqlConfig) GetDsn() string {
	// host=127.0.0.1 user=postgres password=szkj1234567890 dbname=eeds port=55433 sslmode=disable TimeZone=Asia/Shanghai

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", m.Host, m.Port, m.Username, m.Dbname)
}
