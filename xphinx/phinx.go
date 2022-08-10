package xphinx

import (
	"github.com/myskull/common/xLog"
	"github.com/myskull/common/xjson"
	"github.com/myskull/common/xmysql"
	"time"
)

type XPhinx struct {
	version  string
	id       int
	builder  *xmysql.XBuilder
	migrate  func(builder *xmysql.XBuilder) error
	rollback func(builder *xmysql.XBuilder) error
}

var command = "migrate" // 默认是合并
var default_id = 0
var isInit = false

func SetCommand(_command string) {
	command = _command
}

/**
新建一条版本号:会通过数据库进行校验版本号是否存在
*/
func New(version string) *XPhinx {
	if !isInit {
		initTable()
		isInit = true
	}

	default_id += 1
	id := default_id
	builder := xmysql.NewBuilder()
	builder.StartTrans()
	xphinx := &XPhinx{
		id:      id,
		version: version,
		builder: builder,
		migrate: nil,
	}
	return xphinx
}

func (this *XPhinx) Migrate(migrate func(builder *xmysql.XBuilder) error) *XPhinx {
	this.migrate = migrate
	return this
}

func (this *XPhinx) Rollback(rollback func(builder *xmysql.XBuilder) error) *XPhinx {
	this.rollback = rollback
	return this
}

// 执行命令
func (this *XPhinx) Do() error {
	var err error
	if command == "migrate" {
		// 合并
		v := this.get()
		if v != nil {
			return nil
		}
		err = this.migrate(this.builder)
		if err == nil {
			this.add()
		}
	} else {
		// 回滚 总是回滚最后一条
		v := this.get()
		if v == nil {
			return nil
		}
		err = this.rollback(this.builder)
		if err == nil {
			// 删除版本数据
			err = this.del()
		}
	}
	if err != nil {
		this.builder.Rollback()
	} else {
		this.builder.Commit()
	}
	return err
}

type XVersion struct {
	Id      uint32 `db:"id"`
	Version string `db:"version"`
	Addtime int64  `db:"addtime"`
	Date    string `db:"date"`
	Name    string `db:"name"`
}

const (
	Table = "xphinx"
)

func (this *XPhinx) get() *XVersion {
	v := &XVersion{}
	err := this.builder.Table(Table).Where("version='%v'", this.version).FindObj(v)
	if err != nil {
		if err.Error() != "not found" {
			xLog.Error("读取xphinx版本失败:%v", err)
		}
		return nil
	}
	return v
}

func (this *XPhinx) del() error {
	_, err := this.builder.Table(Table).Where("version='%v'", this.version).Del()
	return err
}

func (this *XPhinx) add() error {
	_, err := this.builder.Table(Table).Where("version='%v'", this.version).Add(xjson.M{
		"version": this.version,
		"addtime": time.Now().Unix(),
	})
	return err
}

func initTable() error {
	_, err := xmysql.NewBuilder().Exce("create table if not exists `%v`(`id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',`version` varchar(100) NOT NULL DEFAULT '' COMMENT '版本号',`addtime` int unsigned not null DEFAULT 0 COMMENT '添加时间',PRIMARY KEY (`id`), unique KEY `version` (`version`))ENGINE=InnoDB COMMENT='xphinx表';", Table)
	if err != nil {
		xLog.Error("创建表失败:%v", err)
	}
	return err
}
