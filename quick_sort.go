package main
import "fmt"

func partition(a []int, l, r int) int {
	part:=l-1;
	pivot := a[r];
	for i:=l; i<r; i++ {
		if a[i] <= pivot {
			part++;
			a[part], a[i] = a[i], a[part];
		}
	}
	part++;
	a[part], a[r] = a[r], a[part];
	
	return part;
}

func quick_sort(a []int, l,r int) {
	if l>=r {
		return;
	}
// 	fmt.Println(l,r);
	part := partition(a, l, r);
	quick_sort(a, l, part-1);
	quick_sort(a, part+1, r);
}

func main(){
	
	var t,n int;
	fmt.Scanf("%d\n",&t);
	
	for ;t>0; t-- {
		fmt.Scanf("%d\n", &n);
// 		fmt.Println(n);
		a := make([]int, n);
		for i:=0; i<n; i++ {
			fmt.Scanf("%d",&a[i]);
// 			fmt.Printf("%d ", a[i]);
		}
// 		fmt.Println();
		quick_sort(a,0,n-1);
		for i:=0; i<n; i++ {
			fmt.Printf("%d ",a[i]);
		}
		fmt.Println();
	}
	
}
