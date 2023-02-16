package log

type noop struct{}

func (n *noop) Read(...ReadOption) ([]Record, error) {
	return nil, nil
}

func (n *noop) Write(Record) error {
	return nil
}

func (n *noop) Stream() (Stream, error) {
	return nil, nil
}

func NewLog(opts ...Option) Log {
	return new(noop)
}
