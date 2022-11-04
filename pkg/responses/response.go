package responses

type Response struct {
	Status string
	Msg    string
}

func (r *Response) New(status string, msg string) {
	r.Status = status
	r.Msg = msg
}
