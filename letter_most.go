package main
	
import (
	"fmt"
)
	
func main() {
	var n int
	var s string
	fmt.Scanf("%d\n",&n)
	fmt.Scanf("%s",&s)
	m:=make(map[rune]int)
	for _, v:= range s {
		m[v]++
	}
	ans := 1
	for _,v := range m {
		if v>ans {
			ans = v
		}
	}
	fmt.Println(ans)
}