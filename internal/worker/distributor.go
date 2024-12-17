package worker

import "github.com/hibiken/asynq"

/**
*
* distributor
* <p>
* distributor file
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
* @since 14/12/2024
*
 */

type Distributor struct {
	client *asynq.Client
}

type IDistributor interface {
}

func NewDistributor(option *asynq.RedisClientOpt) IDistributor {
	client := asynq.NewClient(option)
	return Distributor{
		client: client}
}
