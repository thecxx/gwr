package gwr

import (
	"net/http"
	"sync/atomic"
)

type Option func(*Gwr)

func WithBuffer() Option {
	return func(g *Gwr) {

	}
}

type Gwr struct {
	r     *http.Request
	w     http.ResponseWriter
	using uint32
}

func Attach(w http.ResponseWriter, r *http.Request, opts ...Option) Gwr {
	return Gwr{w: w, r: r, using: 1}
}

func (g *Gwr) Method() string {
	return g.r.Method
}

func (g *Gwr) Param(name string) string {
	if g.r.Method != http.MethodGet {
		return g.Form(name)
	}
	return g.Query(name)
}

func (g *Gwr) Query(name string) string {
	return ""
}

func (g *Gwr) Form(name string) string {
	return ""
}

func (g *Gwr) Validate(vptr interface{}, fun func(string) error) error {
	return nil
}

func (g *Gwr) Cookie(name string, unref bool) (*http.Cookie, bool) {
	c, err := g.r.Cookie(name)
	return c, err != nil
}

func (g *Gwr) Remote() *remote {
	return &remote{}
}

func (g *Gwr) Flush() {
	if atomic.CompareAndSwapUint32(&g.using, 1, 0) {

	}
}

// func Forward(w http.ResponseWriter, r *http.Request, path string, code int) {
// 	Redirect(w, r, path, code)
// }

// func Redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
// 	http.Redirect(w, r, url, code)
// }

// func Query(w http.ResponseWriter, r *http.Request, name string) string {
// 	return ""
// }

// func QueryBind(w http.ResponseWriter, r *http.Request, vptr interface{}) {

// }

// func Form(w http.ResponseWriter, r *http.Request, name string) {

// }

// func Cookie(w http.ResponseWriter, r *http.Request, name string, unref bool) {

// }
