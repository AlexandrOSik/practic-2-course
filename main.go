package main

import (
	_ "website/routers"
	_ "website/utils"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run("0.0.0.0:8000")
}
