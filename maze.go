package main

import (
	"os"
	"fmt"
)

func main() {
	maze := readMaze("maze.in") //读maze.in文件

	//验证读取是否正确
	for _,row := range maze{
		for _,val := range row{
			fmt.Printf("%3d ",val)
		}
		fmt.Println()
	}

	fmt.Println("--------")

	steps :=walk(maze,point{0,0},point{len(maze)-1,len(maze[0])-1})

	//打印走的路径
	for _,row := range steps{
		for _,val := range row{
			fmt.Printf("%3d ",val)
		}
		fmt.Println()
	}
}

type point struct {
	i,j int
}
//走的路径，指下一个方向
var dirs = []point{
	{-1,0},{0,-1},{1,0},{0,1}}

func (p point)add(r point) point {
	return point{p.i+r.i,p.j+r.j}
}

//获取点point在grid位置的值
func (p point)at(grid [][]int) (int,bool){
	if p.i < 0 || p.i >= len(grid){
		return 0,false
	}
	if p.j < 0 || p.j >=len(grid[p.i]){
		return 0,false
	}
	return grid[p.i][p.j],true
}

func walk(maze [][]int,start,end point) [][]int{
	steps := make([][]int,len(maze))
	for i := range steps{
		steps[i] = make([]int,len(maze[i]))
	}
	Q := []point{start}
	for len(Q) > 0{
		cur := Q[0]
		Q = Q[1:]//切片，去掉cur，依次循环下去cur都不是以前的值

		if cur == end{
			break
		}

		for _,dir := range dirs{
			next := cur.add(dir)
			//maze at next is 0
			//and steps at next is 0
			//and next != start
			val,ok := next.at(maze)
			if !ok || val == 1{
				continue
			}
			val,ok = next.at(steps)
			if !ok || val != 0{
				continue
			}
			if next == start{
				continue
			}
			curSteps,_ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q,next)
		}
	}
	return steps
}

func readMaze(filename string) [][]int {
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
    var row,col int
	fmt.Fscanf(file,"%d %d",&row,&col) //读取文件第一行两个数并跟据row、col的地址，将其赋值
	fmt.Println("row:",row,"col:",col) //row: 6 col: 5
	/*  测试fmt.Fscanf函数
		fmt.Fscanf(file,"%d",&row)
		fmt.Print("row:",row)//row:6
		fmt.Fscanf(file,"%d",&ces)
		fmt.Print("ces:",ces)//ces:0，换行会变为0
	*/
	maze := make([][]int,row)//声明一个row行的二维切片

	for  i:=range maze{
		maze[i] = make([]int,col)
		for j:= range maze[i]{
			_,err := fmt.Fscanf(file,"%d",&maze[i][j])
			//其中文件中换行符会导致error，所以采用以下方式多读取一行
			if err!=nil && err.Error() == "unexpected newline"{
				fmt.Fscanf(file, "%d")
			}
		}
	}
	return maze
}