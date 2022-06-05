package middleware

import (
	"context"
	"log"
	"time"

	"github.com/binbin6363/web-project/server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Dao struct {
	commDB *gorm.DB
}

var d *Dao

func GetDao() *Dao {
	return d
}

// New creates Dao instance
func New(dsn string) error {
	d = &Dao{}

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	cli, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("governance dao: New db client (mysql.ivy.governance.common) error(%v)", err)
		return err
	}
	d.commDB = cli
	return nil
}

func (d *Dao) db(ctx context.Context) *gorm.DB {
	return d.commDB
}

// CreateEmailTask 创建邮件任务
func (d *Dao) CreateEmailTask(ctx context.Context, taskInfo *model.EmailTask) (id int, err error) {
	r := d.db(ctx)

	log.Printf("CreateEmailTask:%v", taskInfo)
	res := r.Debug().Create(taskInfo)
	if res.Error != nil {
		log.Printf("CreateEmailTask read db error(%v) id(%d)", res.Error, id)
	}
	return taskInfo.ID, res.Error
}

// ScanEmailTask 扫描邮件任务，一次扫2条。因为内容可能比较大
func (d *Dao) ScanEmailTask(ctx context.Context) (taskList []*model.EmailTask, err error) {
	r := d.db(ctx)

	r = r.Debug().Where("email_state = 0").Where("triggle_time <= ?", time.Now().Unix())
	r = r.Order(clause.OrderByColumn{Column: clause.Column{Name: "triggle_time"}, Desc: false}).Limit(2)

	taskList = make([]*model.EmailTask, 0)
	if err = r.Debug().Find(&taskList).Error; err != nil {
		log.Printf("ScanEmailTask read db error(%v)", err)
	}
	log.Printf("ScanEmailTask got task cnt:%d", len(taskList))
	return
}

// UpdateEmailTask 更新邮件任务
func (d *Dao) UpdateEmailTask(ctx context.Context, taskInfo *model.EmailTask) (id int, err error) {
	r := d.db(ctx)

	log.Printf("UpdateEmailTask:%v", taskInfo)
	res := r.Debug().Updates(taskInfo)
	if res.Error != nil {
		log.Printf("UpdateEmailTask read db error(%v) id(%d)", res.Error, id)
	}
	return taskInfo.ID, res.Error
}

// QueryEmailTask 查询邮件任务详情。
func (d *Dao) QueryEmailTask(ctx context.Context, creator string) (taskList []*model.EmailTask, err error) {
	r := d.db(ctx)

	r = r.Debug().Where("email_creator = ?", creator)
	r = r.Order(clause.OrderByColumn{Column: clause.Column{Name: "triggle_time"}, Desc: false}).Limit(10)

	taskList = make([]*model.EmailTask, 0)
	if err = r.Debug().Find(&taskList).Error; err != nil {
		log.Printf("QueryEmailTask read db error(%v)", err)
	}
	log.Printf("QueryEmailTask got task cnt:%d", len(taskList))
	return
}
