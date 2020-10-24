package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id          int
	Content     string
	AddTime     int64
	UserId      int
	Stamp       int
	Status      int
	PraiseCount int
	EpisodesId  int
	VideoId     int
}

func init() {
	orm.RegisterModel(new(Comment))
}

func GetCommentList(episodesId int, offset int, limit int) (int64, []Comment, error) {
	o := orm.NewOrm()
	var comments []Comment
	var sql string
	sql = "select id from comment where status=1 and episodes_id=?"
	num, _ := o.Raw(sql, episodesId).QueryRows(&comments)
	sql = "select id,content,add_time,user_id,stamp,praise_count,episodes_id" +
		" from comment" +
		" where status=1 and episodes_id=?" +
		" order by add_time desc limit ?,?"
	_, err := o.Raw(sql, episodesId, offset, limit).QueryRows(&comments)
	return num, comments, err
}

func SaveComment(content string, uid int, episodesId int, videoId int) error {
	o := orm.NewOrm()
	var sql string
	var comment Comment
	comment.Content = content
	comment.UserId = uid
	comment.EpisodesId = episodesId
	comment.VideoId = videoId
	comment.Stamp = 0
	comment.Status = 1
	comment.AddTime = time.Now().Unix()
	_, err := o.Insert(&comment)
	if err == nil {
		// 修改视频的总评论数
		sql = "update video set comment=comment+1 where id=?"
		o.Raw(sql, videoId).Exec()
		// 修改视频剧集的评论数
		sql = "update video_episodes set comment=comment+1 where id=?"
		o.Raw(sql, episodesId).Exec()
	}
	return err
}
