package controllers

import (
	"lian/util"
	"lian/models"
)

type MainController struct {
	BaseController
}

var queryp = util.P{}

func (c *MainController) Get() {
	c.TplName = "form.html"

}
//获取数据分页
func (c *MainController) Getdata() {
	page_size := 10
	reurnDataList := &[]models.Node{}
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
	totalcount := util.D("uploads").Find(queryp).Count()
	datalist := util.D("uploads").Find(queryp).Page(totalcount-(curlpage)*page_size, page_size).AllData()

	for k, v := range *datalist {
		v.Number = (totalcount - (curlpage)*page_size + k + 1)
		*reurnDataList = append(*reurnDataList, v)
	}
	c.Data["json"] = map[string]interface{}{"totalcount": totalcount, "data": reurnDataList, "culpage": page}
	c.ServeJSON()
	return

}
