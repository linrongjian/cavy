package wswrap

type IWswrap interface {
	Init() error

	// Dial() (*Conn, *http.Response, error)
}
