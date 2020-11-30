package main;
import (
	"fmt";
	"strconv";
	)
	
func check(s string) bool {
	n , _ := strconv.Atoi(s);
	if n%2==0 {
		return true;
	} else {
		return false;
	}
}
func main() {
	var s string;
	fmt.Scanf("%s",&s);
	//fmt.Println(s);
	
	n:=len(s);
	pre := make([]int, n);
	
	for i:=0; i<n; i++ {
		if(check(string(s[i]))) {
			pre[i] = 1;
		}
	}
	
	for i:=n-2; i>=0; i-- {
		pre[i] = pre[i] + pre[i+1];
	}
	for i:=0; i<n; i++ {
		fmt.Printf("%d ",pre[i]);
	}
}