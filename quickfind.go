package main

// Somehow create an array and leave the state
var root []int

func createSet(size int){
	for i := 0; i < size; i++ {
		root = append(root, i)
	}
}

func find(x int) int {
	return root[x]
}

func union(x int, y int) []int{
	rootX := find(x)
	rootY := find(y)
	if rootX != rootY{
		for i := 0; i < len(root); i++ {
			if root[i] == rootY{
				root[i] = rootX
			}
		}
	}
	return root
}

func connected(x int, y int) bool {
	return find(x) == find(y)
}