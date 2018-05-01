package webServer

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type BeeferController struct {
	beego.Controller
}

type APIController struct {
    beego.Controller
}

type WebServer struct {
	BeeCtrl BeeferController
}

func (c *BeeferController) Get() {
	c.Data["Website"] = "Bill Website"
    c.Data["Email"] = WebData.Email
	c.TplName = "html/index.html"
}

func (c *APIController) Get() {
	fmt.Printf("Get API Test\n")
	c.Ctx.WriteString("hello")
}

func init() {

}

func StartServer(conf *WebConfig) {
	fmt.Printf("Web server is listening port %d\n", conf.HttpPort)

	//router
	beego.Router("/", &BeeferController{})
	beego.Router("/api", &APIController{})

	//static
	beego.SetStaticPath("/js", "views/js")
	beego.SetStaticPath("/css", "views/css")

	beego.Run(":" + strconv.Itoa(conf.HttpPort))
}
