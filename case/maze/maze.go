package main

import (
	"os"
	"fmt"
)

func ReadFile(filename string) [][]int {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("打开文件错误")
		panic(err)
	}
	defer f.Close()
	//读取前两个数字作为行数和列数
	var row, col int
	fmt.Fscanf(f, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze { //遍历二维数组返回行标
		//创建列数
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(f,"%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1,0 }, {0, 1},
}
//走的方法
func (p point) add(dir point) point {
	return point{p.i + dir.i, p.j + dir.j}
}

//查看点所在格子的值
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >=len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}
func walk(maze [][]int, start, end point) [][]int{
	//定义走的步子新格子,全是0
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//创建一个可探索点的队列
	o := []point{start}
	for len(o) > 0 {
		cur := o[0]
		o = o[1:]
		if cur == end {
			break
		}

	//开始探索
	for _, dir := range dirs {
		next := cur.add(dir)
		//走道阻塞或者越界
		val, ok := next.at(maze)
		if  !ok || val == 1 {
			continue
		}
		//判断没有走过
		val, ok = next.at(steps)
		if  !ok || val != 0 {
			continue
		}
		//判断不是起点
		if next == start {
			continue
		}
			//1,将格子数填进去
			//2，将探索的点加进队列里
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps+1

			o = append(o, next)
	}
	}
	return  steps
}

func main() {
	maze := ReadFile("case/maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
	//steps:=walk(maze,point{0,0},point{len(maze)-1,len(maze[0])-1})
	//for _, row := range steps{
	//	for _, val := range row {
	//		fmt.Printf("%3d ", val)
	//	}
	//	fmt.Println()
	//}
}
