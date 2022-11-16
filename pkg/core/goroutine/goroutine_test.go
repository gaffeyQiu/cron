package goroutine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	ch := make(chan int, 1)
	setN := func() {
		ch <- 1
		panic("exec f() has panic")
	}
	assert.Panics(t, setN)

	Run(setN)
	res := <-ch
	assert.Equal(t, 1, res)
}

func TestId(t *testing.T) {
	assert.Greater(t, Id(), uint64(1))
}
