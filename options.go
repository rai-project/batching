package batching

import "context"

type Options struct {
	ctx       context.Context
	batchsize int
}

type Option func(*Options)

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.ctx = ctx
	}
}

func BatchSize(batchsize int) Option {
	return func(o *Options) {
		o.batchsize = batchsize
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		ctx:       context.Background(),
		batchsize: 128,
	}
	for _, o := range opts {
		o(options)
	}
	return options
}
