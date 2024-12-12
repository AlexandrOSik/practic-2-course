package controllers

import (
	"slices"
	utils "website/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type PageController struct {
	beego.Controller
}

func (c *PageController) Get() {
	c.Data["FlagExists"] = false
	c.Data["Flag"] = utils.Const.FLAG

	pageName := c.GetString(":page")
	pageIndex, pageExists := utils.Const.PAGE_INDEXES[pageName]
	if !pageExists {
		c.Ctx.Request.Response.StatusCode = 404
		c.Ctx.Request.Response.Status = "Not Found"
		return
	}

	page := utils.Const.PAGES[pageIndex]
	links := make([]map[string]string, len(page.Transitions))

	idx := slices.Index(utils.Const.KEY_PAGES, pageIndex)
	if idx >= 0 {
		cookie, cookieExists := c.GetSecureCookie(utils.Const.COOKIE_SECRET, "items")
		if !slices.Contains(utils.GetItemIds(cookie, cookieExists), utils.Const.KEY_PAGE_SYMBOL[idx]) {
			c.Data["Form"] = true
			c.Data["Item"] = utils.Const.KEY_PAGE_ITEMS[idx]
			c.Data["FormCheck"] = utils.Const.KEY_PAGE_CHECK[idx]
			c.Data["FormValue"] = utils.Const.KEY_PAGE_SYMBOL[idx]
			c.Data["ActLink"] = "/" + utils.Const.ACT_PAGE
		}
	}

	c.Data["Teller"] = page.Teller
	c.Data["Title"] = page.Title
	c.Data["Words"] = page.Words
	c.Data["Place"] = page.Place

	for i := 0; i < len(page.Transitions); i++ {
		to := utils.Const.PAGES[page.Transitions[i]]
		links[i] = map[string]string{
			"Text": to.Title,
			"Link": "/place/" + to.Name,
		}
	}
	if pageIndex == utils.Const.FLAG_PAGE_GATE {
		links = append(links, map[string]string{
			"Text": "Место нахождения флага",
			"Link": "/" + utils.Const.FLAG_PAGE,
		})
	}
	c.Data["Links"] = links

	cookie, cookieExists := c.GetSecureCookie(utils.Const.COOKIE_SECRET, "items")
	c.Data["Items"] = utils.GetItemExistance(cookie, cookieExists)

	c.TplName = "page.tpl"
	c.Layout = "layout.html"
}
