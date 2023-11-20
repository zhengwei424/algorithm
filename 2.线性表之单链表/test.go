package main

import "fmt"

// Node 单链表节点
type Node struct {
	Data interface{} // 节点数据域
	Next *Node       // 节点指针域
}

// LinkList
type LinkList struct {
	Header *Node // 指向第一个节点(带头结点的单链表就指向头结点，不带头结点的单链表就指向首元节点）
	Length int
}

// +++++++++++++++带头结点的单链表+++++++++++++++++++++++++++++
// 单链表初始化
func LinkListInit() *LinkList {
	l := new(Node)
	l.Data = nil
	l.Next = nil
	return &LinkList{l, 0}
}

// 向单链表中插入结点
func (l *LinkList) Insert(elem interface{}) {
	elemPtr := &Node{elem, nil}
	if l.Length == 0 {
		l.Header.Next = elemPtr
	} else if l.Length > 0 {
		elemPtr.Next = l.Header.Next
		l.Header.Next = elemPtr
	}
	l.Length++
}

// 向单链表末尾追加结点
func (l *LinkList) Append(elem interface{}) {
	elemPtr := &Node{elem, nil}
	if l.Length == 0 {
		l.Header.Next = elemPtr
	} else if l.Length > 0 {
		current := l.Header
		for current.Next != nil { // 遍历单链表，直到最后一个结点
			current = current.Next
		}
		current.Next = elemPtr
	}
	l.Length++
}

// 在指定索引的结点之后插入一个结点
func (l *LinkList) InsertTo(index int, elem interface{}) bool {
	// 只允许在结点之间插入
	if index >= l.Length-1 || index < 0 {
		return false
	}
	current := l.Header
	elemPtr := &Node{elem, nil}
	for i := 0; i < l.Length-1; i++ {
		current = current.Next
		if i == index {
			elemPtr.Next = current.Next
			current.Next = elemPtr
		}
	}
	l.Length++
	return true

}

// 删除指定索引的结点
func (l *LinkList) Delete(index int) bool {
	if index >= l.Length || index < 0 {
		return false
	}
	j := 0
	current := l.Header
	for j < index {
		current = current.Next
		fmt.Println("+++++: ", j)
		j++
	}
	current.Next = current.Next.Next
	//tmp := current.Next.Next
	//current.Next = tmp
	l.Length--
	return true
}

// 遍历链表
func (l *LinkList) Scan() {
	current := l.Header
	for i := 0; i < l.Length; i++ {
		current = current.Next
		fmt.Printf("第%d的结点数据是%v，下一个结点地址是%v, %p\n", i, current.Data, current.Next, current.Next)
	}
}

// 单链表返回指定索引的值
//func (l *LinkList) getElem(index int) interface{} {
//	l = l
//}

//

// +++++++++++++++不带头结点的单链表+++++++++++++++++++++++++++++

func main() {
	ll := LinkListInit()
	ll.Append("aa")
	ll.Append(123)
	ll.Append("cc")
	ll.Append("dd")
	ll.Scan()
	ll.Delete(2)
	ll.Scan()
	ll.Insert("ee")
	ll.InsertTo(1, "ff")
	ll.Scan()
}
