package plug

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

type P struct {
	name string
}

func (m *P) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(nil, nil)
}

func mgen(name string) PlugFunc {
	if name != "" {
		return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			fmt.Println(name)
			next(nil, nil)
		}
	}

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		next(nil, nil)
	}
}

func TestPlugFunc(t *testing.T) {
	builder := NewBuilder()

	for i := 0; i < 10; i++ {
		builder.PlugFunc(mgen(strconv.Itoa(i)))
	}

	builder.Build().ServeHTTP(nil, nil)
}

func BenchmarkPlugFunc(b *testing.B) {
	builder := NewBuilder()

	for i := 0; i < 10; i++ {
		builder.PlugFunc(mgen(""))
	}

	h := builder.Build()

	for i := 0; i < b.N; i++ {
		h.ServeHTTP(nil, nil)
	}
}

func BenchmarkPlug(b *testing.B) {
	builder := NewBuilder()

	for i := 0; i < 10; i++ {
		builder.Plug(&P{})
	}

	h := builder.Build()

	for i := 0; i < b.N; i++ {
		h.ServeHTTP(nil, nil)
	}
}
