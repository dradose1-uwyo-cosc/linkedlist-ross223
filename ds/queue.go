package ds

import "fmt"

type Queue struct {
    list LinkedList
}

func (q *Queue) Push(value string) {
        var tempNode Node
        tempNode.data = value
        if q.list.Size == 0 {
                q.list.Head = &tempNode
                q.list.Tail = &tempNode
        } else {
                q.list.Tail.Next = &tempNode
                q.list.Tail = &tempNode
        }
	q.list.Size++
}// add tail node
func (q *Queue) Pop() (string, error) {
        if q.list.Size == 0 {
                return "No", fmt.Errorf("List is Empty")
        } 
	temp := q.list.Head.data
        q.list.Head = q.list.Head.Next
        q.list.Size--
        return temp, nil
}// remove the head
