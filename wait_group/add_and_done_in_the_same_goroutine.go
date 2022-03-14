package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/********************************************************************
created:    2022-03-13
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

func main() {
	wrongExample()
	rightExample()
}

func wrongExample() {
	var wg sync.WaitGroup
	var data int32

	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			atomic.AddInt32(&data, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("wrong example: data", data)
}

func rightExample() {
	var wg sync.WaitGroup
	var data int32

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("right example: data", data)
}
