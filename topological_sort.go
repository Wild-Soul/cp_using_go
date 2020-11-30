package main
	
import (
	"fmt"
	"sort"
)
	
const N = 12
	
var adj = make([][]int,N)
var ans []int
	
func dfs(u int, vis []bool) {
	// fmt.Println("Visiting ", u)
	vis[u] = true
	for _, v := range adj[u] {
		if vis[v] {
			continue
		}
		// fmt.Println("Child ", v, " of ", u)
		dfs(v,vis)
	}
	ans = append(ans,u)
}
	
func main() {
	var n,m,u,v int
	
	fmt.Scanf("%d %d\n",&n,&m)
	for i:=0; i<m; i++ {
		fmt.Scanf("%d %d\n",&u,&v)
		adj[u] = append(adj[u], v)
	}
	
	for i:=1; i<=n; i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(adj[i])))
	}
	vis := make([]bool, n+1)
	for i:=n; i>0; i-- {
		if !vis[i] {
			dfs(i, vis)
		}
		
	}
	
	for i:=n-1; i>=0; i-- {
		fmt.Printf("%d ", ans[i])
	}
}