package controllers

import (
	"redfind/util"
)

type MainController struct {
	BaseController
}

var queryp = util.P{}

func (c *MainController) Get() {
	//下面显示多少个页码框，一定要用奇数，3，5，7
	number := 5
	//page_size
	page_size := 10
	explorerSearch := c.GetString("explorerSearch")
	page := c.GetString("page")
	curlpage := util.ToInt(page)
	if curlpage < 1 {
		curlpage = 1
	}
	if len(explorerSearch) > 0 {
		queryp["data"] = explorerSearch
	} else {
		delete(queryp, "data")
	}
	datalist := util.D("test").Find(queryp).Page((curlpage-1)*page_size, page_size-1).AllData()
	totalcount := util.D("test").Find(queryp).Count()
	c.Data["page"] = PagerHtml(totalcount, page_size, curlpage, explorerSearch, number)
	c.Data["datalist"] = datalist
	c.TplName = "form.html"

}
