package main

import (
	"fmt"
	"time"
)

type Foo func(s string) string

func Decorator(f Foo) Foo {
	return func(ss string) string {
		fmt.Println("decorator start")
		res := f(ss)
		fmt.Println("decorator end")
		return res
	}
}

func Hello(s string) string {
	fmt.Println(s)
	return fmt.Sprintf("%s", s)
}

func main() {
	a := Decorator(Hello)
	a("test")
	nums := []int{1, 2, 0, 3, 4}
	fmt.Println("before nums = ", nums)
	timeNow := time.Now()
	moveZeroes(nums)
	fmt.Println("cost ", time.Since(timeNow))
	fmt.Println("after nums = ", nums)
}

//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
// 1 2 0 3 4
// 1 2 3 0 4
//

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	//j := 0
	//for i := 0; i < len(nums); i++ {
	//	if nums[i] != 0 {
	//		nums[j] = nums[i]
	//		j++
	//	}
	//}
	//
	//for j < len(nums) {
	//	nums[j] = 0
	//	j++
	//}
}
