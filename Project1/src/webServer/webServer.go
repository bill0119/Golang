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

func StartServer(httpPort int) {
	beego.Router("/", &BeeferController{})
	beego.Run(":" + strconv.Itoa(httpPort))
}
