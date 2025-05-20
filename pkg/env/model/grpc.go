package model

/**
 * Grpc
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

type Grpc struct {
	Server   string `mapstructure:"server" env:"server"`
	Client   string `mapstructure:"client" env:"client"`
	Protocol string `mapstructure:"protocol" env:"protocol"`
	Cert     string `mapstructure:"cert" env:"cert"`
	Key      string `mapstructure:"key" env:"key"`
}
