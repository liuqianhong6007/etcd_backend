package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/etcd_backend/internal"
)

func main() {
	// 读取配置
	internal.ReadConf()

	// 初始化 etcd
	internal.Init(internal.GetConfig().EtcdAddr)

	// 开启 http 服务
	r := gin.Default()
	internal.RegisterRoute(r)
	serverAddr := fmt.Sprintf("%s:%d", internal.GetConfig().Host, internal.GetConfig().Port)
	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}
