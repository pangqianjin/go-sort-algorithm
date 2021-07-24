package main

import "fmt"
// 归并排序基于分治思想
func mergeSort(arr []int)[]int {
	if len(arr)==0 || len(arr)==1{
		return arr
	}

	middle := len(arr)/2
	left := arr[:middle]
	right := arr[middle:]

	return merge(mergeSort(left), mergeSort(right))
}

// 辅助函数，用于将两个有序数组合并，并保持有序
func merge(left, right []int)[]int{
	result := make([]int, 0)
	for len(left)>0 && len(right)>0{
		if left[0] < right[0]{
			result = append(result, left[0])
			left = left[1:]
		}else{
			result = append(result, right[0])
			right = right[1:]
		}
	}
	// 放入剩余元素
	result = append(result, left...)
	result = append(result, right...)

	return result
}

func main()  {
	var arr = []int{2, 44, 1, 0, -22, 56, -78}
	arr = mergeSort(arr)
	fmt.Println(arr)// [-78 -22 0 1 2 44 56]
}
