package env

import (
	"errors"
	"fmt"
	model2 "github.com/dbacilio88/poc-golang-grpc-microservice/pkg/env/model"
	"github.com/dbacilio88/poc-golang-grpc-microservice/pkg/env/mq"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
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
	Server    model2.Server    `mapstructure:"server" yaml:"server"`
	Ssh       model2.Ssh       `mapstructure:"ssh" yaml:"ssh"`
	Scheduler model2.Scheduler `mapstructure:"scheduler" yaml:"scheduler"`
	Database  model2.Database  `mapstructure:"database" yaml:"database"`
	Rabbitmq  mq.Rabbitmq      `mapstructure:"rabbitmq" yaml:"rabbitmq"`
	Grpc      model2.Grpc      `mapstructure:"grpc" yaml:"grpc"`
	Workspace model2.Workspace `mapstructure:"workspace" yaml:"workspace"`
}

func LoadProperties() string {

	err := godotenv.Load(".env")

	if err != nil {
		return fmt.Sprintf("Error loading .env file %s", err)
	}

	path := os.Getenv("CONFIG_PATH")

	if path == "" {
		return fmt.Sprint("Error loading .env path is empty")
	}

	viper.SetConfigName("application")
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	//viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var file viper.ConfigFileNotFoundError
		if errors.As(err, &file) {
			return fmt.Sprintf("Error reading config file, %s", file)
		}
		return fmt.Sprintf("Error reading viper config file, %s", file)
	}

	err = viper.WriteConfig()

	if err != nil {
		return fmt.Sprintf("Error writing config file, %s", err)
	}

	if err = viper.Unmarshal(&YAML); err != nil {
		return fmt.Sprintf("Error unmarshalling config, %s", err)
	}

	return "Successfully loaded config"
}
