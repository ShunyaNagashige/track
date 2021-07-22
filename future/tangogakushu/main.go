package main

import "fmt"

func main() {
	m := make(map[string]bool)
	var s string
	fmt.Scanf("%s", &s)
	n:=len(s)
	subs := ""

	//bitsを１からにすることで，，空文字を除外する
	for bits := 1; bits < (1 << uint64(n)); bits++ {
		subs = ""

		// bitsの各要素についてのループ
		for i := 0; i < n; i++ {
			// bitsのi個目の要素の状態がonかどうかチェック
			if (bits>>uint64(i))&1 == 1 {
				subs += s[i : i+1]
			}
		}

		if _, ok := m[subs]; !ok {
			m[subs] = true
		}
	}

	fmt.Println(len(m))
}
