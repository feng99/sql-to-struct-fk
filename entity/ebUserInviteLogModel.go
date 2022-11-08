package entity

type EbUserInviteLog struct {
	Id int64 `gorm:"id"` //主键,自增

	Account string `gorm:"account"` //中青uid

	BeInvitedUid string `gorm:"be_invited_uid"` //被邀请人中青uid

	BeInvitedOpenid string `gorm:"be_invited_openid"` //被邀请人微信openid

	RoomId int64 `gorm:"room_id"` //房间id

	CreateTime string `gorm:"create_time"` //添加时间

	UpdateTime string `gorm:"update_time"` //更新时间

	IsDelete int64 `gorm:"is_delete"` //是否删除 1:是:0否

	IsSign int64 `gorm:"is_sign"` //是否已发送过信令

	OriginAccount string `gorm:"origin_account"` //邀请链 根节点中青uid

}
