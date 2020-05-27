package gomock_usage

// 待测试的函数
// 我们需要模拟访问DB时连接失败和连接成功
func QueryFromDBV1(k string) (string, error) {
	// 可能有些数据输入处理、输出处理啥的，这里反正都略去

	db := NewDBStore()	// 获得DBStore的一个实例
	return db.Get(k)	// 调用接口实例的方法
}

// 对于函数，似乎只能将依赖作为参数传入，这样才能进行mock测试
func QueryFromDBV2(db DBStore, k string) (string, error) {
	// 可能有些数据输入处理、输出处理啥的，这里反正都略去

	return db.Get(k)	// 调用接口实例的方法
}

type QueryHandler struct {
	db DBStore		// 内嵌一个DBStore接口，用以依赖注入
}

// 我们在New方法中注入db依赖（而该New方法往往在main.go或者其他执行初始化的地方进行，这里省略）
func NewQueryHandler(db DBStore) *QueryHandler {
	return &QueryHandler{
		db: db,
	}
}

func (h *QueryHandler) QueryFromDB(k string) (string, error) {
	v, err := h.db.Get(k)
	return v, err
}
