package children

import (
	box2 "github.com/jwhittle933/streamline/media/mp4/box"
	unknown2 "github.com/jwhittle933/streamline/media/mp4/box/unknown"
)

type Registry map[string]box2.Factory

func (r Registry) Names() []string {
	keys := make([]string, 0)

	for k, _ := range r {
		keys = append(keys, k)
	}

	return keys
}

func (r Registry) Get(name string) box2.Factory {
	if fac, ok := r[name]; ok {
		return fac
	}

	return unknown2.New
}
