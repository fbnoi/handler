package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	D int
}

func TestHandler(t *testing.T) {
	s := &Test{0}
	New[*Test]().Then(func(p *Test, next func(*Test)) {
		p.D++
		next(p)
	}).Then(func(p *Test, next func(*Test)) {
		p.D++
		next(p)
	}).Final(func(p *Test) {
		p.D++
	}).Handle(s)

	assert.Equal(t, 3, s.D)
}
