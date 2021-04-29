package bitstream

type Stream chan byte

type ReadOnlyStream <-chan byte

type WriteOnlyStream chan<- byte
