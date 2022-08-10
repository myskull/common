package xmysql

import (
	"github.com/myskull/common/xjson"
	"reflect"
)

func Unmarshal(v interface{}, data xjson.M) {
	_type := reflect.TypeOf(v)
	_value := reflect.ValueOf(v)
	_valueE := _value.Elem()
	for i := 0; i < _value.Elem().NumField(); i++ {
		field_name := _type.Elem().Field(i).Tag.Get("db")
		if field_name == "" {
			continue
		}
		switch _type.Elem().Field(i).Type.String() {
		case "uint32":
			_valueE.Field(i).SetUint(data.Uint64(field_name, 0))
		case "uint64":
			_valueE.Field(i).SetUint(data.Uint64(field_name, 0))
		case "int32":
			_valueE.Field(i).SetInt(data.Int64(field_name, 0))
		case "int64":
			_valueE.Field(i).SetInt(data.Int64(field_name, 0))
		case "string":
			_valueE.Field(i).SetString(data.Get(field_name, ""))
		case "bool":
			_valueE.Field(i).SetBool(data.Bool(field_name, false))
		case "float32":
			_valueE.Field(i).SetFloat(data.Float64(field_name, 0))
		case "float64":
			_valueE.Field(i).SetFloat(data.Float64(field_name, 0))
		}
	}
}
