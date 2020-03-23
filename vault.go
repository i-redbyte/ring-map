package main

type Vault struct {
	size        int
	values      [][2]string
	oldPosition int
	position    int
}

func NewVault(size int) *Vault {
	rm := new(Vault)
	rm.size = 0
	rm.values = make([][2]string, size)
	rm.position = 0
	return rm
}

func (v *Vault) Put(key, value string) {
	if v.size < len(v.values) {
		v.values[v.position][0] = key
		v.values[v.position][1] = value
		v.position = (v.position + 1) % len(v.values)
	} else {
		v.values[v.oldPosition][0] = key
		v.values[v.oldPosition][1] = value
	}
	if v.size < cap(v.values) {
		v.size++
	}
}

func (v *Vault) Len() int {
	return v.size
}

func (v *Vault) Keys() []string {
	var res []string
	for i := 0; i < len(v.values); i++ {
		if v.values[i][0] != "" {
			res = append(res, v.values[i][0])
		}
	}
	return res
}

func (v *Vault) Get(key string) string {
	for i := 0; i < v.size; i++ {
		if v.values[i][0] == key {
			if i == len(v.values)-1 {
				v.oldPosition = 0
			} else {
				v.oldPosition = i + 1
			}
			return v.values[i][1]
		}
	}
	return ""
}
