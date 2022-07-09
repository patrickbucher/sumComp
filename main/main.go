package main

import (
	"fmt"
	"time"

	"github.com/patrickbucher/sumcomp"
)

func main() {
	cache := sumcomp.NewCache(100)
	erica := time.NewTicker(4 * time.Second)
	koni := time.NewTicker(3 * time.Second)
	randy := time.NewTicker(2 * time.Second)
	bitchard := time.NewTicker(1 * time.Second)
	bankrupty := time.NewTicker(1 * time.Minute)
	write := func(name string, c *sumcomp.Cache) {
		summary := sumcomp.RandomSummary()
		fmt.Printf("%s just summarized %s\n", name, summary)
		cache.Publish(summary)
	}
	read := func(name string, c *sumcomp.Cache) {
		dataId, err := cache.GetRandomDataId()
		if err != nil {
			return
		}
		summary, err := cache.GetSummary(dataId)
		if err != nil {
			return
		}
		fmt.Printf("%s just read %s\n", name, summary)
	}
	go func() {
		for {
			select {
			case <-erica.C:
				write("Erica", cache)
			case <-koni.C:
				write("Koni", cache)
			case <-randy.C:
				read("Randy", cache)
			case <-bitchard.C:
				read("Andy", cache)
			}
		}
	}()
	<-bankrupty.C
	fmt.Println("sumComp just went bust")
}
