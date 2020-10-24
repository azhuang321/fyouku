package controllers

import (
	"demo/models"
	"demo/utils"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// UserController ...
type UserController struct {
	beego.Controller
}

// SaveRegister 用户注册
func (c *UserController) SaveRegister() {
	var (
		mobile   string
		password string
		// err      error
	)
	mobile = c.GetString("mobile")
	password = c.GetString("password")

	// 判空
	if mobile == "" {
		c.Data["json"] = ReturnError(4001, "手机号不能为空")
		c.ServeJSON()
		return
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		c.Data["json"] = ReturnError(4002, "手机号格式不正确")
		c.ServeJSON()
		return
	}
	if password == "" {
		c.Data["json"] = ReturnError(4003, "密码不能为空")
		c.ServeJSON()
		return
	}

	// 判断手机号是否已经注册
	status := models.IsUserMobile(mobile)
	if status {
		c.Data["json"] = ReturnError(4005, "此手机号已经注册")
		c.ServeJSON()
		return
	}
	id, err := models.UserSave(mobile, MD5V(password))
	if err != nil {
		c.Data["json"] = ReturnError(5000, err)
		c.ServeJSON()
		return
	}
	c.Data["json"] = ReturnSuccess(0, "注册成功", id, 0)
	c.ServeJSON()
	return
}

// LoginDo 用户登录
func (c *UserController) LoginDo() {
	mobile := c.GetString("mobile")
	password := c.GetString("password")

	// 入参校验
	if mobile == "" {
		c.Data["json"] = ReturnError(4001, "手机号不能为空")
		c.ServeJSON()
		return
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		c.Data["json"] = ReturnError(4002, "手机号格式不正确")
		c.ServeJSON()
		return
	}
	if password == "" {
		c.Data["json"] = ReturnError(4003, "密码不能为空")
		c.ServeJSON()
		return
	}

	uid, name := models.IsMobileLogin(mobile, MD5V(password))
	if uid == 0 {
		c.Data["json"] = ReturnError(4004, "手机号或密码不正确")
		c.ServeJSON()
		return
	}

	c.Data["json"] = ReturnSuccess(0, "登录成功", map[string]interface{}{
		"uid":      uid,
		"username": name,
	}, 1)
	c.ServeJSON()
	return
}

// 批量发送通知消息
func (c *UserController) SendMessageDo() {
	uids := c.GetString("uids")
	content := c.GetString("content")
	if uids == "" {
		c.Data["json"] = ReturnError(4001, "请填写接收人~")
		c.ServeJSON()
		return
	}
	if content == "" {
		c.Data["json"] = ReturnError(4001, "请填写发送内容")
		c.ServeJSON()
		return
	}
	messageId, err := models.SendMessageDo(content)
	if err != nil {
		c.Data["json"] = ReturnError(5000, "发送失败, 请联系客服")
		c.ServeJSON()
		return
	}
	uidConfig := strings.Split(uids, ",")
	for _, v := range uidConfig {
		userId, _ := strconv.Atoi(v)
		models.SendMessageUser(userId, messageId)
	}
	c.Data["json"] = ReturnSuccess(0, "发送成功~", "", 1)
	c.ServeJSON()
	return
}

// 上传视频文件
func (c *UserController) UploadVideo() {
	var (
		err   error
		title string
	)
	r := *c.Ctx.Request
	// 获取表单提交的数据
	uid := r.FormValue("uid")
	// 获取文件流
	file, header, _ := r.FormFile("file")
	// 转换文件流为二进制
	b, _ := ioutil.ReadAll(file)
	// 生成文件名
	filename := strings.Split(header.Filename, ".")
	filename[0] = utils.GetVideoName(uid)
	// 文件保存的位置
	var fileDir = "/crr/golang/dome/static/video/" + filename[0] + "." + filename[1]
	// 播放地址
	var playUrl = "/static/video/" + filename[0] + "." + filename[1]
	err = ioutil.WriteFile(fileDir, b, 0777)
	if err == nil {
		title = utils.ReturnSuccess(0, playUrl, nil, 1)
	} else {
		title = utils.ReturnError(5000, "上传失败, 请联系客服")
	}
	c.Ctx.WriteString(title)
}
