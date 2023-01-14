package main

// iteratively finds all unique paths
func (g *Graph) dfs(start, end *Room, visited []bool, path []*Room, paths *[][]*Room) {
	visited[start.id] = true
	path = append(path, start)

	if start == end {
		pathCopy := make([]*Room, len(path))
		copy(pathCopy, path)
		*paths = append(*paths, pathCopy)
	} else {
		for _, neighbor := range start.neighbors {
			if !visited[neighbor.id] {
				g.dfs(neighbor, end, visited, path, paths)
			}
		}
	}
	path = path[:len(path)-1]
	visited[start.id] = false
}

func (g *Graph) findPaths(start, end *Room) [][]*Room {
	visited := make([]bool, len(g.rooms))
	path := []*Room{}
	paths := [][]*Room{}

	g.dfs(start, end, visited, path, &paths)
	return paths
}
