package mq

import "fmt"

/**
 * Rabbitmq
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

type Rabbitmq struct {
	Protocol   string     `mapstructure:"protocol" env:"protocol"`
	Host       string     `mapstructure:"host" env:"host"`
	Port       int        `mapstructure:"port" env:"port"`
	User       string     `mapstructure:"user" env:"user"`
	Password   string     `mapstructure:"password" env:"password"`
	Vhost      string     `mapstructure:"vhost" env:"vhost"`
	Exchange   Exchange   `mapstructure:"exchange" env:"exchange"`
	Queue      Queue      `mapstructure:"queue" env:"queue"`
	RoutingKey RoutingKey `mapstructure:"routing_key" env:"routing_key"`
	TlsEnabled bool       `mapstructure:"tls_enabled" env:"tls_enabled"`
	Binding    Binding    `mapstructure:"binding" env:"binding"`
}

func (r *Rabbitmq) GetUri() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d",
		r.Protocol,
		r.User,
		r.Password,
		r.Host,
		r.Port,
	)
}
