package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("======> Start:")

	v, err := NewVault(3)
	if err != nil {
		log.Fatal(err)
	}
	v.Put("key1", "value1!")
	fmt.Println(v.Len()) // 1
	v.Put("key2", "value2!")
	fmt.Println(v.Len()) // 2
	v.Put("key3", "value3!")
	fmt.Println(v.Len()) // 3
	fmt.Println(v.Get("key1"))
	v.Put("key4", "value4!")
	fmt.Println(v.Len())
	fmt.Println(v.Keys())
	v.Clear()
	fmt.Println(v.Len())
	fmt.Println("======> End.")
}
