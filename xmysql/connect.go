package xmysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/myskull/common/xconfig"
	"time"
)

func Connect() error {
	if xMysql.db != nil {
		return nil
	}
	database := xconfig.GetStr("mysql", "database", "")
	hostname := xconfig.GetStr("mysql", "hostname", "")
	username := xconfig.GetStr("mysql", "username", "")
	password := xconfig.GetStr("mysql", "password", "")
	max_open_conns := xconfig.GetInt("mysql", "max_open_conns", 10)
	max_idle_conns := xconfig.GetInt("mysql", "max_idle_conns", 10) //
	max_left_time := xconfig.GetInt("mysql", "max_left_time", 10)   // 分钟
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v)/%v", username, password, hostname, database))
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute * time.Duration(max_left_time)) // 每个链接闲置3分钟
	db.SetMaxIdleConns(max_idle_conns)
	db.SetMaxOpenConns(max_open_conns) // 空置的链接
	xMysql.db = db
	return ping()
}

func ping() error {
	return xMysql.db.Ping()
}
