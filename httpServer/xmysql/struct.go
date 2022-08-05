package xmysql

import (
	"database/sql"
)

type XMysql struct {
	db *sql.DB
}

var xMysql = XMysql{}
