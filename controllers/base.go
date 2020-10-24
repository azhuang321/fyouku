package controllers

import (
	"demo/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// 获取频道地区列表
func (c *BaseController) ChannelRegion() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
		return
	}

	num, regions, err := models.GetChannelRegion(channelId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", regions, num)
	c.ServeJSON()
	return
}

// 获取频道类型列表
func (c *BaseController) ChannelType() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
		return
	}
	num, regions, err := models.GetChannelType(channelId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", regions, num)
	c.ServeJSON()
	return
}
