package model

/**
 * Server
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

// Server use mapstructure in document github.com/go-viper/mapstructure/v2
type Server struct {
	Host        string `mapstructure:"host" env:"host"`
	Port        int    `mapstructure:"port" env:"port"`
	Name        string `mapstructure:"name" env:"name"`
	Timeout     int    `mapstructure:"timeout" env:"timeout"`
	Logging     string `mapstructure:"logging" env:"logging"`
	Environment string `mapstructure:"environment" env:"environment"`
	Logs        string `mapstructure:"logs" env:"logs"`
}
