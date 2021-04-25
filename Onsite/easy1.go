package main

import (
	"fmt"
)

func reverseArrInt(data []int)[]int{
	x :=[]int{}
	for idx,_ := range data{
		x = append(x,data[len(data)-1-idx])
	}
	return x
}

func reverseArrStr(data []string)[]string{
	x :=[]string{}
	for idx,_ := range data{
		x = append(x,data[len(data)-1-idx])
	}
	return x
}


func main(){
	strings := []string{"a","b","x","z"}
	// numbers := []int{1,3,5,7,8,9}
	
	fmt.Printf("%v",reverseArrStr(strings))
	// fmt.Printf("%v",reverseArrInt(numbers))
	
}