package model

import "fmt"

/**
 * Database
 * <p>
 * This file contains core data structures and logic used throughout the application.
 *
 * <p><strong>Copyright © 2025 – All rights reserved.</strong></p>
 *
 * <p>This source code is distributed under a collaborative license.</p>
 *
 * <p>
 * Contributions, suggestions, and improvements are welcome!
 * You are free to fork, modify, and submit pull requests under the terms of the repository's license.
 * Please ensure proper attribution to the original author(s) and preserve this notice in derivative works.
 * </p>
 *
 * @author Christian Bacilio De La Cruz
 * @email dbacilio88@outlook.es
 * @since 5/8/2025
 */

type Database struct {
	Host     string `mapstructure:"host" env:"host"`
	Port     int    `mapstructure:"port" env:"port"`
	User     string `mapstructure:"user" env:"user"`
	Password string `mapstructure:"password" env:"password"`
	Dbname   string `mapstructure:"db_name" env:"db_name"`
	Driver   string `mapstructure:"driver" env:"driver"`
	IsTest   bool   `mapstructure:"is_test" env:"is_test"`
}

func (p *Database) GetUrl() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		p.Driver,
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.Dbname)
}
