// Package sinf (Protection Scheme Information)
package sinf

const (
	SINF string = "sinf"
)

type Box struct {
	//
}

func (Box) Type() string {
	return SINF
}