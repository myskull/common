package xmysql

import (
	"database/sql"
	"fmt"
	"github.com/myskull/common/xjson"
	"reflect"
)

type XBuilder struct {
	sql           string
	where         string
	limit_n       int
	limit_m       int
	having        string // 跟group结合使用
	group         string
	table         string
	field         string
	order         string
	alias         string // 别名
	join          string // LEFT JOIN /  RIGHT JOIN
	joinTable     string // join的表名字
	joinOn        string // join的条件
	error         error
	duplicateData string // on duplicate key update
	db            *sql.DB
	tx            *sql.Tx // 事务指针
}

func NewBuilder(table ...string) *XBuilder {
	_table := ""
	if len(table) > 0 {
		_table = table[0]
	}
	builder := &XBuilder{
		table: _table,
		field: "*",
	}
	err := Connect()
	if err != nil {
		builder.error = err
	}
	builder.db = xMysql.db
	builder.tx = nil
	return builder
}

func (this *XBuilder) Alias(alias string) *XBuilder {
	this.alias = alias
	return this
}

func (this *XBuilder) Table(table string) *XBuilder {
	this.table = table
	return this
}

func (this *XBuilder) Order(order string) *XBuilder {
	this.order = order
	return this
}

func (this *XBuilder) Group(group string) *XBuilder {
	this.group = group
	return this
}

func (this *XBuilder) Where(where string, format ...interface{}) *XBuilder {
	_where := where
	if len(format) > 0 {
		_where = fmt.Sprintf(where, format...)
	}
	this.where = _where
	return this
}

// limit n,m  -> Limit(n,m)
// limit n	  -> Limit(n)
func (this *XBuilder) Limit(n int, m ...int) *XBuilder {
	if len(m) == 0 {
		this.limit_n = 0
		this.limit_m = n
	} else {
		this.limit_n = n
		this.limit_m = m[0]
	}
	return this
}

func (this *XBuilder) Field(field string) *XBuilder {
	this.field = field
	return this
}

func (this *XBuilder) Having(having string) *XBuilder {
	this.having = having
	return this
}

func (this *XBuilder) fetchSql() string {
	sql := fmt.Sprintf("SELECT %v from %v ", this.field, this.table)
	if this.alias != "" {
		sql += fmt.Sprintf(" as %v ", this.alias)
	}

	if this.join != "" {
		sql += fmt.Sprintf(" %v %v on %v ", this.join, this.joinTable, this.joinOn)
	}

	if this.where != "" {
		sql += fmt.Sprintf(" where %v ", this.where)
	}
	if this.group != "" {
		sql += fmt.Sprintf(" group by %v ", this.group)
	}
	if this.having != "" {
		sql += fmt.Sprintf(" having %v ", this.group)
	}
	if this.order != "" {
		sql += fmt.Sprintf(" order by %v ", this.order)
	}

	if this.limit_n > 0 || this.limit_m > 0 {
		sql += fmt.Sprintf(" limit %v,%v ", this.limit_n, this.limit_m)
	}
	return sql
}

func (this *XBuilder) LeftJoin(table string, on string) *XBuilder {
	this.join = " LEFT JOIN "
	this.joinTable = table
	this.joinOn = on
	return this
}

func (this *XBuilder) RightJoin(table string, on string) *XBuilder {
	this.join = " RIGHT JOIN "
	this.joinTable = table
	this.joinOn = on
	return this
}

func (this *XBuilder) query(_sql string) (xjson.A, error) {
	this.sql = _sql
	var rows *sql.Rows
	var err error
	if this.tx != nil {
		rows, err = this.tx.Query(this.sql)
	} else {
		rows, err = this.db.Query(this.sql)
	}
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	var values = make([]interface{}, len(columns))
	var scanRow = make([]interface{}, len(columns))
	for i := 0; i < len(values); i++ {
		scanRow[i] = &values[i]
	}
	result := xjson.A{}
	for rows.Next() {
		err = rows.Scan(scanRow...)
		if err != nil {
			return nil, err
		}
		row := xjson.M{}
		for i := 0; i < len(values); i++ {
			val := values[i]
			if val != nil {
				row[columns[i]] = string(val.([]byte))
			} else {
				row[columns[i]] = ""
			}
		}
		result = append(result, row)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("not found")
	}
	return result, nil
}

// 传入一个v结构体，通过结构体进行序列化
// @return 查询是否出错
func (this *XBuilder) Select() (xjson.A, error) {
	if this.error != nil {
		return nil, this.error
	}
	return this.query(this.fetchSql())
}

// 传入一个v结构体，通过结构体进行序列化
// @return 查询是否出错
func (this *XBuilder) Find() (xjson.M, error) {
	if this.error != nil {
		return nil, this.error
	}
	this.Limit(1)
	rows, err := this.query(this.fetchSql())
	if err != nil {
		return nil, err
	}
	return rows[0], nil
}

// 传入一个v结构体，通过结构体进行序列化
// @return 查询是否出错
func (this *XBuilder) FindObj(v interface{}) error {
	if this.error != nil {
		return this.error
	}
	result, err := this.Find()
	if err != nil {
		return err
	}
	Unmarshal(v, result)
	return nil
}

