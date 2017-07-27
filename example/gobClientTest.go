package main

import (
	"encoding/gob"
	"log"

	"github.com/Greyh4t/teleport"
	"github.com/Greyh4t/teleport/debug"
)

var tp = teleport.New()

type X struct {
	Name string
}

type Y struct {
	Name string
	Z    Z
}

type Z struct {
	Name string
}

func init() {
	//注册interface{}里存放的数据，否则会反序列化失败
	gob.Register(X{})
	gob.Register(Y{})
	gob.Register(Z{})
}

func main() {
	debug.Debug = true
	tp.SetUID("clientUID", "serverUID").SetAPI(teleport.API{
		"report": new(report),
	})
	tp.Client("127.0.0.1", ":20125")
	tp.Request(X{"x报道"}, "report", "flag", "serverUID")
	var z = Z{"z报道"}
	tp.Request(Y{"y报道", z}, "report", "flag", "serverUID")
	tp.Request(Z{"z报道"}, "report", "flag", "serverUID")
	select {}
}

type report struct{}

func (*report) Process(receive *teleport.NetData) *teleport.NetData {
	if receive.Status == teleport.SUCCESS {
		log.Printf("%v", receive.Body)
	}
	if receive.Status == teleport.FAILURE {
		log.Printf("%v", "请求处理失败！")
	}
	return nil
}
