package xmysql

import (
	"fmt"
	"github.com/myskull/common/httpServer/xconfig"
	"github.com/myskull/common/httpServer/xjson"
	"testing"
)

func init() {
	xconfig.New("aa.conf")
	err := Connect()
	fmt.Println(err)
}
func TestConnect(t *testing.T) {

}

func TestNewBuilder(t *testing.T) {
	//obj := []string{}
	_, err := NewBuilder("bn_user").Where("uid=%d", 1).Save(xjson.M{
		"agent_id": 0,
	})
	fmt.Println(err)
	result, err := NewBuilder("bn_user").Where("uid=%d", 1).Find()
	fmt.Println(result)
	fmt.Println(err)
}

func TestXBuilder_Add(t *testing.T) {
	id, err := NewBuilder("bn_user").OnDuplicateKey("agent_id = agent_id + 10").Add(xjson.M{
		"agent_id": 2,
		"code":     "xxxx",
	})
	fmt.Println(id)
	fmt.Println(err)
	result, err := NewBuilder("bn_user").Where("code='xxxx'").Find()
	fmt.Println(result)
	fmt.Println(err)
}

func TestXBuilder_LeftJoin(t *testing.T) {
	fmt.Println(10 * 1024 * 1024)
	return
	result, err := NewBuilder("bn_user a").LeftJoin("bn_user_ex b", "a.uid = b.uid").Where("a.uid=%d", 1).Find()
	fmt.Println(result)
	fmt.Println(err)
}

type TestObj struct {
	ID uint32 `db:"id" json:"id"`
}

func TestConnect2(t *testing.T) {
	data := &TestObj{}
	origin_data := xjson.M{
		"id": 1,
	}
	Unmarshal(data, origin_data)
	fmt.Println(data)
}
