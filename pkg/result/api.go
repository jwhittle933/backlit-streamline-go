package result

type Result struct {
	Success interface{}
	Error   error
}

func NewSuccess(s interface{}) *Result {
	return &Result{Success: s, Error: nil}
}

func NewError(e error) *Result {
	return &Result{Success: nil, Error: e}
}

func (r *Result) isSuccess() bool {
	return r.Success != nil && r.Error == nil
}

func (r *Result) isError() bool {
	return r.Error == nil
}

func (r *Result) Next(fn func(data interface{}) *Result) *Result {
	if r.isError() {
		return r
	}

	return fn(r.Success)
}
