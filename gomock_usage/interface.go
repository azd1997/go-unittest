package gomock_usage

import "fmt"

// 依赖接口
type DBStore interface {
	Put(k string, v string) error	// 存数据
	Get(k string) (string, error)	// 取数据
}

func NewDBStore() DBStore {
	return &redisStore{}
}


// 默认使用的redis存储。我们假设开发环境无法连通测试网络中的redis，或者这个redis存储的类还没实现
type redisStore struct {}

func (s *redisStore) Put(k, v string) error {
	fmt.Printf("redisStore: put (%s : %s)\n", k, v)
	return nil
}

func (s *redisStore) Get(k string) (string, error) {
	fmt.Printf("redisStore: get (%s)\n", k)
	return k, nil	// 这里直接将k作为v返回
}

