package main

import (
	"fmt"
	"time"
)

/*
	切片内的元素往一个新的切片内转移的过程，效率最高的方法
	1. 普通的遍历append
	2. 已知容量的情况下，对新的切片进行带有容量值的初始化，然后进行append
	3. 已知容量的情况下，对新的切片进行带有长度值的初始化，然后进行下标元素的赋值替换
*/
// 初始化一个原始切片
func initSlice1(endIndex int) []int64 {
	s1 := []int64{}
	for i := 0; i < endIndex; i++ {
		s1 = append(s1, int64(i))
	}
	return s1
}

func main() {
	// 初始化一个带有10000000容量的切片
	s1 := initSlice1(10000000)
	// 进行三次取样，查看时间
	for i := 0; i < 3; i++ {
		SliceQuick3(s1)
	}
}

// SliceQuick1
// 普通的遍历append
//time now = 2022-07-25 10:56:56.0984786 +0800 CST m=+0.775080901
//time now = 2022-07-25 10:56:56.2042404 +0800 CST m=+0.880842701
//time now = 2022-07-25 10:56:56.3089455 +0800 CST m=+0.985547801
//time now = 2022-07-25 10:56:56.3338862 +0800 CST m=+1.010488501
//time now = 2022-07-25 10:56:56.3474395 +0800 CST m=+1.024041801
//time now = 2022-07-25 10:56:56.4109107 +0800 CST m=+1.087513001
func SliceQuick1(from []int64) {
	to := make([]int64, 0)
	fmt.Printf("time now = %v\n", time.Now())
	for _, v := range from {
		to = append(to, v)
	}
	fmt.Printf("time now = %v\n", time.Now())
}

// SliceQuick2
// 已知容量的情况下，对新的切片进行带有容量值的初始化，然后进行append
//time now = 2022-07-25 11:04:24.1822221 +0800 CST m=+0.985531301
//time now = 2022-07-25 11:04:25.9223641 +0800 CST m=+2.725673301
//time now = 2022-07-25 11:04:26.2852423 +0800 CST m=+3.088551501
//time now = 2022-07-25 11:04:26.4270343 +0800 CST m=+3.230343501
//time now = 2022-07-25 11:04:26.580933 +0800 CST m=+3.384242201
//time now = 2022-07-25 11:04:26.7329694 +0800 CST m=+3.536278601
func SliceQuick2(from []int64) []int64 {
	to := make([]int64, 0, len(from))
	fmt.Printf("time now = %v\n", time.Now())
	for _, v := range from {
		to = append(to, v)
	}
	fmt.Printf("time now = %v\n", time.Now())
	return to
}

// SliceQuick3
// 已知容量的情况下，对新的切片进行带有长度值的初始化，然后进行下标元素的赋值替换
//time now = 2022-07-25 11:06:19.6404986 +0800 CST m=+2.596835801
//time now = 2022-07-25 11:06:20.2189896 +0800 CST m=+3.175326801
//time now = 2022-07-25 11:06:20.9655849 +0800 CST m=+3.921922101
//time now = 2022-07-25 11:06:21.1607221 +0800 CST m=+4.117059301
//time now = 2022-07-25 11:06:21.2509867 +0800 CST m=+4.207323901
//time now = 2022-07-25 11:06:21.3638859 +0800 CST m=+4.320223101
func SliceQuick3(from []int64) []int64 {
	to := make([]int64, len(from))
	fmt.Printf("time now = %v\n", time.Now())
	for k, v := range from {
		to[k] = v
	}
	fmt.Printf("time now = %v\n", time.Now())
	return to
}
