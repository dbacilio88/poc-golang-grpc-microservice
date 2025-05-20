package model

/**
 * Ssh
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

type Ssh struct {
	Host       string `mapstructure:"host" env:"host"`
	Port       string `mapstructure:"port" env:"port"`
	Protocol   string `mapstructure:"protocol" env:"protocol"`
	Username   string `mapstructure:"username" env:"username"`
	Password   string `mapstructure:"password" env:"password"`
	PrivateKey string `mapstructure:"private_key" env:"private_key"`
	PublicKey  string `mapstructure:"public_key" env:"public_key"`
	KnownHosts string `mapstructure:"known_hosts" env:"known_hosts"`
	SftpPath   string `mapstructure:"sftp_path" env:"sftp_path"`
	Enable     bool   `mapstructure:"enable" env:"enable"`
}
