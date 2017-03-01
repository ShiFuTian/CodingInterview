package leetcode

func islandPerimeter(grid [][]int) int {
	var m, n int
	if m = len(grid); m == 0 {
		return 0
	}
	if n = len(grid[0]); n == 0 {
		return 0
	}

	var perimeter int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				perimeter += 4
				if i > 0 && grid[i-1][j] == 1 {
					perimeter -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					perimeter -= 2
				}
			}
		}
	}
	return perimeter
}
