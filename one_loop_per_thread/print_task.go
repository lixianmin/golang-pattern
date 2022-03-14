package main

import (
	"fmt"
)

/********************************************************************
created:    2022-03-13
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type PrintTask struct {
	id int
}

func (my *PrintTask) Run() {
	fmt.Println("print task: ", my)
}
