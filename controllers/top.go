package controllers

import (
	"demo/models"

	"github.com/astaxie/beego"
)

type TopController struct {
	beego.Controller
}

// 根据频道获取排行
func (c *TopController) ChannelTop() {
	// 获取频道
	channelId, err := c.GetInt("channelId")

	if channelId == 0 || err != nil {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
		return
	}
	num, videos, err := models.RedisGetChannelTop(channelId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
		return
	}

	c.Data["json"] = ReturnSuccess(0, "success", videos, num)
	c.ServeJSON()
	return
}

// 根据类型获取排行榜
func (c *TopController) TypeTop() {
	typeId, err := c.GetInt("typeId")
	if typeId == 0 || err != nil {
		c.Data["json"] = ReturnError(4004, "必须指定类型")
	}
	num, videos, err := models.RedisGetTypeTop(typeId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", videos, num)
	c.ServeJSON()
	return
}
