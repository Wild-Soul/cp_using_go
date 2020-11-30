package main
	
import (
	"fmt"
	"sort"
	)
	
func root(x int, parent[]int) int {
	for ; x!=parent[x]; {
		parent[x] = parent[parent[x]];
		x = parent[x];
	}
	return x;
}
	
func join(u,v int, parent, size []int ) {
	var pu,pv int;
	pu = root(u, parent);
	pv = root(v, parent);
	
	if pu!=pv {
		if size[pu] > size[pv] {
			size[pu] += size[pv];
			parent[pv] = pu;
		} else {
			size[pv] += size[pu];
			parent[pu] = pv;
		}
	}
}
	
	
func initialize(parent, size []int) {
	n := len(parent);
	for i:=0; i<n; i++ {
		parent[i] = i;
		size[i] = 1;
	}
}
	
func main() {
	var n,m,u,v int;
	fmt.Scanf("%d %d\n", &n, &m);
	parent := make([]int, n+1);
	size := make([]int, n+1);
	initialize(parent,size);
	
	for i:=0; i<m; i++ {
		fmt.Scanf("%d %d\n", &u, &v);
		// fmt.Println(parent[u], size[u]);
		join(u,v,parent,size);
		
		set := make(map[int]bool);
		var ans []int;
		
		for i:=1; i<=n; i++ {
			x := i;
			for ; x!=parent[x]; {
				x = parent[x];
			}
			
			_, ok := set[x];
			if !ok {
				ans = append(ans,size[x]);
				set[x] = true;
			}
		}
		
		sort.Ints(ans);
		for i:=0; i<len(ans); i++ {
			fmt.Printf("%d ",ans[i]);
		}
		fmt.Println();
	}
}