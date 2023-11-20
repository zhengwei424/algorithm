package main

import "fmt"

// List 数据结构线性表之顺序表
type List struct {
	// 表长度
	Len int
	// 指向顺序表空间的指针
	Ptr *[]interface{}
	// 表容量(表示长度的最大值)
	Capacity int
}

// ListInit 初始化
func (l *List) ListInit(capacity int) {
	l.Capacity = capacity
	l.Len = 0
	tmp := make([]interface{}, capacity)
	l.Ptr = &tmp
}

// ListIsEmpty 判断为空
func (l *List) ListIsEmpty() bool {
	if l.Len == 0 {
		return true
	} else {
		return false
	}
}

// ListIsFull 判断长度达到最大值
func (l *List) ListIsFull() bool {
	if l.Len == l.Capacity {
		return true
	}
	return false
}

// ListGet 根据下标get元素
func (l *List) ListGet(index int) (interface{}, bool) {
	if index < 0 || index > l.Len-1 {
		return nil, false
	}
	return (*l.Ptr)[index], true
}

// 查找

// 1. 根据传入的值，返回第一个匹配的元素索引
func (l *List) ListQuery(elem interface{}) (int, bool) {
	for index, item := range *l.Ptr {
		if item == elem {
			return index, true
		}
	}
	return -1, false
}

// 2. 寻找给定元素的前驱并返回
func (l *List) ListQueryElemPre(elem interface{}) (interface{}, bool) {
	index, _ := l.ListQuery(elem)
	if index == -1 || index == 0 {
		return nil, false
	}
	return (*l.Ptr)[index-1], true
}

// 3. 寻找给定元素的后继并返回索引
func (l *List) ListQueryElemNext(elem interface{}) (interface{}, bool) {
	index, _ := l.ListQuery(elem)
	if index == -1 || index == l.Len-1 {
		return nil, false
	}
	return (*l.Ptr)[index+1], true
}

// 向指定索引插入元素
func (l *List) ListInsert(index int, elem interface{}) bool {
	// 允许index=l.Len，即等效于在原表的末尾追加一个;也允许为空
	if index < 0 || index > l.Len || l.ListIsFull() {
		return false
	}
	lv := *l.Ptr
	for i := l.Len - 1; i >= index; i-- {
		lv[i] = lv[i-1]
	}
	lv[index] = elem
	l.Len++
	return true

}

// 删除指定索引的元素
func (l *List) ListDelete(index int) bool {
	if l.ListIsEmpty() || index < 0 || index > l.Len-1 {
		return false
	}
	lv := *l.Ptr
	for i := index; i < l.Len-1; i++ {
		lv[i] = lv[i+1]
	}
	l.Len--
	return true
}

// 遍历列表
func (l *List) ListTraverse() {
	// 不能使用for...range...遍历，这会以capacity的长度遍历整个list，而不是length
	for i := 0; i < l.Len; i++ {
		fmt.Println((*l.Ptr)[i])
	}

}

// 删除列表，释放内存空间
func (l *List) ListClear() {
	l.Len = 0
	l.Ptr = nil
}

// 测试
func main() {
	var li List
	li.ListInit(4)
	fmt.Println(li.ListIsEmpty())
	fmt.Println(li.ListIsFull())

	type s struct {
		name string
		age  int
	}
	student1 := s{name: "abc", age: 1}
	student2 := s{name: "def", age: 2}
	li.ListInsert(0, student1)
	li.ListInsert(1, student2)
	li.ListInsert(2, "1000")
	li.ListInsert(3, "xyz")
	li.ListInsert(4, "aaa")
	fmt.Println("++++++++++++++++++++++++++")
	li.ListTraverse()
	li.ListDelete(2)
	fmt.Println("++++++++++++++++++++++++++")
	li.ListTraverse()
	li.ListClear()
	fmt.Println(li.ListIsEmpty())

}
