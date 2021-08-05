package goutil

// here, we import the graph we defined in the previous section as the `graph` package
func DFS(g *Graph, startVertex *Vertex, visitCb func(string)) {
	// we maintain a map of visited nodes to prevent visiting the same
	// node more than once
	visited := map[string]bool{}

	if startVertex == nil {
		return
	}
	visited[startVertex.Key] = true
	visitCb(startVertex.Key)

	// for each of the adjacent vertices, call the function recursively
	// if it hasn't yet been visited
	for _, v := range startVertex.Vertices {
		if visited[v.Key] {
			continue
		}
		DFS(g, v, visitCb)
	}
}

// There is a better way to do this. TODO: Refactor
func FindPath(g *Graph, startVertex *Vertex, endVertex *Vertex, currentPath []string, path []string, visited map[string]bool) []string {

	if visited[startVertex.Key] {
		return path
	}

	visited[startVertex.Key] = true
	currentPath = append(currentPath, startVertex.Key)

	if startVertex.Key == endVertex.Key {

		path = append(path, currentPath...)

		visited[startVertex.Key] = false

		if len(currentPath) > 0 {
			currentPath = currentPath[:len(currentPath)-1]
		}

		return path
	}

	for _, next := range startVertex.Vertices {

		path = FindPath(g, next, endVertex, currentPath, path, visited)
	}

	if len(currentPath) > 0 {
		currentPath = currentPath[:len(currentPath)-1]
	}

	visited[startVertex.Key] = false

	return path
}
