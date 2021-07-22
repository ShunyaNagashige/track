package main

import "fmt"
var n,m int
var a []int
var b []int

func main(){
	fmt.Scanf("%d %d",&n,&m)
	
	a=make([]int,m)
	for i:=0;i<m;i++{
		fmt.Scanf("%d",&a[i])
	}

	b=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanf("%d",&b[i])
	}

	start:=make([]int,n)
	for i:=0;i<n;i++{
		start[i]=i
	}

	indexmem:=make([]int,m+1)
	ans1:=swap(1,0,indexmem,start)
	ans2:=swap(2,0,indexmem,start)
}

func swap(indexa,flag int,indexmem,amida []int)[]int{
	indexa--
	for i:=0;i<len(amida);i++{
		if amida[i]!=b[i]{
			break
		}
		return indexmem
	}

	if flag==1{
		indexmem=append(indexmem,indexa)
		tmp:=amida[indexa]
		amida[indexa]=amida[indexa+1]
		amida[indexa+1]=tmp
		return swap(a[indexa+1],flag,indexmem,amida)
	}
	
	for i:=0;i<n;i++{
		if indexa==i{
			tmp:=amida[indexa]
			amida[indexa]=amida[indexa+1]
			amida[indexa+1]=tmp
			return swap(a[indexa+1],flag,indexmem,amida)
		}
		flag=1
		i--
		tmp:=amida[i]
		amida[indexa]=amida[indexa+1]
		amida[indexa+1]=tmp
		return swap(a[indexa+1],flag,indexmem,amida)
	}

	return nil
}