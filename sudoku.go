package main
import (
	"fmt"
	"time"
)

type Sudoku struct{
	grid [9][9]int
} 

func (s *Sudoku) fillGrid(){
	s.fillBox(0,0)
	s.fillBox(3,3)
	s.fillBox(6,6)
	s.fillOtherCells(0,3)
}

func (s *Sudoku) isSafe(r int, c int, num int) bool{
	for i := 0; i < 9; i++ {
		if s.grid[r][i] == num{
			return false
		}
		if s.grid[i][c] == num{
			return false
		}
	}
	x := r/3 * 3
	y := c/3 * 3
	for i := 0; i < 3; i++ {
		for j:= 0; j < 3; j++ {
			if s.grid[x+i][y+j] == num{
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) fillBox(r int, c int){
	var num int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for{
				num = time.Now().Nanosecond()%9+1
				if s.isSafe(r,c,num){
					s.grid[r+i][c+j] = num
					break
				}
			}
		}
	}
}

func (s *Sudoku) fillOtherCells(r int, c int) bool{
	if c >= 9{
		c = 0
		r += 1
	}

	if r < 3{
		if c < 3{
			c = 3			
		}
	} else if r < 6{
			if c == 3{
				c = 6
			}
	} else {
		if c == 6{
			c = 0
			r += 1
		}
	}

	if r >= 9 {
		return true
	}

	for i := 1; i <= 9; i++ {
		if s.isSafe(r,c,i) {
			s.grid[r][c] = i
			if s.fillOtherCells(r,c+1) {
				return true
			}
			s.grid[r][c] = 0
		}
	}

	return false
}

func (s *Sudoku) generateFinalGrid(){
	emptyCells := map[int]bool{}
	i := 0
	for i < 38 {
		index := time.Now().Nanosecond()%81
		if emptyCells[index] == false{
			emptyCells[index] = true
			r := index/9
			c := index%9
			s.grid[r][c] = 0
			i++
		} 
	}
}

func main(){
	st := time.Now()
	var s Sudoku
	s.fillGrid()
	for _,v := range s.grid {
		fmt.Println(v)		
	}
	fmt.Println(" ----User Grid---- ")
	s.generateFinalGrid()
	for _,v := range s.grid {
		fmt.Println(v)		
	}
	fmt.Println(time.Now().Sub(st))
}