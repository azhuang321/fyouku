package models

import "github.com/astaxie/beego/orm"

type Region struct {
	Id   int
	Name string
}
type Type struct {
	Id   int
	Name string
}

func GetChannelRegion(channelId int) (int64, []Region, error) {
	o := orm.NewOrm()
	var regions []Region
	sql := "select id,name " +
		"from channel_region " +
		"where status=1 and channel_id=? " +
		"order by sort desc"
	num, err := o.Raw(sql, channelId).QueryRows(&regions)
	return num, regions, err
}

func GetChannelType(channelId int) (int64, []Type, error) {
	o := orm.NewOrm()
	var types []Type
	sql := `select id,name
	 from channel_type
	 where status=1 and channel_id=?
	 order by sort desc`
	num, err := o.Raw(sql, channelId).QueryRows(&types)
	return num, types, err
}
