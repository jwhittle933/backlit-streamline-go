package result

// Result represents a system operation that can succeed or fail
type Result interface {
	Bind(func(data interface{}) Result) Result
	Ok() interface{}
	Err() error
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

func (o Ok) Bind(fn func(data interface{}) Result) Result {
	return fn(o)
}

func (o Ok) Ok() interface{} {
	return o.data
}

func (o Ok) Err() error {
	return nil
}

type Err struct {
	err error
}

func (e Err) Bind(_ func(data interface{}) Result) Result {
	return e
}

func (e Err) Ok() interface{} {
	return nil
}

func (e Err) Err() error {
	return e.err
}