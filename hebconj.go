package main

type Node interface {
	Process()
	Status() bool
}

type DumbNode struct {
	weight1 float64
	weight2 float64
	child1  Node
	child2  Node
}

type Input struct {
	status bool
}

func (i *Input) Process() {

}

func (i *Input) Status() bool {
	return i.status
}

func (i *DumbNode) Process() {

}

func main() {
	println("Hello, world")

	i1 := Input{false}
	i2 := Input{false}

	n := DumbNode{0.9, 0.1, &i1, &i2}
	train(n)
}

func train(n Node) {
	// run 20 training rounds on n (should be enough)
}

