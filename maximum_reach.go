package main;
	
import "fmt"
	
func main() {
	var n int;
	fmt.Scanf("%d",&n);
	arr := make([]int, n);
	for i:=0; i<n; i++ {
		fmt.Scanf("%d", &arr[i]);
	}
	
	ans := make([]int, n);
	ans[n-1] = 0;
	for i:=n-2; i>=0; i-- {
		for j:=i+1; j<n; j++ {
			if arr[j] > arr[i] {
				ans[i] = ans[j] + 1;
				break;
			} else if ans[j] == 0 {
				break;
			}
		}
	}
	for i:=0; i<n; i++ {
		fmt.Printf("%d ", ans[i]);
	}
}