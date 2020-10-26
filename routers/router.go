package routers

import (
	"demo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{})

	// 用户注册
	beego.Router("/register/save", &controllers.UserController{}, "post:SaveRegister")
	// 用户登录
	beego.Router("/register/login", &controllers.UserController{}, "post:LoginDo")
	// 批量发送通知消息
	beego.Router("/user/SendMessageDo", &controllers.UserController{}, "post:SendMessageDo")
	// 上传视频
	beego.Router("/video/UploadVideo", &controllers.UserController{}, "post:UploadVideo")

	// 获取顶部广告
	beego.Router("/video/ChannelAdvert", &controllers.VideoController{}, "post:ChannelAdvert")
	// 正在热播列表
	beego.Router("/video/ChannelHotList", &controllers.VideoController{}, "post:ChannelHotList")
	// 根据频道地区获取推荐的视频
	beego.Router("/video/ChannelRecommendRegionList", &controllers.VideoController{}, "post:ChannelRecommendRegionList")
	// 根据频道类型获取视频
	beego.Router("/video/GetChannelRecomendTypeList", &controllers.VideoController{}, "post:GetChannelRecomendTypeList")
	// 视频筛选列表
	beego.Router("/video/ChannelVideo", &controllers.VideoController{}, "post:ChannelVideo")
	// 获取视频详情
	beego.Router("/video/VideoInfo", &controllers.VideoController{}, "post:VideoInfo")
	// 获取视频剧集列表
	beego.Router("/video/VideoEpisodesList", &controllers.VideoController{}, "post:VideoEpisodesList")
	// 我的视频管理
	beego.Router("/video/UserVideo", &controllers.VideoController{}, "post:UserVideo")

	// 获取频道地区列表
	beego.Router("/base/ChannelRegion", &controllers.BaseController{}, "post:ChannelRegion")
	// 获取频道类型列表
	beego.Router("/base/ChannelType", &controllers.BaseController{}, "post:ChannelType")

	// 获取评论列表
	beego.Router("/comment/List", &controllers.CommentController{}, "post:List")
	beego.Router("/comment/Save", &controllers.CommentController{}, "post:Save")

	// 根据频道获取排行榜
	beego.Router("/top/ChannelTop", &controllers.TopController{}, "post:ChannelTop")
	// 根据类型获取排行榜
	beego.Router("/top/TypeTop", &controllers.TopController{}, "post:TypeTop")

	// wsdemo
	beego.Router("/test/ws", &controllers.TestController{}, "get:WsFunc")
	beego.Router("/test/index", &controllers.TestController{}, "get:Get")

	// 弹幕
	beego.Router("/barrage/BarrageWs", &controllers.BarrageController{}, "get:BarrageWs")
	beego.Router("/barrage/Save", &controllers.BarrageController{}, "post:Save")

	// redis
	beego.Router("/redis/demo", &controllers.RedisDemoController{}, "post:Demo")

	// mq
	beego.Router("/mq/push", &controllers.MqDemoController{}, "post:Push")
	beego.Router("/mq/PushFanout", &controllers.MqDemoController{}, "post:PushFanout")
	beego.Router("/mq/PushDirect", &controllers.MqDemoController{}, "post:PushDirect")
	beego.Router("/mq/PushTopic", &controllers.MqDemoController{}, "post:PushTopic")
	beego.Router("/mq/PushTopicTwo", &controllers.MqDemoController{}, "post:PushTopicTwo")
	beego.Router("/mq/PushDlx", &controllers.MqDemoController{}, "post:PushDlx")
	beego.Router("/mq/PushTwoDlx", &controllers.MqDemoController{}, "post:PushTwoDlx")

}
