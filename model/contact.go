package model

import "time"

const (
	CONCAT_CATE_USER     = 0x01 //用户
	CONCAT_CATE_COMUNITY = 0x02 //群组
)

type Contact struct {
	Id       int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Ownerid  int64     `xorm:"bigint(20)" form:"ownerid" json:"ownerid"` // 账号角色
	Dstobj   int64     `xorm:"bigint(20)" form:"dstobj" json:"dstobj"`   // 对应好友
	Cate     int       `xorm:"int(11)" form:"cate" json:"cate"`          // 什么角色
	Memo     string    `xorm:"varchar(120)" form:"memo" json:"memo"`     // 什么角色
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"` // 什么角色
}
