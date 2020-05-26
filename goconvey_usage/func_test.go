package goconvey_usage

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// V1版本的测试代码，使用testing框架
func TestFuncV1(t *testing.T) {
	str := "azdloveazdaawerrz"
	should := 7

	ret := maxLengthOfNonRepeatSubStr(str)
	if ret != should {
		t.Errorf("计算 %s 最长不重复子串错误：应该为 %d， 计算为 %d ", str, should, ret)
	}
}

// V2版本的测试代码，使用testing框架，表格驱动测试
func TestFuncV2(t *testing.T) {
	tests := []struct{
		s string
		l int
	} {
		//常规测试
		{"azdloveazdaawerrz", 7},
		{"helloworld", 5},
		{"happynewyear", 3},  //注意这里故意写错一个
		//边缘测试
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcabc", 3},
	}

	for _, tt := range tests {
		calcL := maxLengthOfNonRepeatSubStr(tt.s)
		if tt.l != calcL {
			t.Errorf("计算 %s 最长不重复子串错误：应该为 %d， 计算为 %d ", tt.s, tt.l, calcL)
		}
	}
}

// V3版本的测试代码，使用goconvey框架
func TestFuncV3(t *testing.T) {
	convey.Convey(
		"TestFuncV3 should return the max length of non-repeat substr", t, func() {
			str := "azdloveazdaawerrz"
			should := 7
			convey.So(maxLengthOfNonRepeatSubStr(str), convey.ShouldEqual, should)
		})
}

// V4版本的测试代码，使用goconvey框架，表格驱动测试
func TestFuncV4(t *testing.T) {
	// Convey可以任意嵌套，并且只需在最上层的Convey中传入t
	convey.Convey("TestFuncV4 should return the max length of non-repeat substr", t, func() {
		tests := []struct{
			s string
			l int
		} {
			//常规测试
			{"azdloveazdaawerrz", 7},
			{"helloworld", 5},
			{"happynewyear", 5},  //注意这里故意写错一个
			//边缘测试
			{"", 0},
			{"b", 1},
			{"bbbbbbbb", 1},
			{"abcabcabc", 3},
		}

		for _, tt := range tests {
			convey.So(maxLengthOfNonRepeatSubStr(tt.s), convey.ShouldEqual, tt.l)
		}
	})
}

// V5版本的测试代码，使用goconvey框架，convey嵌套
func TestFuncV5(t *testing.T) {
	// Convey可以任意嵌套，并且只需在最上层的Convey中传入t
	convey.Convey("TestFuncV4", t, func() {

		convey.Convey("一般情况", func() {
			str := "azdloveazdaawerrz"
			should := 7
			convey.So(maxLengthOfNonRepeatSubStr(str), convey.ShouldEqual, should)
		})

		convey.Convey("边界情况", func() {
			str := "helloworld"
			should := 5
			convey.So(maxLengthOfNonRepeatSubStr(str), convey.ShouldEqual, should)
		})
	})
}
