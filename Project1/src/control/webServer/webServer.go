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
	c.Data["Website"] = "Bill Website"
    c.Data["Email"] = "bill0119@gmail.com"
    c.TplName = "html/index.html"
}

func init() {

}

func StartServer(conf *WebConfig) {
	fmt.Printf("Web server is listening port %d\n", conf.HttpPort)

	//router
	beego.Router("/", &BeeferController{})

	//static
	beego.SetStaticPath("/js", "views/js")

	beego.Run(":" + strconv.Itoa(conf.HttpPort))
}
