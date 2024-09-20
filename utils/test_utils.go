package utils

import (
	"container/list"
	"testing"

)

//import "golang.org/x/tools/go/packages"

const C = 1
//const C int64 = 1

/*var V struct { X int }*/

var V struct { X, Y int }

var Var int

// old
/*type T struct { X int }*/
// new
type T struct { X, Y int }



// Node 定义依赖节点的数据结构
type Node struct {
	DependencyName string
	Depth          int
	GO111MODULE    bool
	NewVersion     string
	DownloadDir    string
	Parent         *Node
	Children       []*Node
	SMTConstraints []*Constraint
}

type Constraint struct {
	SymbolName string
	LowerBound string
	UpperBound string
}



// AddChild 在树节点中添加子节点
/*func (node *Node) AddChild(child *Node) {
	node.Children = append(node.Children, child)
}*/

func (node *Node) AddChild(child *Node) {
	node.Children = append(node.Children, child)
}


type U1 int //old
type U2 int
func F1(v U2) {
}

// Shape 接口定义
type Shape interface {
	Area() float64
	Round() float64
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 方法实现
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Round 方法实现
func (r Rectangle) Round() float64 {
	return C
}


func Round() float64 {
	return C
}


func TestGoroutineProfile(t *testing.T) []string {
	apilist := list.New()
	// 添加一些元素到 apilist
	apilist.PushBack("api1")
	apilist.PushBack("api2")

	// 将 apilist 转换为 []string
	var apiStrings []string
	for e := apilist.Front(); e != nil; e = e.Next() {
		apiStrings = append(apiStrings, e.Value.(string))
	}

	return apiStrings
}
