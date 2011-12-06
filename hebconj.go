package main

import "math"

type Node interface {
	Process()
	Status() bool
}

type DumbNode struct {
	weight1 float64
	weight2 float64
	child1  Node
	child2  Node

	output bool
}

type Input struct {
	status bool
}

func (i *Input) Process() {
	// nothing to be done
}

func (i *Input) Status() bool {
	return i.status
}

func (i *DumbNode) Process() {
	i.output = math.Floor(i.weight1*i.weight2) == 1
}

func (i *DumbNode) Status() bool {
	return i.output
}

func main() {
	println("Hello, world")

	i1 := Input{false}
	i2 := Input{false}

	n := &DumbNode{0.9, 0.1, &i1, &i2, false}
	train(n)
}

func train(n Node) {
	// run 20 training rounds on n (should be enough)
}
