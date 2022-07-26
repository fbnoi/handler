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
	New[*Test]().Then(func(p *Pack[*Test], next func(*Pack[*Test])) {
		p.Data.D++
		next(p)
	}).Then(func(p *Pack[*Test], next func(*Pack[*Test])) {
		next(p)
		p.Data.D++
	}).Final(func(p *Pack[*Test]) {
		p.Data.D++
	}).Handle(P(s))

	assert.Equal(t, 3, s.D)
}
