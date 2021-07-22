package main

import "fmt"

var c []int
var n int
var count int

func main(){
	fmt.Scanf("%d",&n)

	c=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanf("%d",&c[i])
	}

	count=0
	for i:=0;i<n;i++{
		search(0,c[i],0,0,0)
	}
}

//num:並べた個数，value:並べようとしている商品の目立ちやすさ
//flag:手前の方が目立ちやすくなるようなことをしたかどうか，beforeValue:一つ前の商品の目立ちやすさ
func search(num,value,flag,max,beforeValue int){
	if num==n{
		count++
	}

	if beforeValue>value&&max>value{
		return
	}

	num++
	if beforeValue>value{
		if flag==1{
			return
		}else{
			flag=1
			search(num,c[num],flag,max,value)
		}
	}
	search(num,c[num],flag,max,value)
}