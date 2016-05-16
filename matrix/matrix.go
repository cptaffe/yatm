package matrix

import (
	"net/http"

	"github.com/cptaffe/yatm/matrix/client"
)

type Matrix struct {
	Mux    *http.ServeMux
	Client *client.Client
}

func (m *Matrix) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Mux.ServeHTTP(w, r)
}

func New(c *client.Client) *Matrix {
	mux := http.NewServeMux()
	mux.Handle("/_matrix/client/", c)
	return &Matrix{
		Mux:    mux,
		Client: c,
	}
}
