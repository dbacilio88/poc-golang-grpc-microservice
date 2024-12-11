package task

import (
	"fmt"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/event"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/event/broker"
	"github.com/madflojo/tasks"
	"go.uber.org/zap"
	"time"
)

/**
*
* tasks
* <p>
* tasks file
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

type Task struct {
	broker *event.BrokerConfig
	log    *zap.Logger
}

type ITask interface {
	Create() *tasks.Scheduler
	Run(task *tasks.Scheduler)
}

func NewTask(log *zap.Logger) *Task {

	eb := event.NewBrokerConfig(log).
		NewBrokerServer(broker.RabbitMqInstance)

	return &Task{
		log:    log,
		broker: eb,
	}
}

func (t *Task) Create() *tasks.Scheduler {
	t.log.Info("Create new task")
	return tasks.New()
}
func (t *Task) Run(task *tasks.Scheduler) {
	t.log.Info("Run new task")
	tsk := &tasks.Task{
		Interval:          time.Minute * 3,
		RunOnce:           false,
		RunSingleInstance: false,
		TaskFunc: func() error {
			fmt.Println("Run task")
			err := t.broker.Publisher([]byte("hola mundo"))
			if err != nil {
				t.log.Error("Error publishing message", zap.Error(err))
				return err
			}

			return nil
		},
		ErrFunc: func(err error) {
			t.log.Error("Error task", zap.Error(err))
		},
	}

	add, err := task.Add(tsk)

	if err != nil {
		t.log.Error("Error adding task", zap.Error(err))
		return
	} else {
		t.log.Info("Added task", zap.Any("add", add))
	}
}
