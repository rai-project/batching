package batching

import (
	"sync"

	"github.com/rai-project/dlframework/steps"
)

type batch struct {
	sync.WaitGroup
}

func NewNaive(runner func([]steps.IDer), data <-chan steps.IDer, opts ...Option) (*batch, error) {
	options := NewOptions(opts...)
	buf := []steps.IDer{}
	btch := new(batch)
	for {
		select {
		case data, ok := <-data:
			if !ok {
				return btch, nil
			}
			buf = append(buf, data)

			if len(buf) == options.batchsize {
				btch.Add(1)
				tmp := make([]steps.IDer, len(buf))
				copy(tmp, buf)
				go func() {
					runner(tmp)
					btch.Done()
				}()
				buf = []steps.IDer{}
			}
		case <-options.ctx.Done():
			return btch, nil
		}
	}
	return btch, nil
}
