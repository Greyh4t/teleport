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
	tp.SetUID("serverUID").SetAPI(teleport.API{
		"report": new(report),
	}).Server(":20125")
	select {}
}

type report struct{}

func (*report) Process(receive *teleport.NetData) *teleport.NetData {
	log.Printf("报到：%v", receive.Body)
	return teleport.ReturnData("服务器：" + receive.From + " 客户端已经报到！")
}
