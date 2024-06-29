package allancestorsofanodeinadirectedacyclicgraph

// https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/description/

func getAllAncestors(result [][]int, index int, mapEdgeChildren map[int][]int) {
	seen := map[int]int{}

	children := mapEdgeChildren[index]
	counter := 0
	for len(children) > counter {
		if _, ok := seen[children[counter]]; !ok {
			seen[children[counter]] = 1
			result[children[counter]] = append(result[children[counter]], index)
			if len(mapEdgeChildren[children[counter]]) > 0 {
				children = append(children, mapEdgeChildren[children[counter]]...)
			}
		}
		counter++
	}
}

func getAncestors(n int, edges [][]int) [][]int {
	mapEdgeChildren := make(map[int][]int)
	result := make([][]int, n)

	for _, val := range edges {
		mapEdgeChildren[val[0]] = append(mapEdgeChildren[val[0]], val[1])
	}

	for index := 0; index < n; index++ {
		getAllAncestors(result, index, mapEdgeChildren)
	}

	return result
}
