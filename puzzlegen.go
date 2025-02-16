package main

import (
	"fmt"
	"math/rand"
	"time"
)

const gridSize = 9

// Generate an empty grid
func createEmptyGrid() [][]int {
	grid := make([][]int, gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}
	return grid
}

// Check if placing a number is valid
func isValid(grid [][]int, row, col, num int) bool {
	for x := 0; x < gridSize; x++ {
		if grid[row][x] == num || grid[x][col] == num {
			return false
		}
	}
	boxStartRow, boxStartCol := row/3*3, col/3*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[boxStartRow+i][boxStartCol+j] == num {
				return false
			}
		}
	}
	return true
}

// Solve the Sudoku grid using backtracking
func solve(grid [][]int) bool {
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if grid[row][col] == 0 {
				for num := 1; num <= gridSize; num++ {
					if isValid(grid, row, col, num) {
						grid[row][col] = num
						if solve(grid) {
							return true
						}
						grid[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

// Remove cells to create the puzzle
func removeCells(grid [][]int, clues int) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	totalCells := gridSize * gridSize
	toRemove := totalCells - clues
	for toRemove > 0 {
		row := random.Intn(gridSize)
		col := random.Intn(gridSize)
		if grid[row][col] != 0 {
			grid[row][col] = 0
			toRemove--
		}
	}
}

// Convert the grid to the desired format
func formatGrid(grid [][]int) []string {
	formatted := make([]string, gridSize)
	for i := 0; i < gridSize; i++ {
		row := ""
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 0 {
				row += "."
			} else {
				row += fmt.Sprintf("%d", grid[i][j])
			}
		}
		formatted[i] = row
	}
	return formatted
}

func main() {
	grid := createEmptyGrid()
	if solve(grid) {
		removeCells(grid, 30) // Adjust the number of clues as needed
		formattedGrid := formatGrid(grid)
		for _, row := range formattedGrid {
			fmt.Printf("\"%s\"\n", row)
		}
	} else {
		fmt.Println("Failed to generate a Sudoku puzzle.")
	}
}
