package webServer

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type BeeferController struct {
	beego.Controller
}

type WebServer struct {
	BeeCtrl BeeferController
}

func (c *BeeferController) Get() {
	c.Ctx.WriteString("Hello Beego")
}

func init() {

}

func StopServer() {

}

func StartServer(conf *WebConfig) {
	fmt.Printf("Web server is listening port %d\n", conf.HttpPort)
	beego.Router("/", &BeeferController{})
	beego.Run(":" + strconv.Itoa(conf.HttpPort))
}
