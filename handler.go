package handler

func P[T any](d T) *Pack[T] {
	return &Pack[T]{
		Data: d,
		idx:  -1,
	}
}

func New[T any]() *Handler[T] {
	return &Handler[T]{}
}

type Pack[T any] struct {
	Data T
	idx  int
}

type Handler[T any] struct {
	mds      []func(*Pack[T], func(*Pack[T]))
	endpoint func(*Pack[T])
}

func (h *Handler[T]) Then(mds ...func(*Pack[T], func(*Pack[T]))) *Handler[T] {
	h.mds = append(h.mds, mds...)

	return h
}

func (h *Handler[T]) Final(end func(*Pack[T])) *Handler[T] {
	h.endpoint = end

	return h
}

func (h *Handler[T]) Handle(p *Pack[T]) {
	p.idx++
	if p.idx < len(h.mds) {
		md := h.mds[p.idx]
		md(p, h.next())
	} else if h.endpoint != nil {
		h.endpoint(p)
	}
}

func (h *Handler[T]) next() func(*Pack[T]) {
	return func(p *Pack[T]) {
		h.Handle(p)
	}
}
