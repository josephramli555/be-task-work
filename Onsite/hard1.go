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
	N :=4
	a := [][]int{
		{1,3,4},
		{98,71,23},
		{52,31,63},
	}
	x := a[0]

	for _,val := range a[1:]{
		x = append(x,val...)
	}
	res := startSort(x)
	fmt.Printf("%v \n",res)
	fmt.Printf("%v",res[N-1])
}