package main

import (
	"fmt"
	"unsafe"
)

type OrderInfo struct {
	OrderCode   rune
	Amount      int
	OrderNumber uint16
	Items       []string
	IsReady     bool
}

type SmallOrderInfo struct {
	Amount      int
	Items       []string
	OrderCode   rune
	OrderNumber uint16
	IsReady     bool
}

func main() {
	var o OrderInfo
	fmt.Println("OrderInfo:")
	fmt.Printf("  Total size:          %d bytes\n", unsafe.Sizeof(o))
	fmt.Printf("  OrderCode  offset:   %d\n", unsafe.Offsetof(o.OrderCode))
	fmt.Printf("  Amount     offset:   %d\n", unsafe.Offsetof(o.Amount))
	fmt.Printf("  OrderNumber offset:  %d\n", unsafe.Offsetof(o.OrderNumber))
	fmt.Printf("  Items      offset:   %d\n", unsafe.Offsetof(o.Items))
	fmt.Printf("  IsReady    offset:   %d\n", unsafe.Offsetof(o.IsReady))

	fmt.Println()

	var s SmallOrderInfo
	fmt.Println("SmallOrderInfo:")
	fmt.Printf("  Total size:          %d bytes\n", unsafe.Sizeof(s))
	fmt.Printf("  Amount     offset:   %d\n", unsafe.Offsetof(s.Amount))
	fmt.Printf("  Items      offset:   %d\n", unsafe.Offsetof(s.Items))
	fmt.Printf("  OrderCode  offset:   %d\n", unsafe.Offsetof(s.OrderCode))
	fmt.Printf("  OrderNumber offset:  %d\n", unsafe.Offsetof(s.OrderNumber))
	fmt.Printf("  IsReady    offset:   %d\n", unsafe.Offsetof(s.IsReady))
}
