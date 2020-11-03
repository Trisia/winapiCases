package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"net"
	"net/http"
	"os"
)

type program struct {
	server *http.Server
}

func (p *program) Start(s service.Service) error {
	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	p.server = &http.Server{Handler: r}
	l, err := net.Listen("tcp", ":30000")
	if err != nil {
		return err
	}
	go p.server.Serve(l)
	return nil
}

func (p *program) Stop(s service.Service) error {
	if p.server != nil {
		p.server.Shutdown(context.TODO())
	}
	return nil
}

func main() {
	pgm := &program{}
	cfg := &service.Config{
		Name:        "TestWebService",
		DisplayName: "TestWebService",
		Description: "测试Web服务",
	}
	s, err := service.New(pgm, cfg)
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		_ = service.Control(s, os.Args[1])
		return
	}
	err = s.Run()
	if err != nil {
		panic(err)
	}
}
