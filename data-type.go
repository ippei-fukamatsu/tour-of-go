package main

import (
	"fmt"
)

// データ型を急いで学ぶ
// https://dev.classmethod.jp/go/golang-7/
func main() {
	// range
	// 配列の要素数分ループをする
	fmt.Println("range ===")
	num_array := []int{2, 3, 4}
	// for loop
	for i := 0; i < len(num_array); i++ {
		fmt.Printf("index:%d,value:%d\n", i, num_array[i])
	}

	// range
	for i, value := range num_array {
		fmt.Printf("index:%d,value:%d\n", i, value)
	}

	// Array
	fmt.Println("array ===")
	var array [5]int
	fmt.Println("array: ", array)
	array[4] = 100
	fmt.Println("array[4]:", array[4])

	arrry_2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arrry_2[:]:", arrry_2[:])
	fmt.Println("arrry_2[1:4]:", arrry_2[1:4])

	// Slice
	// JavaでいうListみたいなもの
	fmt.Println("Slice ===")

	// TODO: 続き
}
