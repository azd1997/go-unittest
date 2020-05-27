package gomock_usage

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"unittest/gomock_usage/mock_gomock_usage"
)

// 对应于QueryFromDBV1
// 尽管我们已经mock好了两个DBStore的实现：mockDBSucc/mockDBFail，但是我们没办法用其替换默认的redisStore来测试
func TestQueryFromDBV1(t *testing.T) {

}

// 对应于QueryFromDBV2
// 现在我们可以用mock实现替换默认的redisStore来测试
func TestQueryFromDBV2(t *testing.T) {
	k := "key"
	var v string
	var err error

	// 测试成功访问数据库
	db := &mockDBSucc{}
	v, err = QueryFromDBV2(db, k)
	if err != nil || v != k {
		t.Error(err)
	}

	// 测试访问数据库失败
	db1 := &mockDBFail{}
	v, err = QueryFromDBV2(db1, k)
	if err == nil || v != "" {
		t.Error(errors.New("err should not be nil"))
	}
}

// 为结构体的方法去做mock测试时，可以将依赖接口嵌入到结构体中，
// 再在结构体初始化时将依赖注入。（可以是在结构体的new方法中或者是setDb之类的方法中）
// 这样也可以去替换生产环境的redis，使用mock结构测试
func TestQueryHandler_QueryFromDB(t *testing.T) {
	k := "key"
	var v string
	var err error

	// 测试成功访问数据库
	db := &mockDBSucc{}
	qh := NewQueryHandler(db)
	v, err = qh.QueryFromDB(k)
	if err != nil || v != k {
		t.Error(err)
	}

	// 测试访问数据库失败
	db1 := &mockDBFail{}
	qh = NewQueryHandler(db1)
	v, err = qh.QueryFromDB(k)
	if err == nil || v != "" {
		t.Error(errors.New("err should not be nil"))
	}
}

// 继续测试QueryHandler_QueryFromDB，但是采用gomock框架
// 1. 使用mockgen生成mock代码：mockgen -source interface.go -destination mock_gomock_usage/gomock_mock.go
// 2. 在本测试函数中调用生成的mock代码，设置mock结构体的预期行为
// 3. 编写测试逻辑
//
// 使用gomock静态设置返回值
func TestQueryHandler_QueryFromDBV2(t *testing.T) {
	// 传入t，生成gomock.Controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// 根据ctrl生成mock结构体
	mockDB := mock_gomock_usage.NewMockDBStore(ctrl)
	// 模拟正常访问数据库，只允许调用1次
	mockDB.EXPECT().Get("key").Return("key", nil).Times(1)
	// 模拟异常访问数据库，只允许调用1次
	mockDB.EXPECT().Get("key").Return("", errors.New("get op failed")).Times(1)

	k := "key"
	var v string
	var err error

	// 测试成功访问数据库
	qh := NewQueryHandler(mockDB)
	v, err = qh.QueryFromDB(k)
	if err != nil || v != k {
		t.Error(err)
	}

	// 测试访问数据库失败
	v, err = qh.QueryFromDB(k)
	if err == nil || v != "" {
		t.Error(errors.New("err should not be nil"))
	}
	fmt.Println(err)
}

// 继续测试QueryHandler_QueryFromDB，但是采用gomock框架
// 1. 使用mockgen生成mock代码：mockgen -source interface.go -destination mock_gomock_usage/gomock_mock.go
// 2. 在本测试函数中调用生成的mock代码，设置mock结构体的预期行为
// 3. 编写测试逻辑
//
// 使用gomock动态设置返回值
func TestQueryHandler_QueryFromDBV3(t *testing.T) {
	// 传入t，生成gomock.Controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// 根据ctrl生成mock结构体
	mockDB := mock_gomock_usage.NewMockDBStore(ctrl)
	// 设置模拟数据库的动态返回规则，允许调用任意次
	mockDB.EXPECT().Get(gomock.Any()).DoAndReturn(func(k string) (string, error) {
		// 伪造一个规则，当k不为""时正常返回，为""时返回错误
		if k == "" {
			return "", errors.New("get op failed")
		} else {
			return k, nil
		}
	}).AnyTimes()

	var v string
	var err error

	// 测试成功访问数据库
	qh := NewQueryHandler(mockDB)
	v, err = qh.QueryFromDB("getkey")
	if err != nil || v != "getkey" {
		t.Error(err)
	}

	// 测试访问数据库失败
	v, err = qh.QueryFromDB("")
	if err == nil || v != "" {
		t.Error(errors.New("err should not be nil"))
	}
	fmt.Println(err)
}

// 继续测试QueryHandler_QueryFromDB，但是采用gomock框架
// 1. 使用mockgen生成mock代码：mockgen -source interface.go -destination mock_gomock_usage/gomock_mock.go
// 2. 在本测试函数中调用生成的mock代码，设置mock结构体的预期行为
// 3. 编写测试逻辑
//
// 使用gomock检测调用顺序
func TestQueryHandler_QueryFromDBV4(t *testing.T) {
	// 传入t，生成gomock.Controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// 根据ctrl生成mock结构体
	mockDB := mock_gomock_usage.NewMockDBStore(ctrl)
	// 设置模拟数据库的预期行为：
	// 接下来只允许调用一次Get("suc")，并且之后只允许调用一次Get("fail")
	gomock.InOrder(
		mockDB.EXPECT().Get("suc").Return("suc", nil).Times(1),
		mockDB.EXPECT().Get("fail").Return("", errors.New("get op failed")).Times(1),
	)

	var v string
	var err error

	// 测试成功访问数据库
	qh := NewQueryHandler(mockDB)
	v, err = qh.QueryFromDB("suc")
	if err != nil || v != "suc" {
		t.Error(err)
	}

	// 测试访问数据库失败
	v, err = qh.QueryFromDB("fail")
	if err == nil || v != "" {
		t.Error(errors.New("err should not be nil"))
	}
	fmt.Println(err)
}

// 继续测试QueryHandler_QueryFromDB， 使用gomock + testify断言
func TestQueryHandler_QueryFromDBV5(t *testing.T) {
	// 传入t，生成gomock.Controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// 根据ctrl生成mock结构体
	mockDB := mock_gomock_usage.NewMockDBStore(ctrl)
	// 设置模拟数据库的预期行为：
	// 接下来只允许调用一次Get("suc")，并且之后只允许调用一次Get("fail")
	gomock.InOrder(
		mockDB.EXPECT().Get("suc").Return("suc", nil).Times(1),
		mockDB.EXPECT().Get("fail").Return("", errors.New("get op failed")).Times(1),
	)

	var v string
	var err error

	// 测试成功访问数据库
	qh := NewQueryHandler(mockDB)
	v, err = qh.QueryFromDB("suc")
	assert.Nil(t, err)
	assert.Equal(t, "suc", v)

	// 测试访问数据库失败
	v, err = qh.QueryFromDB("fail")
	assert.NotNil(t, err)
	assert.Equal(t, "", v)
	fmt.Println(err)
}