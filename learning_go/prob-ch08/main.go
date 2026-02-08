package main

import (
	"fmt"
)

// i
type Doublable interface {
	int | float64
}

func Double[T Doublable](d T) T {
	return d * 2
}

// ii
type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type MyInt int

func (i MyInt) String() string {
	return "This is MyInt!"
}

func Print[T Printable](p T) {
	fmt.Println(p.String())
}

// iii
type ListItem[T comparable] struct {
	Val      T
	NextItem *ListItem[T]
}

type SinglyLinkedList[T comparable] struct {
	FirstItem *ListItem[T]
	Size      int
}

func (list *SinglyLinkedList[T]) Add(itemToAdd ListItem[T]) {
	if list.FirstItem == nil {
		list.FirstItem = &itemToAdd
		list.Size++
		return
	}
	curItem := list.FirstItem
	for {
		if curItem.NextItem == nil {
			curItem.NextItem = &itemToAdd
			list.Size++
			return
		}
		curItem = curItem.NextItem
	}
}

func (list *SinglyLinkedList[T]) Insert(itemToAdd ListItem[T], index int) {
	if index < 0 || index >= list.Size {
		// 存在しないインデックス
		return
	} else if index == 0 {
		// 先頭に追加する場合
		itemToAdd.NextItem = list.FirstItem
		list.FirstItem = &itemToAdd
		list.Size++
		return
	}

	curIndex := 0
	curItem := list.FirstItem

	// indexの1つまでの要素を見る
	for curIndex < index-1 {
		curIndex++
		curItem = curItem.NextItem
	}

	// indexの位置にitemToAddをセット
	itemToAdd.NextItem = curItem.NextItem
	curItem.NextItem = &itemToAdd
	list.Size++
}

func (list *SinglyLinkedList[T]) Index(item ListItem[T]) int {
	curIndex := 0
	curItem := list.FirstItem
	for {
		if curItem == nil {
			curIndex = -1
			break
		} else if curItem.Val == item.Val {
			break
		}
		curIndex++
		curItem = curItem.NextItem
	}
	return curIndex
}

func main() {
	// var i MyInt = 1
	// Print(i)

	// SinglyLinkedListのサンプルコード
	// ListItem[int]を生成
	item1 := ListItem[int]{Val: 10}
	item2 := ListItem[int]{Val: 20}
	item3 := ListItem[int]{Val: 30}
	item4 := ListItem[int]{Val: 40}
	item5 := ListItem[int]{Val: 50}

	// SinglyLinkedList[int]を生成
	var list SinglyLinkedList[int]

	// 5つAddしてみる
	list.Add(item1)
	list.Add(item2)
	list.Add(item3)
	list.Add(item4)
	list.Add(item5)

	// listの内容をfmt.Printlnで確認
	fmt.Println("After Add:")
	cur := list.FirstItem
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.NextItem
	}

	// Insert: 2番目のindex(インデックス2)に新しい要素を挿入
	itemInsert := ListItem[int]{Val: 999}
	list.Insert(itemInsert, 2)
	fmt.Println("\nAfter Insert(999 at index 2):")
	cur = list.FirstItem
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.NextItem
	}

	// Index: item3(Val: 30)のインデックスを調べる
	idx := list.Index(item3)
	fmt.Println("\nIndex of 30:", idx)

	// Index: itemInsert(Val:999)のインデックスも調べる
	idx = list.Index(itemInsert)
	fmt.Println("\nIndex of 999:", idx)
}
