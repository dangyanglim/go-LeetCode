package main

import (
	//"fmt"
)

func MaxHeapify(arr []int) {
	length:=len(arr)
	end := length
	for i := length/2 - 1; i >= 0; i-- {
		dad := i
		son := dad*2 + 1
		if son < end {
			if son+1 < end && arr[son+1] < arr[son] {
				son++
			}
			if arr[dad] < arr[son] {
			} else {
				arr[son], arr[dad] = arr[dad], arr[son]
			}
		} else {
			return
		}
	}

}
//堆排序
func heapify_sort(arr []int) {
	length:=len(arr)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		MaxHeapify(arr[i:length])
	}
}

//快排
func quickSort(values []int) {
	length:=len(values)
	if length <= 1 {
		return
	}
	if length == 2 {
		if values[0] > values[1] {
			values[0], values[1] = values[1], values[0]
		}
		return
	}
	mid := values[0]
	head, tail := 1, length-1
	for head < tail+1 {
		if values[head] > mid {
			values[head], values[tail] = values[tail], values[head]
			tail--
		} else {
			head++
		}
	}
	if tail == len(values)-1 {
		values[0], values[tail] = values[tail], values[0]
		quickSort(values[:tail])
	} else {
		quickSort(values[:head])
		quickSort(values[head:])
	}
}

//冒泡
func bubbleSort(values []int){
	length:=len(values)
	for i:=0;i<length-1;i++{
		for j:=0;j<length-1-i;j++{
			if values[j]>values[j+1]{
				values[j],values[j+1]=values[j+1],values[j]
			}
		}
	}
}

func main() {

}
