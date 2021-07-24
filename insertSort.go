package main

import "fmt"
/* 从当前元素的前一个元素开始，往前找一个合适的位置，插进去
	这个位置在第一个 比当前元素小 的元素之后
	期间，这两个元素之间的元素都要向后移动一个位置，来把空隙让出来
*/
func insertSort(arr []int) {
	var preIndex, current int
	for i:=1; i<len(arr); i++{
		current = arr[i]
		for preIndex = i-1; preIndex>=0 && arr[preIndex]>current; preIndex--{
			arr[preIndex+1] = arr[preIndex]// 两个元素之间的元素向后移动
		}

		arr[preIndex+1] = current// 放在第一个比当前元素小的元素之后
	}
}

func main()  {
	var arr = []int{2, 44, 1, 0, -22, 56, -78}
	insertSort(arr)
	fmt.Println(arr)// [-78 -22 0 1 2 44 56]
}
