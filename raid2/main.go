package main

import (
	"fmt"
	"os"
	//"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	inputString := ""

	for index := 1; index < 10; index ++{				     //Iterate over all os.Args[1:]
		str := os.Args[index]                              //take one string
		if len(str) != 9{
			fmt.Println("Error")
			return
		} 
		for _, letter := range str{                                //Iterate over symbols in string
			if letter >= '1' && letter <= '9'{                     //digit check
				inputString = inputString + string(letter)
			}else if letter == '.'{                               //46 = '.'
				inputString = inputString + "0"
			}else {
				fmt.Println("Error")
				return
			}   
		}
	}

	grid := parseInput(inputString)

	//DrawTheGrid(grid)

	if backtrack(&grid) {
		DrawTheGrid(grid)
	} else {
		fmt.Println("Error")
	}

}

func backtrack(grid *[9][9]int) bool {
	if !StillZeros(grid) {                                           //if no zeros in the grid
		return true                                              //then go and draw it
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					grid[i][j] = candidate
					if IsGridValid(grid) {
						if backtrack(grid) {
							return true
						}
						grid[i][j] = 0
					} else {
						grid[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func StillZeros(grid *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func IsGridValid(grid *[9][9]int) bool {

	//checking 3x3
	for i := 0; i < 9; i=i+3{
		for j := 0; j < 9; j=j+3{
			if !Check3x3(i, j, grid){
				return false	
			}
			
		}
	}

	//checking rows
	slice := make([]int, 9)
	var iterator int
	for i2 := 0; i2 < 9; i2++{
		iterator = 0
		for j2 := 0; j2 < 9; j2++{
			slice[iterator] = grid[i2][j2] 
			iterator++
		}
		if !SliceValid(slice){
			return false
		}
	}

	//checking columns
	for i3 := 0; i3 < 9; i3++{
		iterator = 0
		for j3 := 0; j3 < 9; j3++{
			slice[iterator] = grid[j3][i3] 
			iterator++
		}
		if !SliceValid(slice){
			return false
		}
	}

	return true
}

func Check3x3(x int, y int, grid *[9][9]int) bool {
	slice := make([]int, 9)
	iterator := 0
	for i := x; i < x+3; i++{
		for j := y; j < y+3; j++{
			slice[iterator] = grid[i][j]
			iterator++
		}
	}	
	//fmt.Println(slice)
	if SliceValid(slice){
		//fmt.Println(slice)
		return true		
	}
	return false
}

func SliceValid(slice []int) bool {
	for i := 0; i < 8; i++{
		for j := i+1; j < 9; j++{
			if slice[i] == slice[j] && slice[i] != 0{
				return false	
			}
		}
	} 
	return true
}

func DrawTheGrid(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func parseInput(input string) [9][9]int {
	board := [9][9]int{}
	i := 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			i1 := Atoi(string(input[i]))
			i++
			board[row][col] = i1
		}
	}
	return board
}

func Atoi(s string) int {
	dec := 0
	signNumb := 0
	index := 0
	var signV rune
	for i, j := range s {
		ed := 0
		if j == '-' || j == '+' {
			index = i
			signNumb++
			if j == '-' {
				signV = '-'
			}
			if index > 0 {
				return 0
			}
		} else if j < '0' || j > '9' || signNumb > 1 {
			return 0
		}
		for i := '1'; i <= j; i++ {
			ed = ed + 1
		}
		dec = dec*10 + ed
	}
	if signV == '-' {
		return dec * -1
	} else {
		return dec
	}
}
