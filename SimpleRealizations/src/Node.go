package src

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func newNode(value int) (a *Node) {
	a.value = value
	return a
}
func (n *Node) init(value int) {

}
