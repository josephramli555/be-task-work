package main

import (
	"fmt"
)

func zeroSum(data []int){
	ans := []string{}
	maps := make(map[int]int)
	for i:=0;i<len(data)-1;i++{
		for j:=i+1;j<len(data);j++{
			x :=(data[i]+data[j])
			if _,exist := maps[-1*x];exist{
				answer:=fmt.Sprintf("%d %d %d",-1*x,data[i],data[j])
				ans = append(ans,answer)
			}else{
				maps[x] = x
			}
		}
	}
	for _,a := range ans{
		fmt.Println(a)
	}
}

func main(){
	zeroSum([]int{-1,0,1,-4,3})
}