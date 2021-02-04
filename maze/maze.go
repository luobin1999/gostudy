package main

import "fmt"

/*  读文件有问题
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d%d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
		fmt.Println(maze[i])
	}
	return maze
}
*/
func readMaze(filename string) [6][5]int {
	maze := [6][5]int{
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	return maze
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [6][5]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func (p point) at1(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [6][5]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			// maze at next is 0
			// and steps at next is 0
			// and next != start
			if val, ok := next.at(maze); !ok || val == 1 {
				continue
			}
			if val, ok := next.at1(steps); !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			val, _ := cur.at1(steps)
			steps[next.i][next.j] = val + 1
			Q = append(Q, next)
		}
	}
	return steps
}

/**
根据steps获取最短路径
*/
func getRoads(steps [][]int) []point {
	roads := make([]point, 0)
	Q := make([]point, 0)
	end := point{len(steps) - 1, len(steps[0]) - 1}
	roads = append(roads, end)
	Q = append(Q, end)
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		val, _ := cur.at1(steps)
		for _, dir := range dirs {
			pre := cur.add(dir)
			preVal, _ := pre.at1(steps)
			if preVal == val-1 {
				roads = append(roads, pre)
				if preVal == 0 {
					break
				}
				Q = append(Q, pre)
			}
		}
	}
	return roads
}

func main() {
	maze := readMaze("maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	fmt.Println("================")
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
	fmt.Println("================")
	raods := getRoads(steps)
	fmt.Println(raods)
}

/*
//咱们先从main函数开始捋
func readMaze(filename string)[][]int{ //查找文件函数 作用是以数组形式返回文件内容
	file, err := os.Open(filename)   //打开文件返回file指针
	if err != nil {
		panic(err)
	}
	var row, col int //定义存放行列变量
	fmt.Fscanf(file,"%d %d", &row, &col)//将文件头两个数字 也就是行列存入变量 注意是放入地址
	maze := make([][]int,row)  //开始循环创建 注意此行代码只是创建行数并没有创建列数 [][]int只是代表创建二维个数
	for i := range maze{   //找到每一行
		maze[i] = make([]int, col) //为每一行创建列数
		for j := range maze[i]{   //获得每一行的列数
			fmt.Fscanf(file,"%d",&maze[i][j])//为每一行列元素赋值 0/1
		}
	}
	return maze //将二维数组返回
}
type point struct{   //坐标结构体 尽量不用xy命名
	i, j int
}
var dirs = [4]point{ //上左下右顺序坐标点
	{-1,0},{0,-1},{1,0},{0,1},
}
func (p point)Add(r point) point{  //返回下一步的四种可能
	return point{p.i+r.i,p.j+r.j}
}
func (p point)at(grid [][]int) (int,bool){ // 判断下一步能否通过 是否越界
	if p.i < 0||p.i >= len(grid){
		return 0, false
	}
	if p.j <0 || p.j >= len(grid[p.j]){
		return 0, false
	}
	return grid[p.i][p.j], true
}
func walk(maze [][]int, start , end point)[][]int{ //整个算法核心
	steps := make([][]int, len(maze))  //创建一个存放路线的二维数组
	for i := range steps{  //同样创建列
		steps[i] = make([] int, len(maze[i]))
	}
	Q := []point{start}   //定义一个变量 及赋值变量坐标 这是算法的几个关键点之一
	for len(Q) > 0{   // 这个循环是判断我们的起点到终点能不能走通 走不通会结束循环
		cur := Q[0]   //获得当前没有用过的坐标
		Q = Q[1:]     //将上一步使用过的坐标排除 go语言的切片是不是很好用
		if cur == end{    //当坐标从0.0一步一步移动到指定结束坐标 循环结束
			break
		}
		for _, dir := range dirs{  //数组在上方定义 实际就是坐标点的上左下右 每一个坐标点都有四个方位 每次进行比较
			next := cur.Add(dir) //next装的坐标为上左下右的坐标点 next的含义是当前坐标点的下一个坐标点
			//由于我们采用的是上左下右的判断方式 越是到最后的坐标 越是最节省时间的 所以关键点在于 如果上坐标可以移动
			//下坐标也可以移动 下坐标会覆盖上坐标 因为我们是求最短路线 关键点之一

			val ,ok := next.at(maze)   //函数作用是判断当前坐标的下一步是否会越界及碰墙
			if !ok || val == 1{//如果会越界及碰墙则结束此次循环 说明这个坐标位置不正确
				continue
			}
			val, ok = next.at(steps) // 判断坐标点的下一步是否在走过的路线上 因为我们每走一步都会在路线上标记步数 除了第一步
			if !ok || val != 0{
				continue
			}
			if  next == start{ //这就是判断第二步是否掉回头到第一步的作用 补充第一个判断的
				continue
			}
			cursteps, _ := cur.at(steps)  //将当前坐标的步数返回给变量
			steps[next.i][next.j] = cursteps+1 //确定下一步的坐标并且让步数加一

			Q = append(Q,next) //将下一步的位置传给变量  关键点之一 如果上左下右没有一个地方可以通过 变量的长度为零 循环结束
		}
	}
	return steps
}
func main() {
	maze := readMaze("maze/maze.in") //函数作用很简单 找到目标文件并将文件内容存到二维数组然后返回给maze
	//函数作用是将迷宫路线以二维数组方式返回 第一个参数是将迷宫传入 第二个参数 确定起始位置 第三个参数 确定结束位置
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
	fmt.Println("=========================")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	//打印路线
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}
*/
