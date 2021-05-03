package children

import (
	"github.com/jwhittle933/streamline/pkg/media/mp4/box"
	"github.com/jwhittle933/streamline/pkg/media/mp4/box/unknown"
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
	if fac, ok := r[name]; ok {
		return fac
	}

	return unknown.New
}