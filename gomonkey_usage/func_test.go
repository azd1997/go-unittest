package gomonkey_usage

import (
	"errors"
	"fmt"
	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 使用gomonkey来进行函数打桩
func TestQuerySomeViaSrv(t *testing.T) {
	patches := gomonkey.ApplyFunc(
		CallOneSrv,		// 第一个参数为待破除的依赖函数名
		func(req string) (string, error) {	// 第二个参数为桩函数
			fmt.Println("call monkey srv")
			if req != "" {
				return req, nil
			}
			return "", errors.New("call failed")
		},
	)
	defer patches.Reset()	// 恢复

	// 简单测试下功能
	var v string
	var err error

	v, err = QuerySomeViaSrv("req")
	assert.Nil(t, err)
	assert.Equal(t, "req", v)

	v, err = QuerySomeViaSrv("")
	assert.NotNil(t, err)
	assert.Equal(t, "", v)
	fmt.Println(err)
}