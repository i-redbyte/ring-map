package main

type RingMap struct {
	size       int
	key, value []string
	position   int
}

func New(size int) *RingMap {
	rm := new(RingMap)
	rm.size = 0
	rm.key = make([]string, size)
	rm.value = make([]string, size)
	rm.position = 0
	return rm
}

func (rm *RingMap) Put(key, value string) {
	rm.key[rm.position] = key
	rm.value[rm.position] = value
	rm.position = (rm.position + 1) % len(rm.key)
	if rm.size < cap(rm.key) {
		rm.size++
	}
}

func (rm *RingMap) Len() int {
	return rm.size
}

func (rm *RingMap) Keys() []string {
	return rm.key
}

func (rm *RingMap) Get(key string) string {
	for i := 0; i < rm.size; i++ {
		if rm.key[i] == key {
			return rm.value[i]
		}
	}
	return ""
}
