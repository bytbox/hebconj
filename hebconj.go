package main

type Node interface {
	Process()
	Status() float64
}

type DumbNode struct {
	weight []float64
	child  []Node
	output float64
}

type Input struct {
	status float64
}

func (i *Input) Process() {
	// nothing to be done
}

func (i *Input) Status() float64 {
	return i.status
}

func (i *DumbNode) Process() {
	i.output = 0
	for c := range i.weight {
		i.output += i.weight[c] * i.child[c].Status()
	}
}

func (i *DumbNode) Status() float64 {
	return i.output
}

func main() {
	println("Hello, world")

	i1 := &Input{0.0}
	i2 := &Input{0.1}
	n := &DumbNode{[]float64{0.1, 0.6}, []Node{i1, i2}, 0}
	n.Process()
	println(n.Status())
	train(n)
}

func train(n Node) {
	// run 20 training rounds on n (should be enough)
}
