package main

type Node interface {
	Process()
	Status() float64
}

type DumbNode struct {
	weight1 float64
	weight2 float64
	child1  Node
	child2  Node

	output float64
}

type Input struct {
	status float64
}

func (i *Input) Process() {

}

func (i *Input) Status() float64 {
	return i.status
}

func (i *DumbNode) Process() {
	i.output = i.weight1 * i.child1.Status() + i.weight2 * i.child2.Status()
}

func (i *DumbNode) Status() float64 {
	return i.output
}

func main() {
	println("Hello, world")

	i1 := Input{0.0}
	i2 := Input{0.0}

	n := &DumbNode{0.9, 0.1, &i1, &i2, 0.0}
	n.Process()
	println(n.Status())
	train(n)
}

func train(n Node) {
	// run 20 training rounds on n (should be enough)
}
