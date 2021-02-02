package lib

import "sync"

type Node struct {
	Left  *Node
	Right *Node
	Data  interface{}
}

func (n *Node) Length() int {
	if n.Left != nil {
		return 1 + n.Left.Length()
	} else {
		return 0
	}
}

func (n *Node) Add(node *Node) {
	if n.Left == nil {
		n.Left = node
	} else if n.Right == nil {
		n.Right = node
	} else {
		if n.Left.Length() <= n.Right.Length() {
			n.Left.Add(node)
		} else {
			n.Right.Add(node)
		}
	}
}

func (n *Node) SearchAsync(data interface{}, waitGroup *sync.WaitGroup, resultChannel chan<- *Node) {

	waitGroup.Add(1)
	defer waitGroup.Done()

	if n.Data == data {
		resultChannel <- n
	} else {
		if n.Left != nil {
			go n.Left.SearchAsync(data, waitGroup, resultChannel)
		}
		if n.Right != nil {
			go n.Right.SearchAsync(data, waitGroup, resultChannel)
		}
	}
}
