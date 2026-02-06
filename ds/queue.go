package ds

type Stack struct {
    list LinkedList
}

func (s *Stack) Push(value string) {
        var tempNode Node
        tempNode.data = value
        tempNode.Next = s.list.Head
        s.list.Head = &tempNode
        s.list.Size++
} // add new head node
func (s *Stack) Pop() (string, bool) {
        if s.list.Size == 0 {
                return "No", false
        }
	temp := s.list.Head.data
        s.list.Head = s.list.Head.Next
        s.list.Size--
        return temp, true
}// remove the head
