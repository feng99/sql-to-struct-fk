package entity

type EbLiveRoom struct {
	Id                   int64  `gorm:"id"`                     // id
	AnchorId             int64  `gorm:"anchor_id"`              // 主播ID
	Name                 string `gorm:"name"`                   // 直播间名字
	Title                string `gorm:"title"`                  // 直播间标题
	CoverImg             string `gorm:"cover_img"`              // 背景图
	ShareImg             string `gorm:"share_img"`              // 分享图
	StartTime            int64  `gorm:"start_time"`             // 直播计划开始时间
	EndTime              int64  `gorm:"end_time"`               // 直播计划结束时间
	LiveType             int64  `gorm:"live_type"`              // 直播间类型 【1: 推流，0：手机直播】
	ScreenType           int64  `gorm:"screen_type"`            // 横屏、竖屏 【1：横屏，0：竖屏】
	CloseLike            int64  `gorm:"close_like"`             // 是否关闭点赞 1 开 0 关
	CloseGoods           int64  `gorm:"close_goods"`            // 是否关闭货架 1 开 0 关
	CloseComment         int64  `gorm:"close_comment"`          // 是否关闭评论 1 开 0 关
	ErrorMsg             string `gorm:"error_msg"`              // 未通过原因
	Status               int64  `gorm:"status"`                 // 审核状态:0-未审核 1-微信审核 2-审核通过 -1-审核未通过
	LiveStatus           int64  `gorm:"live_status"`            // 直播状态:0-未开始 1-直播中 2-已结束 3-超时未开播
	VisibleRange         int64  `gorm:"visible_range"`          // 可见范围:1-全员可见 2-指定用户可见
	VisibleRangeConf     string `gorm:"visible_range_conf"`     // 可见范围配置 visible_range = 1 时存场景 ID 列表 visible_range = 2 时存中青 UID 列表 逗号分隔
	OrderPrompt          int64  `gorm:"order_prompt"`           // 下单提示条 0-未设置 1-真实交易 2-真实+水军
	InvitationList       int64  `gorm:"invitation_list"`        // 邀请榜 0-未设置 1-真实用户 2-真实+水军
	InvitationListNumber int64  `gorm:"invitation_list_number"` // 邀请榜展示的用户数
	OnlineAudience       int64  `gorm:"online_audience"`        // 在线观众 0-未设置 1-真实用户 2-真实+水军
	OnlineAudienceNumber int64  `gorm:"online_audience_number"` // 在线观众展示数
	Mark                 string `gorm:"mark"`                   // 备注
	ServerAddress        string `gorm:"server_address"`         // 服务器地址
	PrivateKey           string `gorm:"private_key"`            // 流密钥
	ReplayStatus         int64  `gorm:"replay_status"`          // 回放状态
	Sort                 int64  `gorm:"sort"`                   // 直播间排序,多个直播间时使用
	MerchantId           int64  `gorm:"merchant_id"`            // 所属商家id
	OpenType             int64  `gorm:"open_type"`              // 开播类型: 0-主播端app 1-管理后台(obs)
	IsShow               int64  `gorm:"is_show"`                // 是否显示
	IsDelete             int64  `gorm:"is_delete"`              // 是否删除
	CreateTime           int64  `gorm:"create_time"`            // 创建时间
	UpdateTime           int64  `gorm:"update_time"`            // 修改时间

}
