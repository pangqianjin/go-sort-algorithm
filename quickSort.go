package main

import "fmt"
// 第一种写法，需要三个参数
func quickSort1(arr []int, low, high int){
	var i, j = low, high
	if i<j{
		pivot := arr[i]// 选取arr[i]作为基准数
		for i<j{// 或者i!=j
			for ; i<j && arr[j]>pivot; j--{}
			arr[i] = arr[j]// 把右侧小于基准数的元素放到左侧

			for ; i<j && arr[i]<pivot; i++ {}
			arr[j] = arr[i]// 把左侧大于基准数的元素放到右侧
		}
		arr[i] = pivot// 基准数还原为arr[i]
		quickSort1(arr, low, i-1)
		quickSort1(arr, i+1, high)
	}
}

// 第二种写法，只需要一个参数，但是需要有返回值
func quickSort2(arr []int)[]int{
	if len(arr)==0{
		return arr
	}

	result := make([]int, 0)
	pivot := arr[len(arr)-1]// 取基准数
	arr = arr[:len(arr)-1]// 弹出最后一个元素

	var left, right  []int
	for i:=0; i<len(arr); i++{
		if arr[i]<pivot{
			left = append(left, arr[i])// 小于基准数的放在左侧
		}else{
			right = append(right, arr[i])// 大于基准数的放在右侧
		}
	}

	result = append(result, quickSort2(left)...)// 放入左侧排序结果
	result = append(result, pivot)// 放入基准数
	result = append(result, quickSort2(right)...)// 放入右侧排序结果

	return result
}

func main() {
	var arr1 = []int{2, 44, 1, 0, -22, 56, -78}
	var arr2 = []int{2, 44, 1, 0, -22, 56, -78}
	quickSort1(arr1, 0, len(arr1)-1)
	arr2 = quickSort2(arr2)
	fmt.Println(arr1)// [-78 -22 0 1 2 44 56]
	fmt.Println(arr2)// [-78 -22 0 1 2 44 56]
}
