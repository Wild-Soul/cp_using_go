package main
import "fmt"
	
func preprocess( dp [][]float64, N int ) {
	
	for i:=1; i<N; i++ {
		for j:=1; j<N; j++ {
			if i==0 || j==0 {
				dp[i][j] = 1.0;
				continue;
			}
			if j>=2 {
				dp[i][j] = float64(i)/float64(i+j) + (float64(j)/float64(i+j))*(float64(j-1)/float64(i+j-1))*dp[i][j-2];
			} else {
				dp[i][j] = float64(i)/float64(i+j);
			}
		}
	}
}
	
func main() {
	
	var t,g,r int;
	const N int = 1004;
	dp := make([][]float64,N);
	for i:=0; i<N; i++ {
		dp[i] = make([]float64, N);
		for j:=0; j<N; j++ {
			dp[i][j] = 1.0;
		}
	}
	preprocess(dp,N);
	fmt.Scanf("%d",&t);
	
	for i:=0; i<t; i++ {
		fmt.Scanf("%d %d",&r,&g);
		fmt.Printf("%0.6f\n", dp[r][g]);
	}
}