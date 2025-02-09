package controllers

import (
	"website/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Post() {
	c.TplName = "index-ok.tpl"
	c.Layout = "layout.html"
	cookie, cookieExists := c.GetSecureCookie(utils.Const.COOKIE_SECRET, "items")
	c.Data["Items"] = utils.GetItemExistance(cookie, cookieExists)
	flag := c.GetString("flag")
	if flag == utils.Const.FLAG {
		c.Data["Title"] = "Конец пути"
		c.Data["FlagCorrect"] = true
		c.Data["Flag"] = flag
	} else {
		c.Data["Title"] = "Начало пути"
		c.Data["FlagCorrect"] = false
	}
}

func (c *MainController) Get() {
	c.Data["Title"] = "Начало пути"
	c.Data["Link"] = "/place/" + utils.Const.PAGES[0].Name

	cookie, cookieExists := c.GetSecureCookie(utils.Const.COOKIE_SECRET, "items")
	c.Data["Items"] = utils.GetItemExistance(cookie, cookieExists)

	c.TplName = "index.tpl"
	c.Layout = "layout.html"
}
