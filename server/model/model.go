package model

type AddEmailReq struct {
	Subject     string `json:"subject"`     // 邮件标题
	Receiver    string `json:"receiver"`    // 邮件收件人
	Body        string `json:"body"`        // 邮件正文
	Creator     string `json:"creator"`     // 创建者ID
	Mode        string `json:"mode"`        // 邮件发送模式。随机/定向
	Public      bool   `json:"public"`      // 邮件是否公开可见
	TriggleDate string `json:"triggleDate"` // 邮件投递时间，日期
}

type AddEmailRsp struct {
	ID int `json:"id"` // 邮件ID
}

// 发送邮件需要的结构体
type EmailInfo struct {
	Receiver   string // 收件人
	CCReceiver string // 抄送人
	Subject    string // 邮件主题
	Body       string // 邮件内容
}

type QueryEmailReq struct {
	Creator string `json:"creator"` // 创建者ID
}

type EmailBasic struct {
	Receiver    string `json:"receiver"`     // 收件人
	Subject     string `json:"subject"`      // 邮件主题
	Mode        string `json:"mode"`         // 邮件发送模式。随机/定向
	Public      bool   `json:"public"`       // 邮件是否公开可见
	TriggleDate string `json:"triggle_date"` // 邮件投递时间，日期
	State       string `json:"state"`        // 邮件投递状态
}

type QueryEmailRsp struct {
	ID        int          `json:"id"`   // 邮件ID
	BasicList []EmailBasic `json:"list"` // 邮件列表，不包含内容
}

// ==================== gorm db 信息
// 邮件信息
type EmailTask struct {
	ID          int    `gorm:"column:id;primarykey;AUTO_INCREMENT"`
	Subject     string `gorm:"column:email_subject"`
	Body        string `gorm:"column:email_body"`
	Creator     string `gorm:"column:email_creator"`  // 创建者ID
	Receiver    string `gorm:"column:email_receiver"` // 收件人，非公开情况下使用
	Public      bool   `gorm:"column:email_public"`   // 是公开还是非公开
	ContentType string `gorm:"column:content_type"`   // 邮件正文类型。http还是plain
	Mode        string `gorm:"column:triggle_mode"`   // 邮件触发方式，随机还是指定用户
	State       int    `gorm:"column:email_state"`    // 邮件状态。0待发送，1已发送，2发送失败
	TriggleTime int    `gorm:"column:triggle_time"`   // 邮件触发时间戳
	CreateTime  int    `gorm:"column:create_time"`    // 邮件创建时间戳
	UpdateTime  int    `gorm:"column:update_time"`    // 邮件更新时间戳，用于任务触发设置时间
}

func (EmailTask) TableName() string {
	return "emails"
}
