package NumbersNode

import (
	"sync"
	"unicode/utf8"
)

type Node struct {
	sync.Mutex
	Value interface{}
	Leaves []*Node
	IsEnd bool
}


func NewNode () *Node {
	var node = &Node{
		Leaves:make([]*Node, 10),
		IsEnd:false,
	}
	return node
}


func (N *Node) Insert(value interface{}, number string) bool {

	if len(number) == 0 {
		N.Lock()
		N.IsEnd = true
		N.Value = value
		N.Unlock()
		return true
	}

	if N.IsEnd {
		return false
	}

	firstNum := getFirstChar(number)
	address := getAddress(firstNum)

	if address == -1 {
		return false
	}

	if N.Leaves[address] == nil {
		node := NewNode()
		N.Leaves[address] = node
	}

	return N.Leaves[address].Insert(value, removeFirstChar(number))
}


func (N *Node) Search (number string) interface{}{

	if N.IsEnd {
		return N.Value
	}

	firstNum := getFirstChar(number)
	address := getAddress(firstNum)

	if address == -1 {
		return nil
	}

	if N.Leaves[address] != nil {
		return N.Leaves[address].Search(removeFirstChar(number))
	}

	return nil

}

func (N *Node) Update (value interface{}, number string) bool {

	if N.IsEnd {
		N.Lock()
		N.Value = value
		N.Unlock()
		return true
	}

	firstNum := getFirstChar(number)
	address := getAddress(firstNum)

	if address == -1 {
		return false
	}

	if N.Leaves[address] != nil {
		return N.Leaves[address].Update(value, removeFirstChar(number))
	}

	return false
}

func (N *Node) Delete (number string) bool {

	if N.IsEnd {
		N.Value = nil
		return true
	}

	firstNum := getFirstChar(number)
	address := getAddress(firstNum)

	if address == -1 {
		return false
	}

	if N.Leaves[address] != nil {
		return N.Leaves[address].Delete(removeFirstChar(number))
	}

	return false

}

func getFirstChar(str string) string  {
	return str[:1]
}

func removeFirstChar(str string) string {
	_, i := utf8.DecodeRuneInString(str)
	return str[i:]
}


func getAddress(number string ) int {
	switch number {
	case "0" :
		return 0
	case "1" :
		return 1
	case "2":
		return 2
	case "3" :
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	default:
		return -1
	}
}