package main

import (
	"fmt"
)

func join(a,b []int) []int{
	var finalArr  []int
	idxa,idxb := 0,0

	lena,lenb := len(a),len(b)

	for i:=0;i<(lena+lenb);i++{
		if idxb < lenb && idxa < lena{
			if a[idxa] >= b[idxb]{
				finalArr = append(finalArr,a[idxa])
				idxa++
			}else{
				finalArr = append(finalArr,b[idxb])
				idxb++
			}
		}else if idxb == lenb{
			finalArr = append(finalArr,a[idxa])
			idxa++
		}else if idxa == lena{
			finalArr = append(finalArr,b[idxb])
			idxb++
		}
	}
	return finalArr
}


func startSort(arr[]int) []int{
	if len(arr)==1{
		return arr
	}

	center := len(arr)/2

	arrA := startSort(arr[:center])
	arrB := startSort(arr[center:])

	return join(arrA,arrB)
}

func main(){
	var a,b,minDiff = 0,0,int(^uint(0)>>1)
	arr := startSort([]int{15, 30, 8, 1, 50, 21, 48})
	for i:=0;i<len(arr)-1;i++{
		if arr[i]-arr[i+1] < minDiff{
			a = arr[i]
			b = arr[i+1]
			minDiff = arr[i]-arr[i+1]
		}
	}
	fmt.Printf("Minimum is %d -%d = %d",a,b,minDiff)
}