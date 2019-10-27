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

  for index := 1; index < 10; index ++{            //Iterate over all os.Args[1:]
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

func DrawTheGrid(grid [9][9]int) { // this function make awesome things ♥‿♥
  fmt.Println("\n♛ ■ ■ ■ ♛ ■ ■ ■ ♛ ■ ■ ■ ♛")
  for horiz := 0; horiz < 9; horiz++ {
    fmt.Print("■ ")
    for vert := 0; vert < 9; vert++ {
      if vert == 3 || vert == 6 {
        fmt.Print("■ ")
      }
      fmt.Printf("%d ", grid[horiz][vert])
      if vert == 8 {
        fmt.Print("■")
      }
    }
    if horiz == 2 || horiz == 5 || horiz == 8 {
      fmt.Println("\n♛ ■ ■ ■ ♛ ■ ■ ■ ♛ ■ ■ ■ ♛")
    } else {
      fmt.Println()
    }
  }
}


func parseInput(input string) [9][9]int { 
  grid := [9][9]int{}
  i := 0
  for horiz := 0; horiz < 9; horiz++ {
    for vert := 0; vert < 9; vert++ {
      j := Atoi(string(input[i]))
      i++
      grid[horiz][vert] = j
    }
  }
  return grid
}

func Atoi(s string) int {
  result := 0
  sign := 1
  for i, r := range s { //i - index; r - all string digits
    if i == 0 { // 0 - exit
      if r == '-' { 
        sign = -1 
        continue
      }
      if r == '+' {
        continue 
      }
    }
    if r < '0' || r > '9' { 
      return 0
    }
    for i := 0; i <= 9; i++ {
      if rune(i) == r-'0' {
        result *= 10
        result += i
        break
      }
    }
  }
  return result * sign
}
