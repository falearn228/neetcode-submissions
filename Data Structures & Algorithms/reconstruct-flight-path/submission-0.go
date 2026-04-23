import "slices"

func findItinerary(tickets [][]string) []string {
	adj := make(map[string][]string, len(tickets))

	// create adj list
	for _, v := range tickets {
		adj[v[0]] = append(adj[v[0]], v[1])
	}

	if len(adj["JFK"]) == 0 {
		return []string{"JFK"}
	}

	// reorder for min lexim
	for key := range adj {
		slices.Sort(adj[key])
	}

	res := make([]string, 0, len(tickets))
	
	var dfs func(string)
	dfs = func(from string) {
		for len(adj[from]) > 0 {
			to := adj[from][0]
			adj[from] = adj[from][1:]
			dfs(to)
		}
		res = append(res, from)
	}
	dfs("JFK")
	slices.Reverse(res)
	return res
}

