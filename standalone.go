package main

import (
	"sync"
	"github.com/tikalk/go-distribution-workshop/apps"
	"github.com/tikalk/go-distribution-workshop/messaging"
)

func Standalone(){
	defer messaging.Stop()
	wg := &sync.WaitGroup{}
	wg.Add(2)


	go apps.ExecuteSimulation(wg)
	go apps.ExecuteDisplay(wg)

	wg.Wait()

}
