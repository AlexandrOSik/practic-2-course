package routers

import (
	"website/controllers"
	"website/utils"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/"+utils.Const.ACT_PAGE, &controllers.ActController{})
	beego.Router("/"+utils.Const.FLAG_PAGE, &controllers.FlagController{})
	beego.Router("/place/:page", &controllers.PageController{})
}
