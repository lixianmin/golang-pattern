package close_chan_pattern

import (
	"fmt"
	"github.com/lixianmin/got/loom"
)

/********************************************************************
created:    2022-03-12
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type PlayerWaitClose struct {
	commandChan chan int
	wc          loom.WaitClose
}

func NewPlayerWaitClose() *PlayerWaitClose {
	var player = &PlayerWaitClose{
		commandChan: make(chan int, 8),
	}

	return player
}

func (my *PlayerWaitClose) goLoop() {
	var closeChan = my.wc.C()

	for {
		select {
		case cmd := <-my.commandChan:
			my.runCommand(cmd)
		case <-closeChan:
			return
		}
	}
}

func (my *PlayerWaitClose) SendCommand(cmd int) {
	select {
	case my.commandChan <- cmd:
	case <-my.wc.C():
	}
}

func (my *PlayerWaitClose) Close() error {
	return my.wc.Close(func() error {
		// here can be reached only once
		return nil
	})
}

func (my *PlayerWaitClose) runCommand(cmd int) {
	fmt.Println(cmd)
}