// 传入一个v结构体，通过结构体进行序列化
// @return 查询是否出错
func (this *XBuilder) SelctObj(v interface{}) error {
	if this.error != nil {
		return this.error
	}
	list, err := this.Select()
	if err != nil {
		return err
	}
	if len(list) > 0 {
		_value := reflect.ValueOf(v)
		_valueE := _value.Elem()
		_valueE = _valueE.Slice(0, _valueE.Cap())
		_element := _valueE.Type().Elem()
		i := 0
		for idx, row := range list {
			// 需要添加
			if _valueE.Len() == idx {
				elemp := reflect.New(_element)
				Unmarshal(elemp.Interface(), row)
				_valueE = reflect.Append(_valueE, elemp.Elem())
			}
			i++
		}
		_value.Elem().Set(_valueE.Slice(0, i))
	}
	return nil
}

func (this *XBuilder) Query(sql string) (xjson.A, error) {
	if this.error != nil {
		return nil, this.error
	}
	return this.query(sql)
}

func (this *XBuilder) Save(data xjson.M) (int64, error) {
	if this.error != nil {
		return 0, this.error
	}
	_sql := fmt.Sprintf("UPDATE %v set ", this.table)
	i := 0
	for key, val := range data {
		if i > 0 {
			_sql += ","
		}
		_sql += fmt.Sprintf(" `%v` = '%v' ", key, val)
		i++
	}
	if this.where != "" {
		_sql += fmt.Sprintf(" where %v ", this.where)
	}
	if this.limit_m > 0 || this.limit_n > 0 {
		_sql += fmt.Sprintf(" limit %v,%v ", this.limit_n, this.limit_m)
	}
	this.sql = _sql
	var result sql.Result
	var err error
	if this.tx != nil {
		result, err = this.tx.Exec(this.sql)
	} else {
		result, err = this.db.Exec(this.sql)
	}
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (this *XBuilder) Del(id ...uint32) (int64, error) {
	if this.error != nil {
		return 0, this.error
	}
	this.sql = fmt.Sprintf(" DELETE FROM %v ", this.table)
	this.sql += " where 1 = 1 "
	if this.where != "" {
		this.sql += fmt.Sprintf(" and %v ", this.where)
	}
	if len(id) > 0 {
		this.sql += fmt.Sprintf(" and id = %v ", id[0])
	}
	if this.order != "" {
		this.sql += fmt.Sprintf(" order by %v ", this.order)
	}
	if this.limit_m > 0 || this.limit_n > 0 {
		this.sql += fmt.Sprintf(" limit %v,%v ", this.limit_n, this.limit_m)
	}
	var result sql.Result
	var err error
	if this.tx != nil {
		result, err = this.tx.Exec(this.sql)
	} else {
		result, err = this.db.Exec(this.sql)
	}
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (this *XBuilder) Add(data xjson.M) (int64, error) {
	if this.error != nil {
		return 0, this.error
	}
	this.sql = "insert into  "
	this.sql += fmt.Sprintf(" %v ", this.table)
	field := ""
	column := ""
	for key, val := range data {
		if field != "" {
			field += ","
		}
		if column != "" {
			column += ","
		}
		column += fmt.Sprintf(" `%v`", key)
		field += fmt.Sprintf("  '%v' ", val)
	}
	this.sql += fmt.Sprintf(" (%v)value(%v)", column, field)
	if this.duplicateData != "" {
		this.sql += fmt.Sprintf(" on duplicate key update %v ", this.duplicateData)
	}
	var result sql.Result
	var err error
	if this.tx != nil {
		result, err = this.tx.Exec(this.sql)
	} else {
		result, err = this.db.Exec(this.sql)
	}
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (this *XBuilder) GetLastSql() string {
	return this.sql
}

func (this *XBuilder) OnDuplicateKey(data string) *XBuilder {
	this.duplicateData = data
	return this
}

func (this *XBuilder) ReplaceInfo(data xjson.M) (int64, error) {
	if this.error != nil {
		return 0, this.error
	}
	this.sql = "replace into  "
	this.sql += fmt.Sprintf(" %v ", this.table)
	field := ""
	column := ""
	for key, val := range data {
		if field != "" {
			field += ","
		}
		if column != "" {
			column += ","
		}
		column += fmt.Sprintf(" `%v`", key)
		field += fmt.Sprintf("  '%v' ", val)
	}
	this.sql += fmt.Sprintf(" (%v)value(%v)", column, field)
	var result sql.Result
	var err error
	if this.tx != nil {
		result, err = this.tx.Exec(this.sql)
	} else {
		result, err = this.db.Exec(this.sql)
	}
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// 事务启动
func (this *XBuilder) StartTrans() error {
	tx, err := this.db.Begin()
	this.tx = tx
	return err
}

func (this *XBuilder) Commit() error {
	return this.tx.Commit()
}

func (this *XBuilder) Rollback() error {
	return this.tx.Rollback()
}

func (this *XBuilder) Exce(_sql string, format ...interface{}) (sql.Result, error) {
	this.sql = fmt.Sprintf(_sql, format...)
	if this.tx != nil {
		return this.tx.Exec(this.sql)
	} else {
		return this.db.Exec(this.sql)
	}
}
