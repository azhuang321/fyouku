package controllers

import (
	"demo/models"

	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"statmp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userinfo"`
}

// 获取评论列表
func (c *CommentController) List() {
	// 获取剧集数
	episodesId, _ := c.GetInt("episodesId")
	// 获取页码信息
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	if episodesId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定视频剧集")
	}
	if limit == 0 {
		limit = 12
	}

	num, comments, err := models.GetCommentList(episodesId, offset, limit)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
		return
	}

	var data []CommentInfo
	var commentInfo CommentInfo
	for _, v := range comments {
		commentInfo.Id = v.Id
		commentInfo.Content = v.Content
		commentInfo.AddTime = v.AddTime
		commentInfo.AddTimeTitle = DateFormat(v.AddTime)
		commentInfo.UserId = v.UserId
		commentInfo.Stamp = v.Stamp
		commentInfo.PraiseCount = v.PraiseCount
		// 获得用户信息
		commentInfo.UserInfo, _ = models.RedisGetUserInfo(v.UserId)
		data = append(data, commentInfo)
	}
	c.Data["json"] = ReturnSuccess(0, "success", data, num)
	c.ServeJSON()
	return
}

// 保存评论
func (c *CommentController) Save() {
	content := c.GetString("content")
	uid, _ := c.GetInt("uid")
	episodesId, _ := c.GetInt("episodesId")
	videoId, _ := c.GetInt("videoId")
	if content == "" {
		c.Data["json"] = ReturnError(4001, "内容不能为空")
		c.ServeJSON()
		return
	}
	if uid == 0 {
		c.Data["json"] = ReturnError(4002, "请先登录")
		c.ServeJSON()
		return
	}
	if episodesId == 0 {
		c.Data["json"] = ReturnError(4003, "必须指定评论剧集")
		c.ServeJSON()
		return
	}
	if videoId == 0 {
		c.Data["json"] = ReturnError(4004, "必须指定视频ID")
		c.ServeJSON()
		return
	}
	err := models.SaveComment(content, uid, episodesId, videoId)
	if err != nil {
		c.Data["json"] = ReturnError(5000, err)
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", "", 1)
	c.ServeJSON()
	return
}
