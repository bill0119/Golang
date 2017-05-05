package webServer

import (
	"github.com/astaxie/beego"
	"strconv"
)

type BeeferController struct {
	beego.Controller
}

func (c *BeeferController) Get() {
	c.Ctx.WriteString("Hello Beego")
}

func init() {

}

func StopServer() {

}

func StartServer(conf *WebConfig) {
	beego.Router("/", &BeeferController{})
	beego.Run(":" + strconv.Itoa(conf.HttpPort))
}
