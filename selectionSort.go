package main

import "fmt"

// 选择排序是改进的冒泡排序，减少了交换元素的次数
// 每次记住arr[i+1]...arr[len(arr)-1]之间最小元素的索引，然后与arr[i]交换
func selectionSort(arr []int) {
	for i:=0; i<len(arr); i++{
		minIndex := i// 默认arr[i]为最小元素
		for j:=i+1; j<len(arr); j++{
			if arr[j] < arr[i]{
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main(){
	var arr = []int{2, 44, 1, 0, -22, 56, -78}
	selectionSort(arr)
	fmt.Println(arr)// [-78 -22 0 1 2 44 56]
}
