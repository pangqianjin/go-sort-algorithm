package main

import "fmt"

// 希尔排序是改进的插入排序,相当于分组插入
func shallSort(arr []int)  {
	var fitIndex, current int
	for step:=len(arr)/2; step>0; step/=2{
		for i:=step; i<len(arr); i++{
			current = arr[i]
			// fitIndex-step是 与current同组的 其他元素的下标
			for fitIndex=i; fitIndex-step>=0 && arr[fitIndex-step]>current; fitIndex-=step{
				arr[fitIndex] = arr[fitIndex-step]
			}

			arr[fitIndex] = current
		}
	}
}

func main()  {
	var arr = []int{2, 44, 1, 0, -22, 56, -78}
	shallSort(arr)
	fmt.Println(arr)// [-78 -22 0 1 2 44 56]
}
