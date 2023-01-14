package resp

type Resp struct {
	Code int
	Msg  string
	Data interface{} `json:"data,omitempty"`
}

func R(code int, msg string, data ...interface{}) Resp {
	r := Resp{
		Code: code,
		Msg:  msg,
	}
	if len(data) > 0 {
		r.Data = data[0]
	}
	return r
}

func E(msg string) Resp {
	return Resp{Code: -1, Msg: msg}
}

func S(msg string, data ...interface{}) Resp {
	return R(0, msg, data...)
}
