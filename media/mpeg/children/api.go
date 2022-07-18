package children

import (
	"github.com/jwhittle933/streamline/media/mpeg/box"
	"github.com/jwhittle933/streamline/media/mpeg/box/unknown"
)

type Registry map[string]box.Factory

func (r Registry) Names() []string {
	keys := make([]string, 0)

	for k, _ := range r {
		keys = append(keys, k)
	}

	return keys
}

func (r Registry) Get(name string) box.Factory {
	if f, ok := r[name]; ok {
		return f
	}

	return unknown.New
}

func (r Registry) Put(name string, fn box.Factory) {
	r[name] = fn
}

func (r Registry) Add(name string, fn box.Factory) {
	if _, ok := r[name]; ok {
		return
	}

	r.Put(name, fn)
}
