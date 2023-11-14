package wsmantesting

type ClientMock struct {
	Err      error
	Response []byte
}

func (c *ClientMock) Post(msg string) (response []byte, err error) {
	return c.Response, c.Err
}
