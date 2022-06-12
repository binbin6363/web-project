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

// EvalMatches .
type EvalMatches struct {
	Value  int               `json:"value"`
	Metric string            `json:"metric"`
	Tags   map[string]string `json:"tags"`
}

// AlertRequest grafana自定义webhook告警请求
type AlertRequest struct {
	Title       string        `json:"title"`       // title
	State       string        `json:"state"`       // state
	RuleUrl     string        `json:"ruleUrl"`     // ruleUrl
	RuleName    string        `json:"ruleName"`    // ruleName
	RuleId      int           `json:"ruleId"`      // ruleId
	OrgId       int           `json:"orgId"`       // orgId
	Message     string        `json:"message"`     // message
	DashboardId int           `json:"dashboardId"` // dashboardId
	EvalMatches []EvalMatches `json:"evalMatches"` // evalMatches
}

// SmsAlertReq UBS 提供的短信告警restapi接口
// 示例：curl -d '{"ip": "10.0.0.1", "source_time": "2022-05-09 16:43:53","alarm_type": "api_default",
// "level": "remain","alarm_name": "FAILURE for production/HTTP", "alarm_content":
// "FAILURE for production HTTP on machine 10.0.0.1", "meta_info": "结算中心", "action": "firing"}'
// http://paas.uat.rfzubs.com:80/o/cw_uac_saas/alarm/collect/event/api/75563987-d529-421d-bfaa-4bab4dfbdc8a/
// -H 'X-Secret:lAkTwemWYln1xzY3pvhtkyJ1Re3tBEIu' -v -k
type SmsAlertReq struct {
	Ip           string `json:"title"`         // 告警对象
	SourceTime   string `json:"source_time"`   // 告警发生的时间，格式：YYYY-MM-DD HH:mm:ss，必须
	AlarmType    string `json:"alarm_type"`    // 告警指标，必须
	Level        string `json:"level"`         // 告警级别，必须，枚举：remain/warning/fatal(提醒/警告/致命)
	AlarmName    string `json:"alarm_name"`    // 告警名称，可选
	AlarmContent string `json:"alarm_content"` // 告警详情，可选
	MetaInfo     string `json:"meta_info"`     // 告警其余信息，可选
	Action       string `json:"action"`        // 告警事件动作，可选，不传默认是firing, 枚举：firing/resolved(产生告警/消除告警)
	BkObjId      string `json:"bk_obj_id"`     // CMDB模型ID，可选,
	BkInstId     string `json:"bk_inst_id"`    // CMDB模型实例ID，可选
}
