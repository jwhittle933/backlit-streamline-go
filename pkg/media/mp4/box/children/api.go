package children

import "github.com/jwhittle933/streamline/pkg/media/mp4/box"

type BoxFactory func(*box.Info) box.Boxed

type Registry map[string]BoxFactory

func (r Registry) Names() []string {
	keys := make([]string, 0)

	for k, _ := range r {
		keys = append(keys, k)
	}

	return keys
}