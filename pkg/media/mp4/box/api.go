package box

type Box interface {
	Type() string
	Version() uint8
}
