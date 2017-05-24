package webServer

import (
	"fmt"
	"os"
	"io/ioutil"
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
	execDirAbsPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(execDirAbsPath + "/web/html/index.html")
	if err != nil {
		panic(err)
	}
	c.Ctx.WriteString(string(data))
}

func init() {

}

func StartServer(conf *WebConfig) {
	fmt.Printf("Web server is listening port %d\n", conf.HttpPort)

	//router
	beego.Router("/", &BeeferController{})


	beego.Run(":" + strconv.Itoa(conf.HttpPort))
}
