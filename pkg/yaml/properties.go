package yaml

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

/**
*
* properties
* <p>
* properties file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is shared under a collaborative license.
* Contributions, suggestions, and improvements are welcome!
* Feel free to fork, modify, and submit pull requests under the terms of the repository's license.
* Please ensure proper attribution to the original author(s) and maintain this notice in derivative works.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 8/12/2024
*
 */

var YAML Properties

type Properties struct {
	Server    Server    `mapstructure:"server" yaml:"server"`
	Ssh       Ssh       `mapstructure:"ssh" yaml:"ssh"`
	Scheduler Scheduler `mapstructure:"scheduler" yaml:"scheduler"`
	Database  Database  `mapstructure:"database" yaml:"database"`
	Rabbitmq  Rabbitmq  `mapstructure:"rabbitmq" yaml:"rabbitmq"`
	Grpc      Grpc      `mapstructure:"grpc" yaml:"grpc"`
}

// Server use mapstructure in document github.com/go-viper/mapstructure/v2
type Server struct {
	Host        string `mapstructure:"host" yaml:"host"`
	Port        int    `mapstructure:"port" yaml:"port"`
	Name        string `mapstructure:"name" yaml:"name"`
	Timeout     int    `mapstructure:"timeout" yaml:"timeout"`
	Logging     string `mapstructure:"logging" yaml:"logging"`
	Environment string `mapstructure:"environment" yaml:"environment"`
	Logs        string `mapstructure:"logs" yaml:"logs"`
}

type Ssh struct {
	Host       string `mapstructure:"host" yaml:"host"`
	Port       string `mapstructure:"port" yaml:"port"`
	Protocol   string `mapstructure:"protocol" yaml:"protocol"`
	Username   string `mapstructure:"username" yaml:"username"`
	Password   string `mapstructure:"password" yaml:"password"`
	PrivateKey string `mapstructure:"private_key" yaml:"private_key"`
	PublicKey  string `mapstructure:"public_key" yaml:"public_key"`
	KnownHosts string `mapstructure:"known_hosts" yaml:"known_hosts"`
	SftpPath   string `mapstructure:"sftp_path" yaml:"sftp_path"`
	Enable     bool   `mapstructure:"enable" yaml:"enable"`
}

type Scheduler struct {
	Enable bool `mapstructure:"enable" yaml:"enable"`
}

type Database struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	Dbname   string `mapstructure:"db_name" yaml:"db_name"`
	Driver   string `mapstructure:"driver" yaml:"driver"`
	IsTest   bool   `mapstructure:"is_test" yaml:"is_test"`
}

type Rabbitmq struct {
	Exchange   Exchange   `mapstructure:"exchange" yaml:"exchange"`
	Host       string     `mapstructure:"host" yaml:"host"`
	Password   string     `mapstructure:"password" yaml:"password"`
	Port       int        `mapstructure:"port" yaml:"port"`
	Protocol   string     `mapstructure:"protocol" yaml:"protocol"`
	Queue      Queue      `mapstructure:"queue" yaml:"queue"`
	RoutingKey RoutingKey `mapstructure:"routing_key" yaml:"routing_key"`
	TlsEnabled bool       `mapstructure:"tls_enabled" yaml:"tls_enabled"`
	User       string     `mapstructure:"user" yaml:"user"`
	Vhost      string     `mapstructure:"vhost" yaml:"vhost"`
}
type Exchange struct {
	Durable bool   `mapstructure:"durable" yaml:"durable"`
	Name    string `mapstructure:"name" yaml:"name"`
	Type    string `mapstructure:"type" yaml:"type"`
}
type Queue struct {
	Name       string `mapstructure:"name" yaml:"name"`
	Durable    bool   `mapstructure:"durable" yaml:"durable"`
	MessageTtl int    `mapstructure:"message_ttl" yaml:"message_ttl"`
	Type       string `mapstructure:"type" yaml:"type"`
}
type RoutingKey struct {
	Request  string `mapstructure:"request" yaml:"request"`
	Response string `mapstructure:"response" yaml:"response"`
}

type Grpc struct {
	Server   string `mapstructure:"server" yaml:"server"`
	Client   string `mapstructure:"client" yaml:"client"`
	Protocol string `mapstructure:"protocol" yaml:"protocol"`
	Cert     string `mapstructure:"cert" yaml:"cert"`
	Key      string `mapstructure:"key" yaml:"key"`
}

type ParameterBroker struct {
	Port       int
	Uri        string
	Exchange   string
	Queue      string
	Vhost      string
	RoutingKey string
}

type IParameterBroker interface {
	GetUri() string
	GetVhost() string
	GetQueueName() string
	GetRoutingKey() string
	GetExchange() string
}

func GetUriDatasource() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		YAML.Database.Driver,
		YAML.Database.User,
		YAML.Database.Password,
		YAML.Database.Host,
		YAML.Database.Port,
		YAML.Database.Dbname)
}

func (r *Rabbitmq) GetUri() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/",
		YAML.Rabbitmq.Protocol,
		YAML.Rabbitmq.User,
		YAML.Rabbitmq.Password,
		YAML.Rabbitmq.Host,
		YAML.Rabbitmq.Port)
}
func (r *Rabbitmq) GetVhost() string {
	return r.Vhost
}
func (r *Rabbitmq) GetQueueName() string {
	return r.Queue.Name
}
func (r *Rabbitmq) GetExchange() string {
	return r.Exchange.Name
}

func (r *Rabbitmq) GetRoutingKey() string {
	return r.RoutingKey.Request
}

func LoadProperties() {

	path := os.Getenv("CONFIG_PATH")

	if path == "" {
		fmt.Println("err")
		log.Fatal("La variable de entorno CONFIG_PATH no está definida")
	}

	viper.SetConfigName("application")
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	// Si se desea pasar el archivo por variable de entorno:
	//viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var file viper.ConfigFileNotFoundError
		if errors.As(err, &file) {
			fmt.Println("Config file not found")
			log.Fatalf("Error reading config file, %s", file)
			return
		}
		return
	}

	err := viper.WriteConfig()

	if err != nil {
		fmt.Println("error writing config file")
		log.Fatalf("Error writing config file, %s", err)
		return
	}

	err = viper.Unmarshal(&YAML)

	if err != nil {
		log.Fatalf("Error unmarshalling config, %s", err)
		return
	}

	log.Println("Successfully loaded config")
}
