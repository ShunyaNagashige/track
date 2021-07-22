package main

import "fmt"

func main() {
	var n, q int
	fmt.Scanf("%d %d", &n, &q)

	cards := make([]int, n)
	for i := 0; i < n; i++ {
		cards[i] = i + 1
	}

	var first, pos int
	for i := 0; i < n; i++ {
		first = cards[0]
		cards = cards[1:]
		pos = i
		// fmt.Println(i)
		// NGコード 代入順序が保証されないからダメでした
		// slice , slice[pos] = append(slice[:pos+1], slice[pos:]...), 追加する要素
		// fmt.Println(cards)
		if len(cards) == pos {
			cards=append(cards, first)
		} else {
			cards = append(cards[:pos+1], cards[pos:]...)
			cards[pos] = first
		}
		// fmt.Println(cards)
	}

	indexmem:=make([]int,n)
	for i:=0;i<n;i++{
		indexmem[cards[i]-1]=i+1
	}

	// fmt.Println(indexmem)

	t := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scanf("%d", &t[i])
	}
	for i := 0; i < q; i++ {
		fmt.Println(indexmem[t[i]-1])
	}
}
