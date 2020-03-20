package main

type RingMap struct {
	size     int
	values   [][2]string
	position int
}

func New(size int) *RingMap {
	rm := new(RingMap)
	rm.size = 0
	rm.values = make([][2]string, size)
	rm.position = 0
	return rm
}

func (rm *RingMap) Put(key, value string) {
	rm.values[rm.position][0] = key
	rm.values[rm.position][1] = value
	rm.position = (rm.position + 1) % len(rm.values)
	if rm.size < cap(rm.values) {
		rm.size++
	}
}

func (rm *RingMap) Len() int {
	return rm.size
}

func (rm *RingMap) Keys() []string {
	res := make([]string, len(rm.values))
	for i := 0; i < len(rm.values); i++ {
		if rm.values[i][0] != "" {
			res = append(res, rm.values[i][0])
		}
	}
	return res
}

func (rm *RingMap) Get(key string) string {
	for i := 0; i < rm.size; i++ {
		if rm.values[i][0] == key {
			return rm.values[i][1]
		}
	}
	return ""
}
