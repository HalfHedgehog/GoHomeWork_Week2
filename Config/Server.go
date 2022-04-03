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
	status      string
	MysqlConfig string
	RedisConfig string
}

func (obj *Server) Open() {
	obj.status = "open"
	fmt.Printf("Server is open , MysqlConfig:%s , RedisConfig : %s", obj.MysqlConfig, obj.RedisConfig)
}

func (obj *Server) Close() {
	obj.status = "close"
}

func (obj *Server) Start() {
	g := new(errgroup.Group)

	g.Go(func() error {
		time.Sleep(1 * time.Second)
		obj.RedisConfig = "RedisConfig"
		return nil
	})
	g.Go(func() error {
		time.Sleep(2 * time.Second)
		obj.RedisConfig = "RedisConfig"
		return nil
	})

	err := g.Wait()
	if err != nil {
		fmt.Println(err)
	}
	obj.Open()
}
