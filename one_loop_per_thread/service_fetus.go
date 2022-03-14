package main

/********************************************************************
created:    2022-03-13
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type ServiceFetus struct {
}

func (my *ServiceFetus) Process(task Task) {
	task.Run()
}
