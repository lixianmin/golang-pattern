package main

/********************************************************************
created:    2022-03-13
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type Task interface {
	Run()
}

func NewTask() Task {
	idGenerator++
	var task = &PrintTask{id: idGenerator}
	return task
}
