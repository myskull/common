package xjson

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type XJson struct {
	jsonData []byte // json解析器
	isList   bool
	data     interface{}
}
type A []interface{}
type M map[string]interface{}

func New(data []byte) *XJson {
	var xj = &XJson{
		jsonData: data,
	}
	err := xj.parse()
	if err != nil {
		return nil
	}
	return xj
}

func (this *XJson) parse() error {
	if string(this.jsonData[0:1]) == "[" {
		fmt.Println(string(this.jsonData), "是列表")
		this.isList = true
		var list A
		err := json.Unmarshal(this.jsonData, &list)
		if err != nil {
			fmt.Printf("非json结构!%+v", string(this.jsonData))
			return err
		}
		var result = A{}
		for _, row := range list {
			_type := reflect.TypeOf(row)
			// 重新解析一下
			if _type.String() == "[]interface {}" {
				// 表示还需要上襦解析
				b, err := json.Marshal(row)
				if err != nil {
					return err
				}
				result = append(result, New(b))
			} else if _type.String() == "map[string]interface {}" {
				result = append(result, row)
			} else {
				result = append(result, row)
			}
		}
		this.data = result
		fmt.Println("列表数据", this.data)
	} else {
		fmt.Println(string(this.jsonData), "不是列表")
		this.isList = false
	}
	return nil
}

func (this *XJson) ForEach(_func func(key, value interface{})) {

}
