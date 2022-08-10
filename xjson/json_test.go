package xjson

import "testing"

func TestNew(t *testing.T) {
	data := []byte(`[1,2,[1,2],{"a":"x"}]`)
	New(data)
}
