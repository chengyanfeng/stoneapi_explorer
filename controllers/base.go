package controllers

import (
	"github.com/astaxie/beego"
	"lian/util"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "tpl/header.html"
	c.Layout = "home.html"
}

func PagerHtml(totalcount int, page_size int, curpage int, mpurl string, number int) string {
	if number%2 == 0 {
		return "请输入为奇数的页码数"
	}
	totalpage := 0
	if totalcount%page_size > 0 {
		totalpage = (totalcount / page_size) + 1
	} else {
		totalpage = (totalcount / page_size)
	}
	//如果返回为空，那就返回为空
	if totalcount == 0 {
		return ""
	}
	html := ""
	//如果返回的条数大于每页的条数才能产生分页
	if totalcount > page_size {
		html = "<ul class=" + "\"" + "pagination" + "\"" + ">" +
			"<li><a>" + util.ToString(totalcount) + "条</a></li>" +
			"<li><a>" + util.ToString(totalpage) + "页</a></li>"
		if curpage > 1 {
			html = html +
				"<li>" +
				"<a href=?page=" + util.ToString(curpage-1) + "&explorerSearch=" + mpurl + " aria-label=" + "\"" + "Previous" + "\"" + ">" +
				"<span aria-hidden=" + "\"" + "true" + "\"" + ">&laquo;</span>" +
				"</a> </li>"
		}
		if totalpage <= number {
			for i := 1; i <= totalpage; i++ {
				if i == curpage {
					html = html + "<li class=" + "\"" + "active" + "\"" + "><a href=?page=" + util.ToString(i) + "&explorerSearch=" + mpurl + " >" + util.ToString(i) + "</a></li>"
				} else {
					html = html + "<li><a href=?page=" + util.ToString(i) + "&explorerSearch=" + mpurl + ">" + util.ToString(i) + "</a></li>"
				}
			}
		} else {
			if curpage > 1 {
				if curpage < totalpage {
					start := curpage - (number-1)/2
					end := curpage + (number-1)/2
					length := 0
					if start < 1 {
						length = 1 - start
						start = 1

						if (end + length) > totalpage {
							length = length - ((end + length) - totalpage)
						}
					}
					if (end + length) > totalpage {
						length = length - ((end + length) - totalpage)
					}
					for i := start; i <= end+length; i++ {
						if i == curpage {
							html = html + "<li class=" + "\"" + "active" + "\"" + "><a href=?page=" + util.ToString(i) + "&explorerSearch=" + mpurl + ">" + util.ToString(i) + "</a></li>"
						} else {
							html = html + "<li><a href=?page=" + util.ToString(i) + "&explorerSearch=" + mpurl + ">" + util.ToString(i) + "</a></li>"
						}
					}
				} else {
					for i := curpage - number + 1; i <= curpage; i++ {
						if i == curpage {
							html = html + "<li class=" + "\"" + "active" + "\"" + "><a href=?page=" + util.ToString(i) + "&explorerSearch=" + mpurl + ">" + util.ToString(i) + "</a></li>"
						} else {
							html = html + "<li><a href=?page=" + util.ToString(i) + "&explorerSearch=" + mpurl + ">" + util.ToString(i) + "</a></li>"
						}
					}
				}
			} else {
				for i := curpage; i <= number; i++ {
					if i == curpage {
						html = html + "<li class=" + "\"" + "active" + "\"" + "><a href=?page=" + util.ToString(i) + "&explorerSearch=" + mpurl + ">" + util.ToString(i) + "</a></li>"
					} else {
						html = html + "<li><a href=?page=" + util.ToString(i) + "&explorerSearch=" + mpurl + ">" + util.ToString(i) + "</a></li>"
					}
				}
			}
		}

		if curpage < totalpage {
			html = html +
				"<li>" +
				"<a href=?page=" + util.ToString(curpage+1) + "&explorerSearch=" + mpurl + " aria-label=" + "\"" + "Next" + "\"" + ">" +
				"<span aria-hidden=" + "\"" + "true" + "\"" + ">&raquo;</span>" +
				"</a> </li>"
		}

		html = html + "</ul> </nav>"

	}

	return html

}
