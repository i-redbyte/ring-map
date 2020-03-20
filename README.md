**An example of a solution to a problem:**

```go
package main

import (
	"fmt"
)

// Vault structure stores limited number of key-value elements.
// It stores most recently addressed elements, so if it already stores maximum elements allowed,
// when we add next one, it should delete least recently used one (either by Get or Put).
// type Vault struct {
// ...
// }

func main() {
	v := NewVault(3) // Create vault to store 3 key-value elements (string: string)

	v.Put("key1", "value1") // Add "key1": "value1"
	fmt.Println(v.Len())    // 1

	v.Put("key2", "value2") // Add "key2": "value2"
	fmt.Println(v.Len())    // 2

	v.Put("key3", "value3") // Add "key3": "value3"
	fmt.Println(v.Len())    // 3

	fmt.Println(v.Get("key1")) // value1

	v.Put("key4", "value4") // Add "key3": "value3" and remove "key2" as least recently used one
	fmt.Println(v.Len())    // 3

	fmt.Println(v.Keys()) // [key1 key4 key3] // keys order doesn't matter here
}

```