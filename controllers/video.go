package controllers

import (
	"demo/models"

	"github.com/astaxie/beego"
)

// VideoController ...
type VideoController struct {
	beego.Controller
}

// ChannelAdvert 获取顶部广告
func (c *VideoController) ChannelAdvert() {
	channelId, _ := c.GetInt("channelId")

	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须制定频道")
		c.ServeJSON()
		return
	}

	num, channels, err := models.GetChannelAdvert(channelId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "请求数据失败, 请稍后重试~")
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", channels, num)
	c.ServeJSON()
	return
}

// ChannelHotList 正在热播列表
func (c *VideoController) ChannelHotList() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
		return
	}

	num, videos, err := models.GetChannelHotList(channelId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", videos, num)
	c.ServeJSON()
	return
}

// ChannelRecommendRegionList 根据频道地区获取推荐的视频
func (c *VideoController) ChannelRecommendRegionList() {
	channelId, _ := c.GetInt("channelId")
	regionId, _ := c.GetInt("regionId")

	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
		return
	}
	if regionId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道地区")
		c.ServeJSON()
		return
	}
	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", videos, num)
	c.ServeJSON()
	return
}

// GetChannelRecomendTypeList 根据频道类型获取视频
func (c *VideoController) GetChannelRecomendTypeList() {
	channelId, _ := c.GetInt("channelId")
	typeId, _ := c.GetInt("typeId")

	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
		return
	}
	if typeId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道类型")
		c.ServeJSON()
		return
	}
	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", videos, num)
	c.ServeJSON()
	return
}

// ChannelVideo 视频筛选列表
func (c *VideoController) ChannelVideo() {
	// 获取频道ID
	channelId, _ := c.GetInt("channelId")
	// 获取频道地区ID
	regionId, _ := c.GetInt("regionId")
	// 获取频道类型ID
	typeId, _ := c.GetInt("typeId")
	// 获取状态
	end := c.GetString("end")
	// 获取排序
	sort := c.GetString("sort")
	// 获取页码
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
		return
	}
	if limit == 0 {
		limit = 12
	}

	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", videos, num)
	c.ServeJSON()
	return
}

// 获取视频详情
func (c *VideoController) VideoInfo() {
	videoId, _ := c.GetInt("videoId")
	if videoId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定视频ID")
		c.ServeJSON()
		return
	}

	video, err := models.RedisGetVideoInfo(videoId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "请求失败, 请稍后重试~")
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", video, 1)
	c.ServeJSON()
	return
}

// 获取视频剧集列表
func (c *VideoController) VideoEpisodesList() {
	videoId, _ := c.GetInt("videoId")
	if videoId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定视频ID")
		c.ServeJSON()
		return
	}
	num, episodes, err := models.GetVideoEpisodesList(videoId)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "请求失败, 请稍后重试~")
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", episodes, num)
	c.ServeJSON()
	return
}

// 我的视频管理
func (c *VideoController) UserVideo() {
	uid, _ := c.GetInt("uid")
	if uid == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定用户")
		c.ServeJSON()
		return
	}
	num, videos, err := models.GetUserVideo(uid)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
	}
	c.Data["json"] = ReturnSuccess(0, "success", videos, num)
	c.ServeJSON()
	return
}

// 保存用户上传视频信息
func (c *VideoController) VideoSave() {
	playUrl := c.GetString("playUrl")
	title := c.GetString("title")
	subTitle := c.GetString("subTitle")
	channelId, _ := c.GetInt("channelId")
	typeId, _ := c.GetInt("typeId")
	regionId, _ := c.GetInt("regionId")
	uid, _ := c.GetInt("uid")
	if uid == 0 {
		c.Data["json"] = ReturnError(4001, "请先登录")
		c.ServeJSON()
		return
	}
	if playUrl == "" {
		c.Data["json"] = ReturnError(4001, "视频地址不能为空")
		c.ServeJSON()
		return
	}
	err := models.SaveVideo(title, subTitle, channelId, regionId, typeId, playUrl, uid)
	if err != nil {
		c.Data["json"] = ReturnError(5000, err)
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "success", nil, 1)
	c.ServeJSON()
	return
}
