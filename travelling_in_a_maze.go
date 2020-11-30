package main
	
import "fmt";
	
func initialize(a, d [][]int, ans [][]bool, n,m int) {
	for i:=0; i<n; i++ {
		a[i] = make([]int, m);
		d[i] = make([]int, m);
		ans[i] = make([]bool, m);
	}
}
	
func input( a[][] int, n,m int) {
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if j==m-1 {
				fmt.Scanf("%d\n",&a[i][j]);
			} else {
				fmt.Scanf("%d",&a[i][j]);
			}
			
		}
	}
}
	
func main() {
	var n,m,t, maxi int;
	fmt.Scanf("%d\n",&t);
	
	for test:=0; test < t; test++ {
	
		fmt.Scanf("%d %d %d\n",&n, &m, &maxi);
		a := make([][]int, n);
		d := make([][]int, n);
		ans := make([][]bool, n);
	
		initialize(a,d,ans, n,m);
		input(a,n,m);
		input(d,n,m);
		// for i:=0; i<n; i++ {
		// 	for j:=0; j<m; j++ {
		// 		fmt.Printf("%d ",a[i][j]);
		// 	}
		// 	fmt.Println();
		// }
		
		for i:=0; i<n; i++ {
			for j:=0; j<m; j++ {
				temp := false;
	
				if i > 0 {
					if ans[i-1][j] && (a[i][j] + (i+j) * d[i][j]) <= maxi {
						temp = temp || true;
					} else {
						temp = temp || false;
					}
				} else {
					temp = true;
				}
	
				if j > 0 {
					if ans[i][j-1] && (a[i][j] + (i+1) * d[i][j]) <= maxi {
						temp = temp || true;
					} else {
						temp = temp || false;
					}
				} else {
					temp = temp || true;
				}
				ans[i][j] = temp;
			}
		}
	
		if ans[n-1][m-1] {
			fmt.Println("YES");
		} else {
			fmt.Println("NO");
		}
	}
}