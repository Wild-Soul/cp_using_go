package main
import "fmt"

func main(){
	
	var t,n,m int;
	fmt.Scanf("%d\n", &t);
	for ;t>0; t-- {
		fmt.Scanf("%d %d\n",&n,&m);
		a := make([]int, n);
		b := make([]int, m);
		for i:=0; i<n; i++ {
			fmt.Scanf("%d",&a[i]);
		}
		for i:=0; i<m; i++ {
			fmt.Scanf("%d",&b[i]);
		}
		c := make([]int, n+m);
		i,j := 0,0;
		idx := 0;
		for ;i<n && j<m; {
			if a[i] < b[j] {
				c[idx] = a[i];
				idx++; i++;
			} else {
				c[idx] = b[j];
				idx++; j++;
			}
		}
		
		for ;i<n; {
			c[idx] = a[i];
			idx++; i++;
		}
		for ;j<m; {
			c[idx] = b[j];
			idx++; j++;
		}
		for i=0; i<(n+m); i++ {
			fmt.Printf("%d ",c[i]);
		}
		fmt.Println();
	}
	
}