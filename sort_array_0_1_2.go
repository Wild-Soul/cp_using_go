package main
import "fmt"

func main(){
	
	var t,n,m int;
	fmt.Scanf("%d\n", &t);
	zero,one,two := 0,0,0;
	
	for ;t>0; t-- {
		fmt.Scanf("%d\n",&n);
		zero,two,one = 0,0,0
		
		for i:=0; i<n; i++ {
		    fmt.Scanf("%d", &m);
		    if m==0 {
		        zero++;
		    } else if m==1 {
		        one++;
		    } else {
		    	two++;
		    }
		}
		
		for i:=0; i<n; i++ {
			if i<zero {
				fmt.Printf("%d ",0);
			} else if i<(zero+one) {
				fmt.Printf("%d ",1);
			} else {
				fmt.Printf("%d ",2);
			}
		}
		fmt.Println();
	}
	
}