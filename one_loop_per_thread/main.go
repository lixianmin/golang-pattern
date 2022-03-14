package main

import "time"

/********************************************************************
created:    2022-03-13
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

func main() {
	var service = NewService()
	for i := 0; i < 3; i++ {
		service.SendTask(NewTask())
	}

	time.Sleep(time.Second * 2)

	// 可以Close()多次
	_ = service.Close()
	_ = service.Close()

	for i := 0; i < 3; i++ {
		service.SendTask(NewTask())
	}

	time.Sleep(time.Hour)
}
