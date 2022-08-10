package xphinx

import (
	"fmt"
	"github.com/myskull/common/xconfig"
	"github.com/myskull/common/xmysql"
	"testing"
)

func init() {
	xconfig.New("aa.conf")
	err := xmysql.Connect()
	fmt.Println(err)
}

func TestNew(t *testing.T) {
	//os.Setenv("XPHINX_COMMAND", "migrate")
	//os.Setenv("XPHINX_COMMAND", "rollback")
	err := New("v1.0.0").Migrate(func(builder *xmysql.XBuilder) error {
		return nil
	}).Rollback(func(builder *xmysql.XBuilder) error {
		return nil
	}).Do()
	fmt.Println(err)

	// create 创建一个版本
	// migrate 合并版本
	// rollback 回滚版本
	// init 初始化
}
