package gomock_usage

import (
	"errors"
	"fmt"
)

// 这里准备了两个DBStore的mock实现

// 正常查询
type mockDBSucc struct {
}

func (s *mockDBSucc) Put(k, v string) error {
	fmt.Printf("mockDBSucc: put (%s : %s)\n", k, v)
	return nil
}

func (s *mockDBSucc) Get(k string) (string, error) {
	fmt.Printf("mockDBSucc: get (%s)\n", k)
	return k, nil	// 这里直接将k作为v返回
}

// 查询失败
type mockDBFail struct {
}

func (s *mockDBFail) Put(k, v string) error {
	fmt.Printf("mockDBFail: put (%s : %s)\n", k, v)
	return nil
}

func (s *mockDBFail) Get(k string) (string, error) {
	fmt.Printf("mockDBFail: get (%s)\n", k)
	return "", errors.New("get op failed")	// 查询失败
}