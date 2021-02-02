package lib

import "sync"

type Network struct {
	Node
}

func (n Network) Add(node *Node) {
	n.Node.Add(node)
}

func (n Network) SearchAsync(data interface{}) *Node {
	var waitGroup sync.WaitGroup

	var resultChannel = make(chan *Node)

	n.Node.SearchAsync(data, &waitGroup, resultChannel)

	var quitChannel = make(chan int)

	go func() {
		waitGroup.Wait()
		quitChannel <- 0
	}()

	for {
		select {
		case result := <-resultChannel:
			return result
		case <-quitChannel:
			return nil
		}
	}

}
