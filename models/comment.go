package models

import (
	"demo/services/mq"
	"encoding/json"
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
		/*
			更新排行榜 - 通过MQ来来实现
			排行榜是根据评论数排序的, 把video_id推给队列, 表示该视频评论数需要+1.
			这里实现的是生产者, 消费者是/mq/top/main.go
			* 特别说明:用队列的异步方式是为了避免在接口中直接执行Mysql的耗时操作, 所有上面的Mysql操作应该挪到异步任务中执行
		*/
		// 实时增加评论数
		videoObj := map[string]int{
			"videoId": videoId,
		}
		videoJson, _ := json.Marshal(videoObj)
		mq.Publish("", "fyouku_top", string(videoJson))

		/*
			延时增加评论数
			数据推给死信队列的A交换机
			这里实现的是生产者, 消费者是/mq/comment/main.go
		*/
		videoCountObj := map[string]int{
			"VideoId":    videoId,
			"EpisodesId": episodesId,
		}
		videoCountJson, _ := json.Marshal(videoCountObj)
		mq.PublishDlx("fyouku.comment.count", string(videoCountJson))
	}
	return err
}
