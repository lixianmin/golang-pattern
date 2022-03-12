package main

/********************************************************************
created:    2022-03-12
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

func main() {
	var data = 1
	go func() {
		for {
			data = 2
		}
	}()

	for {
		data = 3
	}
}
