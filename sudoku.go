package main

import (
	"fmt"
	"os"
	"strconv"
)

const size = 9

type Grid [size][size]int

func main() {
	if len(os.Args) != size+1 {
		fmt.Println("Error")
		return
	}

	grid, valid := parseInput(os.Args[1:])
	if !valid || !isValidSudoku(grid) {
		fmt.Println("Error")
		return
	}

	if solveSudoku(&grid) {
		printGrid(grid)
	} else {
		fmt.Println("Error")
	}
}

func parseInput(args []string) (Grid, bool) {
	var grid Grid
	for i, row := range args {
		if len(row) != size {
			return grid, false
		}
		for j, char := range row {
			if char == '.' {
				grid[i][j] = 0
			} else {
				num, err := strconv.Atoi(string(char))
				if err != nil || num < 1 || num > 9 {
					return grid, false
				}
				grid[i][j] = num
			}
		}
	}
	return grid, true
}

func isValidSudoku(grid Grid) bool {
	rowCheck, colCheck, boxCheck := [size][size]bool{}, [size][size]bool{}, [size][size]bool{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			num := grid[i][j]
			if num == 0 {
				continue
			}
			boxIndex := (i/3)*3 + j/3
			if rowCheck[i][num-1] || colCheck[j][num-1] || boxCheck[boxIndex][num-1] {
				return false
			}
			rowCheck[i][num-1], colCheck[j][num-1], boxCheck[boxIndex][num-1] = true, true, true
		}
	}
	return true
}

func solveSudoku(grid *Grid) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isSafe(grid, i, j, num) {
						grid[i][j] = num
						if solveSudoku(grid) {
							return true
						}
						grid[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func isSafe(grid *Grid, row, col, num int) bool {
	for i := 0; i < size; i++ {
		if grid[row][i] == num || grid[i][col] == num || grid[row/3*3+i/3][col/3*3+i%3] == num {
			return false
		}
	}
	return true
}

func printGrid(grid Grid) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
}
