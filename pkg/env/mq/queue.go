package mq

/**
 * queue
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
 * @since 5/10/2025
 */

type Queue struct {
	Files QueueType `mapstructure:"files" env:"files"`
	Email QueueType `mapstructure:"email" env:"email"`
}

type QueueType struct {
	Name       string `mapstructure:"name" env:"name"`
	Durable    bool   `mapstructure:"durable" env:"durable"`
	MessageTtl int    `mapstructure:"message_ttl" env:"message_ttl"`
	Types      string `mapstructure:"types" env:"types"`
}
