package main

import "fmt"

func main() {
	v := NewVault(3)
	v.Put("key1", "value1")
	fmt.Println(v.Len()) // 1
	v.Put("key2", "value2")
	fmt.Println(v.Len()) // 2
	v.Put("key3", "value3")
	fmt.Println(v.Len()) // 3
	fmt.Println(v.Get("key1"))
	v.Put("key4", "value4")
	fmt.Println(v.Len())
	fmt.Println(v.Keys())
}
