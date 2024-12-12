package controllers

import (
	utils "website/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type FlagController struct {
	beego.Controller
}

func (c *FlagController) Post() {
	cookie, cookieExists := c.GetSecureCookie(utils.Const.COOKIE_SECRET, "items")
	ids := utils.GetItemIds(cookie, cookieExists)
	if len(ids) == len(utils.Const.KEY_PAGES) {
		c.Data["json"] = map[string]string{
			"flag": utils.Const.FLAG,
		}
		c.ServeJSON(true)
	} else {
		c.Ctx.Request.Response.StatusCode = 403
		c.Ctx.Request.Response.Status = "Forbidden"
	}
}

func (c *FlagController) Get() {
	c.Data["Title"] = "Место нахождения флага"
	cookie, cookieExists := c.GetSecureCookie(utils.Const.COOKIE_SECRET, "items")
	c.Data["Items"] = utils.GetItemExistance(cookie, cookieExists)
	items := utils.GetItemIds(cookie, cookieExists)
	c.Data["CanGetFlag"] = len(items) == len(utils.Const.KEY_PAGE_ITEMS)

	c.TplName = "flag.tpl"
	c.Layout = "layout.html"
}
