package controllers

import (
	"stoneapi_explorer/util"
	"stoneapi_explorer/models"
	"math"
	"fmt"
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
	datalist := &[]models.Node{}

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
	totalcount := util.D("uploads",mongp).Find(queryp).Count()
	a:=math.Ceil(util.ToFloat(totalcount)/util.ToFloat(page_size))
	fmt.Print(a)
	if curlpage<util.ToInt(math.Ceil(util.ToFloat(totalcount)/util.ToFloat(page_size))){
		datalist = util.D("uploads",mongp).Find(queryp).Page(totalcount-(curlpage)*page_size, page_size).AllData()
	}else {
		if totalcount%page_size==0{
			datalist = util.D("uploads",mongp).Find(queryp).Page(0, page_size).AllData()
		}else {
		datalist = util.D("uploads",mongp).Find(queryp).Page(0, (totalcount%page_size)).AllData()
		}
	}

	for k, v := range *datalist {
		v.Number = (totalcount - (curlpage)*page_size + k + 1)
		*reurnDataList = append(*reurnDataList, v)
	}
	c.Data["json"] = map[string]interface{}{"totalcount": totalcount, "data": reurnDataList, "culpage": page}
	c.ServeJSON()
	return

}
