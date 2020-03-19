package main

import "fmt"

func main() {
	v := New(3)

	v.Put("key1", "value1")
	fmt.Println(v.Len())
	v.Put("key2", "value2")
	fmt.Println(v.Len())
	v.Put("key3", "value3")
	fmt.Println(v.Len())
	v.Put("key4", "value4")
	fmt.Println(v.Len())
	fmt.Println("GET:", v.Get("key4"))
	fmt.Println("Keys:", v.Keys())
}
