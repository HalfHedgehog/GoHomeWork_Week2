/**
    @Author: qiyou_wu
    @CreateDate: 2022/4/3
    @Description:
**/
package Config

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

type Server struct {
	closeChan   chan int
	status      string
	MysqlConfig string
	RedisConfig string
}

func (obj *Server) Open() {
	obj.status = "open"
	obj.closeChan = make(chan int, 1)
	fmt.Printf("Server is open , MysqlConfig:%s , RedisConfig : %s \n", obj.MysqlConfig, obj.RedisConfig)
}

func (obj *Server) Close() {
	obj.status = "close"
	obj.closeChan <- 1
}

func (obj *Server) Start() {
	g := new(errgroup.Group)

	//启动mysql
	g.Go(func() error {
		time.Sleep(1 * time.Second)
		obj.MysqlConfig = "127.0.0.1:3306"
		return nil
	})
	//启动Redis
	g.Go(func() error {
		time.Sleep(2 * time.Second)
		obj.RedisConfig = "127.0.0.1:6379"
		return nil
	})

	err := g.Wait()
	if err != nil {
		fmt.Println(err)
	}
	obj.Open()

	select {
	case <-obj.closeChan:
		obj.RedisConfig = "close"
		obj.MysqlConfig = "close"
		fmt.Printf("I'm closed , Mysql : %s , Redis : %s", obj.MysqlConfig, obj.RedisConfig)
	}
}
