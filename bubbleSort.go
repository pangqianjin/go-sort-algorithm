package main

import "fmt"

func BubbleSort(arr []int) {
	for i:=len(arr); i>0; i--{
		for j:=0; j<i-1; j++{
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main()  {
	var arr = []int{2, 44, 1, 0, -22, 56, -78}
	BubbleSort(arr)
	fmt.Println(arr)// [-78 -22 0 1 2 44 56]
}