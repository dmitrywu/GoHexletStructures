package main

import "fmt"

type Profile struct {
	ID   int
	Tags []string
	Meta map[string]string
}

func (p Profile) Equal(other Profile) bool {
	// сравнивает ID, элементы Tags (по порядку) и пары Meta. Учитывайте, что коллекции могут быть nil.
	if p.ID != other.ID {
		return false
	}
	if len(other.Tags) != len(p.Tags) {
		return false
	}
	for i := range p.Tags {
		if p.Tags[i] != other.Tags[i] {
			return false
		}
	}
	if len(p.Meta) != len(other.Meta) {
		return false
	}
	for k, v := range p.Meta {
		if ov, ok := other.Meta[k]; !ok || ov != v {
			return false
		}
	}
	return true
}

func (p Profile) Clone() Profile {
	//создающий глубокую копию: новые срез и карта, если коллекции не пустые.
	cp := Profile{ID: p.ID}

	if len(p.Tags) > 0 {
		cp.Tags = make([]string, len(p.Tags))
		copy(cp.Tags, p.Tags)
	}

	if len(p.Meta) > 0 {
		cp.Meta = make(map[string]string, len(p.Meta))
		for k, v := range p.Meta {
			cp.Meta[k] = v
		}
	}
	return cp
}

func main() {
	fmt.Print("\033[H\033[2J")

}
