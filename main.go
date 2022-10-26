package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/ttacon/chalk"
	"reaven-server/config"
	"reaven-server/src/controller"
	_ "reaven-server/src/dao"
)

var port string
var mode string

func init() {
	config.ReadOptions()
	port = config.ConfigInfo.ServerPort
	mode = config.ConfigInfo.Mode
	fmt.Println(chalk.Green, "The service will be started on port"+port)
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	gin.SetMode(mode)
}

func main() {
	r := gin.Default() //携带基础中间件启动
	controller.Router(r)
	err := r.Run(port)
	if err != nil {
		fmt.Println(err.Error())
	}
}
