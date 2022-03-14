package main

import (
	"fmt"
	"sync"
)

/********************************************************************
created:    2022-03-13
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type Service struct {
	taskChan  chan Task
	closeChan chan struct{}
	closeOnce sync.Once
}

func NewService() *Service {
	var service = &Service{
		// service的成员全是线程安全的
		taskChan:  make(chan Task, 8),
		closeChan: make(chan struct{}),
	}

	// one loop per thread
	go service.goLoop()
	return service
}

func (my *Service) goLoop() {
	// 栈变量无线程安全问题, 因此fetus无需担心data race
	var fetus = &ServiceFetus{}

	for {
		select {
		case task := <-my.taskChan:
			fetus.Process(task)
		case <-my.closeChan:
			return
		}
	}
}

func (my *Service) SendTask(task Task) {
	// Service对外开放的接口都设计成线程安全的
	if task != nil {
		select {
		case my.taskChan <- task:
			fmt.Println("receive task:", task)
		case <-my.closeChan:
			fmt.Println("ignored task:", task)
		}
	}
}

// Close 支持幂等
func (my *Service) Close() error {
	my.closeOnce.Do(func() {
		// 只关闭closeChan. 不需要也不允许关闭taskChan
		close(my.closeChan)
		fmt.Println("service is closed")
	})

	return nil
}
