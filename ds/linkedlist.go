                                                                                                                                                                                                                                                                                                                                                                                                                                                                   linkedlist.go                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   
package ds

import "fmt"

type LinkedList struct {
    Head *Node
    Tail *Node
    Size int
}

// Functions

func (l *LinkedList) Insert(value string) {
        var newNode Node
        newNode.data = value
        newNode.Next = nil
        if l.Size == 0 {
                l.Head = &newNode
                l.Tail = &newNode
        } else {
                l.Tail.Next = &newNode
                l.Tail = &newNode
        }
	l.Size++
} // insert at the tail
func (l *LinkedList) InsertAt(position int, value string) error {
        var prevNode *Node
        var curNode *Node
        curNode = l.Head
        for i := 0; i <= position; i++ {
                if i == position {
                        if i == 0 || i == position {
                                var newNode Node
                                newNode.data = value
                                newNode.Next = nil
                                if i == 0 {
                                        l.Head = &newNode
                                }
                                l.Tail = &newNode
                        } else {
                                var newNode Node
                                newNode.data = value
                                newNode.Next = prevNode.Next
                                prevNode.Next = &newNode
                        }
                        l.Size++
                } else {
                        if curNode.Next == nil {
                                return fmt.Errorf("Out of Bounds")
                        }
                        if i != 0 {
                                prevNode = curNode
                        }
                        curNode = curNode.Next
                }
        }
	return nil
} //inserts at a position, returns true if success or false if not, like if pos doesn't exist
func (l *LinkedList) Remove(value string) error {
        var tempNode *Node
        var prevNode *Node
        tempNode = l.Head
        prevNode = l.Head
        var err = true
        for i := 0; i <= l.Size; i++ {
                if tempNode.data == value {
                        prevNode.Next = tempNode.Next
                        err = false
                        l.Size--
                        break
                } else {
                        prevNode = tempNode
                        tempNode = tempNode.Next
                }
        }
	if err {
                return fmt.Errorf("No Occurences of value")
        }
	return nil
}// remove first occurrence of the value
func (l *LinkedList) RemoveAll(value string) error {
        var tempNode *Node
        var prevNode *Node
        var err = true
        tempNode = l.Head
        prevNode = l.Head
        for i := 0; i < l.Size; i++ {
                if tempNode.data == value {
                        if tempNode == l.Head {
                                tempNode = tempNode.Next
                                l.Head = tempNode
                                prevNode = tempNode
                                err = false
                                l.Size--
                        } else {
                                prevNode.Next = tempNode.Next
                                err = false
                                l.Size--
                        }
                } else {
                        prevNode = tempNode
                        tempNode = tempNode.Next
                }
        }
	if err {
                return fmt.Errorf("No Occurences of value")
        }
	return nil
} //remove all occurrences of a value
func (l *LinkedList) RemoveAt(pos int) error {
        var tempNode *Node
        var prevNode *Node
        tempNode = l.Head
        if pos > l.Size {
                return fmt.Errorf("Out of Bounds")
        }
	for i := 0; i <= pos; i++ {
                if i == pos {
                        prevNode.Next = tempNode.Next
                        l.Size--
                        break
                } else {
                        prevNode = tempNode
                        tempNode = tempNode.Next
                }
        }
	return nil
} // remove at a position, if index exists
func (l *LinkedList) IsEmpty() bool {
        if l.Size == 0 {
                return true
        } else {
                return false
        }
} // checks if the linked list is empty
func (l *LinkedList) GetSize() int {
        return l.Size
} // get size of ll
func (l *LinkedList) Reverse() {
        var prevNode *Node
        var tempNode *Node
        var nextNode *Node
        tempNode = l.Head
        prevNode = nil
        for tempNode != nil {
                nextNode = tempNode.Next
                tempNode.Next = prevNode
                prevNode = tempNode
                tempNode = nextNode 
        }
	temp := l.Head
        l.Head = l.Tail
        l.Tail = temp
        return
} //reverse the list
func (l *LinkedList) PrintList() {
        var tempNode *Node
        tempNode = l.Head
        for tempNode != nil {
                fmt.Println(tempNode.data)
                tempNode = tempNode.Next
        }
	return
} //print the list 


