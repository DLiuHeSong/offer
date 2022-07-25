package main

import (
	"unsafe"
)

/*
	每个切片都指向一个底层数组
	每个切片都保存了当前切片的长度、底层数组可用容量
	使用len()计算切片长度时间复杂度为O(1)，不需要遍历切片
	使用cap()计算切片容量时间复杂度为O(1)，不需要遍历切片
	通过函数传递切片时，不会拷贝整个切片，因为切片本身只是个结构体而矣
	使用append()向切片追加元素时有可能触发扩容，扩容后将会生成新的切片
*/
type Slice struct {
	Cap int            // 容量
	Len int            // 长度
	Arr unsafe.Pointer // 指向数组的指针
}

func main() {
	// 1. 首先认识一下slice的结构
	// 2. 切片与arr的区别
	// 3. 切片的扩容原理

}

// 模拟切片的初始化过程
// 切片的模拟是模拟原理，并不是实际的过程
// 其中长度和容量可以不指定，则进行初始化有值状态下，取值的长度
func initSlice() Slice {
	// 具有5个元素的数组
	arr := [5]int{1, 2, 3, 4, 5}
	// slice的长度页数5个
	slice := Slice{Len: len(arr), Cap: len(arr), Arr: unsafe.Pointer(&arr)}
	return slice
}
