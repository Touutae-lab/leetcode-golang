package main

const mod = int(1e9 + 7)
const max = int(1e6 + 1)

func countPaths(grid [][]int) int {
	cached := make([][]int, len(grid))
	for i := range grid {
		cached[i] = make([]int, len(grid[0]))
	}

	result := 0
	for i := range grid {
		for j := range grid[0] {
			result += countSequences(grid, cached, max, i, j)
		}
	}
	return result % mod
}

func countSequences(grid [][]int, cached [][]int, previousValue, i, j int) int {
	// If we are out of bounds or the current value is not greater than the previous value, return 0
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || previousValue <= grid[i][j] {
		return 0
	}

	// if cached[i][j] > 0, it means we have already computed the count for this cell
	if cached[i][j] > 0 {
		return cached[i][j]
	}

	cached[i][j] = 1 // count current path
	// Go right
	cached[i][j] += countSequences(grid, cached, grid[i][j], i-1, j) % mod
	// Go down
	cached[i][j] += countSequences(grid, cached, grid[i][j], i, j+1) % mod
	// Go left
	cached[i][j] += countSequences(grid, cached, grid[i][j], i+1, j) % mod
	// Go up
	cached[i][j] += countSequences(grid, cached, grid[i][j], i, j-1) % mod

	// Store the computed count in dp table
	return cached[i][j]
}
