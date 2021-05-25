package ledg_util

import (
	"fmt"
	"github.com/filecoin-project/lotus/chain/types"
)

//type Node struct {
//	Msg *types.Message
//	MethReturn []byte
//	CallDepth uint64
//}


//type Node struct {
//	Amount int
//}
//type Node struct {
//	p ledger.ActorMethodParams
//}

func (n *ActorMethodParams) String() string {
	return fmt.Sprint(n.Msg.Cid().String())
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*ActorMethodParams
	count int
}

func (s *Stack)  GetCount() int {
	return s.count
}

// Push adds a node to the stack.
func (s *Stack) Push(n *ActorMethodParams) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *ActorMethodParams {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*ActorMethodParams, size),
		size:  size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*ActorMethodParams
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.
func (q *Queue) Push(n *ActorMethodParams) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*ActorMethodParams, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *ActorMethodParams {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

type ActorMethodParams struct {
	Msg   *types.Message
	Ret   []byte
	Depth uint64

}


