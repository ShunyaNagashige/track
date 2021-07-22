package main

import (
	"fmt"
	"sort"
	// "math"
)

func main(){
	var N int
	fmt.Scanf("%d",&N)

	ab:=make(map[string][]string)
	// for i:= range ab{
	// 	ab[i]=make([]bool,math.MaxInt32)
	// }
	
	var a,b string
	for i:=0;i<N;i++{
		fmt.Scanf("%s %s",&a,&b)
		ab[a]=append(ab[a], b)
	}

	// キーをソートする
    slice1 := make([]string, len(ab))
    index := 0
    for key := range ab {
        slice1[index] = key
        index++
    }
    sort.Strings(slice1)

	for _,key:=range slice1{
		
	}
}