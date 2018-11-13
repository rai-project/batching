package batching

func NewNaive(runner func([][]byte), data <-chan []byte, opts ...Option) error {
	options := NewOptions(opts...)
	buf := [][]byte{}
	for {
		select {
		case data, ok := <-data:
			if !ok {
				return nil
			}
			buf = append(buf, data)
			if len(buf) == options.batchsize {
				runner(buf)
				buf = [][]byte{}
			}
		case <-options.ctx.Done():
			return nil
		}
	}
	return nil
}
