package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
}

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

// ReturnSuccess 正确的返回值
func ReturnSuccess(code int, msg interface{}, items interface{}, count int64) *JsonStruct {
	return &JsonStruct{
		Code:  code,
		Msg:   msg,
		Items: items,
		Count: count,
	}

}

func ReturnError(code int, msg interface{}) *JsonStruct {
	return &JsonStruct{
		Code: code,
		Msg:  msg,
	}

}

// 用户密码加密
func MD5V(password string) string {
	h := md5.New()
	h.Write([]byte(password + beego.AppConfig.String("md5code")))
	str := hex.EncodeToString(h.Sum(nil))
	return str
}

// 格式化时间
func DateFormat(time1 int64) string {
	video_time := time.Unix(time1, 0)
	return video_time.Format("2006-01-02")
}
