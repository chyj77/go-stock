package controllers

import (
	"encoding/json"
	"fmt"
	"go-stock/models"

	"github.com/astaxie/beego"
)

type ZtsjController struct {
	beego.Controller
}

func (c *ZtsjController) URLMapping() {
	c.Mapping("Save", c.Save)
	c.Mapping("Index", c.Index)
	c.Mapping("Delete", c.Delete)
}

// @router /ztsj/index [get]
func (this *ZtsjController) Index() {
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", this.Ctx.Request.Header.Get("Origin"))

	var pageIndex int
	var pageSize int
	var ztsjs models.ZtsjResponse
	this.Ctx.Input.Bind(&pageIndex, "pageIndex")
	this.Ctx.Input.Bind(&pageSize, "pageSize")
	if pageSize == 0 {
		pageSize = 10
	}
	if pageIndex > 0 {
		pageIndex = (pageIndex - 1) * pageSize
	}
	ztsjs.PageIndex = pageIndex
	ztsjs.PageSize = pageSize
	ztsjs.TotalCounts, ztsjs.Datas = models.QueryZtsj(pageIndex, pageSize)
	this.Data["json"] = ztsjs
	this.ServeJSON()
}

// @router /ztsj/save [post]
func (this *ZtsjController) Save() {
	var ztsj models.Ztsj
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ztsj); err == nil {
		counts := models.SaveZtsj(&ztsj)
		if counts > 0 {
			this.Data["json"] = "{\"code\":200}"
		} else {
			this.Data["json"] = "{\"code\":400}"
		}
	} else {
		fmt.Println(err.Error())
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

// @router /ztsj/delete [get]
func (this *ZtsjController) Delete() {
	recId, _ := this.GetInt64("recId")

	counts := models.DeleteZtsj(recId)
	if counts > 0 {
		this.Data["json"] = "{\"code\":200}"
	} else {
		this.Data["json"] = "{\"code\":400}"
	}

	this.ServeJSON()
}
