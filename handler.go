package handler

func New[T any]() *Handler[T] {
	return &Handler[T]{}
}

type pack[T any] struct {
	data T
	idx  int
}

type Handler[T any] struct {
	mds      []func(T, func(T))
	endpoint func(T)
}

func (h *Handler[T]) Then(mds ...func(T, func(T))) *Handler[T] {
	h.mds = append(h.mds, mds...)

	return h
}

func (h *Handler[T]) Final(end func(T)) *Handler[T] {
	h.endpoint = end

	return h
}

func (h *Handler[T]) Handle(t T) {
	p := &pack[T]{t, -1}
	h.handle(p)
}

func (h *Handler[T]) handle(p *pack[T]) {
	p.idx++
	if p.idx < len(h.mds) {
		md := h.mds[p.idx]
		md(p.data, h.next(p))
	} else if h.endpoint != nil {
		h.endpoint(p.data)
	}
}

func (h *Handler[T]) next(p *pack[T]) func(T) {
	return func(t T) {
		p.data = t
		h.handle(p)
	}
}
