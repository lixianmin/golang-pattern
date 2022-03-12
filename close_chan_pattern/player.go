package close_chan_pattern

import "fmt"

/********************************************************************
created:    2022-03-12
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type Player struct {
	commandChan chan int
	closeChan   chan struct{}
}

func NewPlayer() *Player {
	var player = &Player{
		commandChan: make(chan int, 8),
		closeChan:   make(chan struct{}),
	}

	return player
}

func (my *Player) goLoop() {
	for {
		select {
		case cmd := <-my.commandChan:
			my.runCommand(cmd)
		case <-my.closeChan:
			return
		}
	}
}

func (my *Player) SendCommand(cmd int) {
	select {
	case my.commandChan <- cmd:
	case <-my.closeChan:
	}
}

func (my *Player) Close() error {
	// todo 多次调用会panic
	close(my.closeChan)
	return nil
}

func (my *Player) runCommand(cmd int) {
	fmt.Println(cmd)
}
