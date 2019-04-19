package search

// 200
type DSU struct {
	Root []int
	Count int
}

func (this *DSU) find(x int) int  {

	if this.Root[x] == x {
		return x
	}
	nx := this.find(this.Root[x])
	this.Root[x] = nx
	return nx
}

func (this *DSU) union(x, y int)  {

	rx, ry := this.find(x), this.find(y)
	if rx != ry {
		this.Root[rx] = ry
		this.Count--
	}
}

func numIslands(grid [][]byte) int {

	if len(grid) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	n := row * col
	dsu := DSU{make([]int, n), 0}
	for i := 0; i < n; i++ {
		dsu.Root[i] = i
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col ; j++ {
			if grid[i][j] == '1' {
				dsu.Count++
			}
		}
	}
	// traverse the graph
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			for k := 0; k < 4; k++ {
				dir := dirs[k]
				x, y := i + dir[0], j + dir[1]
				if x < 0 || x >= row || y < 0 || y >= col {
					continue
				}
				if grid[i][j] == '0' || grid[x][y] == '0' {
					continue
				}
				dsu.union(i * col + j, x * col + y)
			}
		}
	}

	return dsu.Count
}