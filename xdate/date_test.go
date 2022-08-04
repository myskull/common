package xdate

import (
	"fmt"
	"testing"
)

func TestToDate(t *testing.T) {
	//fmt.Println(ToTime("2022-07-01"))
	fmt.Println(ToTime("2022-07-01"))
	fmt.Println(ToTime("2022-07-01 00:00:00"))
	fmt.Println(ToTime("2022/07/01 00:00:00"))
	fmt.Println(ToDate(111, "Y-m-d"))
}
