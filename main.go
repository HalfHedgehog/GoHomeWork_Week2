package main

import (
	"GoHomeWork_week2/metadata"
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go func() {
		time.Sleep(5 * time.Second)

		metadata.MetaData.Source = "我已经被修改"

		ch1 <- 1
	}()

	go func() {
		fmt.Println(metadata.MetaData.Source)
		ch2 <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case <-ch1:
			fmt.Println("goroutine update is Done")
		case <-ch2:
			fmt.Println("goroutine select is Done")
		}
	}
}
