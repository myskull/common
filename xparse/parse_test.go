package xparse

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	fmt.Println(^uint(0) >> 1)
	fmt.Println(Int32("1"))
	fmt.Println(Int32("1.1"))
	fmt.Println(Int32("0.1"))
	fmt.Println(Int32("122"))
	fmt.Println(Int32("129"))
	fmt.Println(Int32("111111"))
	fmt.Println(Int32("-111111"))
	fmt.Println(Int32("11111222223331"))
	fmt.Println(Int32("-11111222223331"))
	fmt.Println(Int32("-111111"))
	fmt.Println(Int32("122ee"))
}
