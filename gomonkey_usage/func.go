package gomonkey_usage

import "fmt"

// gomonkey用于对函数/方法/变量进行打桩
// 这里举一个函数的例子

// 这是一个向下游服务发起调用并得到结果的函数
// 我们假设下游服务还没实现，或者无法在开发环境调用下游服务，需要破除掉这个依赖
func CallOneSrv(req string) (string, error) {
	fmt.Println("call real srv")
	return "", nil
}

// 待测试的函数
// 函数内需要调用CallOneSrv，但现在想要破除掉对CallOneSrv的依赖
// 使用stub模式的话，需要将CallOneSrv赋给一个函数变量，来实现这个函数的替换
// 使用gomonkey，可以更加友好的解决这个问题
func QuerySomeViaSrv(k string) (string, error) {
	// 其他代码略
	v, err := CallOneSrv(k)

	return v, err
}

