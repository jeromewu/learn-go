package linkedlist

import (
	"testing"
)

func TestAppendOneNode(t *testing.T) {
	var list LinkedList
	list.AppendNode(&Node{value: 0})
	if list.GetNode(0).value != 0 {
		t.Error("List[0].value != 0")
	}
}

func TestApeendTwoNode(t *testing.T) {
	var list LinkedList
	list.AppendNode(&Node{value: 0})
	list.AppendNode(&Node{value: 1})
	if list.GetNode(1).value != 1 {
		t.Error("List[1].value != 1")
	}
}

func TestInsertNode(t *testing.T) {
	var list LinkedList
	list.AppendNode(&Node{value: 0})
	list.AppendNode(&Node{value: 1})
	list.InsertNode(0, &Node{value: 1.5})
	if list.GetNode(1).value != 1.5 {
		t.Error("List[1].value != 1.5")
	}
}

func TestDeleteNode(t *testing.T) {
	var list LinkedList
	list.AppendNode(&Node{value: 0})
	list.AppendNode(&Node{value: 1})
	list.AppendNode(&Node{value: 2})
	list.DeleteNode(1)
	if list.GetNode(1).value != 2 {
		t.Error("List[1].value != 2")
	}
}
