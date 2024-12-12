package controllers

import (
	"slices"
	"strings"
	utils "website/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type ActController struct {
	beego.Controller
}

func (c *ActController) Post() {
	cookie, cookieExists := c.GetSecureCookie(utils.Const.COOKIE_SECRET, "items")
	ids := utils.GetItemIds(cookie, cookieExists)
	check := c.GetString("check")
	value := c.GetString("value")
	idx := slices.Index(utils.Const.KEY_PAGE_SYMBOL, value)
	if idx >= 0 {
		if utils.Const.KEY_PAGE_CHECK[idx] == check {
			if !slices.Contains(ids, value) {
				ids = append(ids, value)
			}
			data := strings.Join(ids, ",")
			c.Ctx.SetSecureCookie(utils.Const.COOKIE_SECRET, "items", data)
			c.Ctx.WriteString(strings.Join(ids, ","))
			return
		} else {
			c.Ctx.Request.Response.StatusCode = 403
			c.Ctx.Request.Response.Status = "Forbidden"
		}
	} else {
		c.Ctx.Request.Response.StatusCode = 404
		c.Ctx.Request.Response.Status = "Not Found"
	}

}
