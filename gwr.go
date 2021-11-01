package gwr

import (
	"net/http"
	"sync/atomic"
)

type ErrorConstructor func(s string) error

type Gwr struct {
	w     http.ResponseWriter
	r     *http.Request
	using uint32
}

// Attach to a pair of w/r and return a new Gwr.
func Attach(w http.ResponseWriter, r *http.Request) Gwr {
	return Gwr{w: w, r: r, using: 1}
}

func (g *Gwr) Extract(vptr interface{}, fun ErrorConstructor) error {
	return nil
}

func (g *Gwr) Flush() {
	if atomic.CompareAndSwapUint32(&g.using, 1, 0) {

	}
}

func Extract(r *http.Request, vptr interface{}, fun ErrorConstructor) error {
	return nil
}
