package result

type Binder func(data interface{}) Result

// Result represents a system operation that can succeed or fail
type Result interface {
	Bind(b Binder) Result
	BindAll(bs ...Binder) Result
	Ok() interface{}
	IsOk() bool
	Err() error
	IsErr() bool
}

// Wrap converts data to a Result to perform resilient operations on data
func Wrap(data interface{}) Result {
	return Ok{data: data}
}

func WrapErr(e error) Result {
	return Err{err: e}
}

type Ok struct {
	data interface{}
}

func (o Ok) Bind(b Binder) Result {
	return b(o)
}

func (o Ok) BindAll(bs ...Binder) Result {
	var out Result
	for _, b := range bs {
		out = b(o)
	}

	return out
}

func (o Ok) Ok() interface{} {
	return o.data
}

func (o Ok) IsOk() bool {
	return true
}

func (o Ok) IsErr() bool {
	return false
}

func (o Ok) Err() error {
	return nil
}

type Err struct {
	err error
}

func (e Err) Bind(b Binder) Result {
	return e
}

func (e Err) BindAll(bs ...Binder) Result {
	return e
}

func (e Err) Ok() interface{} {
	return nil
}

func (e Err) Err() error {
	return e.err
}

func (e Err) IsOk() bool {
	return false
}

func (e Err) IsErr() bool {
	return true
}

type PipelineFn func(r Result) Result

func Pipeline(r Result, bs ...Binder) Result {
	return r.BindAll(bs...)
}
