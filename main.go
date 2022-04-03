package main

import (
	"GoHomeWork_week2/Config"
	"time"
)

func main() {
	go Config.Serve.Start()
	time.Sleep(5 * time.Second)
	Config.Serve.Close()
	time.Sleep(5 * time.Second)
}
