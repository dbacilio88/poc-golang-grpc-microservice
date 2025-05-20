package task

import (
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/service"
	"github.com/dbacilio88/poc-golang-grpc-microservice/pkg/shared/store"
	"github.com/dbacilio88/poc-golang-grpc-microservice/pkg/utils"
	"github.com/madflojo/tasks"
	"go.uber.org/zap"
	"path/filepath"
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
	*zap.Logger
	utils.IFiles
}

type ITask interface {
	Create() *tasks.Scheduler
	Run(task *tasks.Scheduler, ws string, msg service.IMessaging)
}

func NewTask(log *zap.Logger) ITask {
	//eb := event.NewBrokerConfig(log).NewBrokerServer(broker.RabbitMqInstance)
	memory := store.NewMemory()
	files := utils.NewFiles(log, memory)
	return &Task{
		Logger: log,
		IFiles: files,
		//	broker: eb,
	}
}

func (t *Task) Create() *tasks.Scheduler {
	t.Info("Create new task")
	return tasks.New()
}
func (t *Task) Run(task *tasks.Scheduler, ws string, msg service.IMessaging) {
	t.Info("Run new task")
	tsk := &tasks.Task{
		Interval:          time.Minute * 1,
		RunOnce:           false,
		RunSingleInstance: false,
		TaskFunc: func() error {
			var err error
			if err = t.IFiles.GenerateFile(); err != nil {
				return err
			}
			//err := t.broker.BrokerPublisher([]byte("hola mundo"))
			//if err != nil {
			//	t.log.Error("Error publishing message", zap.Error(err))
			//	return err
			//	}
			abs, err := filepath.Abs(ws)
			err = t.IFiles.ScanDir(abs, msg)
			return err
		},
		ErrFunc: func(err error) {
			t.Info("Error task ", zap.Error(err))
			//t.Error("Error task", zap.Error(err))
		},
	}

	add, err := task.Add(tsk)

	if err != nil {
		t.Error("Error adding task", zap.Error(err))
		return
	} else {
		t.Info("Added task", zap.Any("add", add))
	}
}
