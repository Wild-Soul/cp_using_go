package main;
	
import (
	"fmt"
	"math"
)
	
func solve(i,j int, a []int, dp [][]int) int {
	if(i==j){
		return 1;
	} else if(i>j) {
		return 0;
	}
	if(dp[i][j] != -1){return dp[i][j];}
	temp:=a[i];
	
	ans := 100000000
	
	for k:=i; k<=j; k++ {
		if temp == a[k] {
			if i==k {
				ans = int(math.Min(float64(ans), float64( 1+solve(k+1, j,a,dp))));
			} else if k==j {
				ans = int(math.Min(float64(ans), float64(solve(i+1, k-1,a,dp))));
			} else {
				ans = int(math.Min(float64(ans), float64(solve(i+1,k-1,a,dp) + solve(k+1,j,a,dp))));
			}
		}
	}
	dp[i][j] = ans;
	return dp[i][j];
}
	
func main() {
	
	var n int;
	fmt.Scanf("%d",&n);
	
	a := make([]int,n);
	
	for i:=0; i<n; i++ {
		fmt.Scanf("%d",&a[i]);
	}
	fmt.Println();
	dp := make([][]int,n);
	for i:=0; i<n; i++ {
		dp[i] = make([]int,n,n);
		for j:=0; j<n; j++ {
			dp[i][j] = -1;
		}
	}
	fmt.Println(solve(0,n-1, a, dp));
	
	// for i:=0; i<n; i++ {
	//     for j:=0; j<n; j++ {
	//         fmt.Printf("%d ",dp[i][j]);
	//     }
	//     fmt.Println();
	// }
}