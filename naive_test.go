package batching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	batchSize = 10
)

func TestNaive(t *testing.T) {
	dataGenerator := make(chan []byte)
	go func() {
		maxGenerator := batchSize
		for {
			if maxGenerator <= 0 {
				close(dataGenerator)
				break
			}
			dataGenerator <- []byte("http://ww4.hdnux.com/photos/41/15/35/8705883/4/920x920.jpg")
			maxGenerator--
		}
	}()
	NewNaive(
		func(data [][]byte) {
			assert.NotEmpty(t, data)
			println(len(data))
		},
		dataGenerator,
		BatchSize(batchSize),
	)
}
