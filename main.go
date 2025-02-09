package main

import (
	_ "website/routers"
	"website/utils"
	_ "website/utils"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	println(utils.Const.FLAG)
	beego.Run("0.0.0.0:8000")
}
