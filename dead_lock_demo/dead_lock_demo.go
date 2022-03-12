package main

import (
	"fmt"
	"github.com/lixianmin/gonsole"
	"github.com/lixianmin/got/loom"
	"github.com/lixianmin/logo"
	"log"
	"net/http"
	"sync"
)

/********************************************************************
created:    2022-03-12
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

func main() {
	var logger = logo.GetLogger().(*logo.Logger)
	logger.SetFilterLevel(logo.LevelDebug)

	var webPort = 8888
	var mux = http.NewServeMux()
	gonsole.NewServer(mux,
		gonsole.WithPort(webPort),
		gonsole.WithPageBody("<H1>This is a very huge body</H1>"),
		gonsole.WithUserPasswords(map[string]string{"xmli": "123456"}),
		gonsole.WithEnablePProf(true),
		gonsole.WithDeadlockIgnores([]string{
			"github.com/lixianmin/logo.(*Logger).goLoop",
			"github.com/lixianmin/road.(*App).goLoop",
			"github.com/lixianmin/road.(*sessionSender).goLoop",
		}),
	)

	var srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: mux,
	}

	loom.Go(goLoop)
	log.Fatal(srv.ListenAndServe())
}

func goLoop(later loom.Later) {
	var lock sync.Mutex
	f1(&lock)
}

func f1(lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()
	f2(lock)
}

func f2(lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()
	f1(lock)
}
