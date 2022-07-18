package result

type Binder func(data interface{}) Result
type Tee func(data interface{})

// Result represents a system operation that can succeed or fail
type Result interface {
	Bind(b Binder) Result
	Tee(t Tee) Result
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
		out = b(o.data)
	}

	return out
}

func (o Ok) Tee(t Tee) Result {
	t(o.data)
	return o
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

func (e Err) Bind(Binder) Result {
	return e
}

func (e Err) BindEither(_ Binder, errBind Binder) Result {
	return errBind(e)
}

func (e Err) BindAll(...Binder) Result {
	return e
}

func (e Err) Tee(Tee) Result {
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

func Pipe(r Result, bs ...Binder) Result {
	return r.BindAll(bs...)
}
