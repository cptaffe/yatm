package client

import (
	"net/http"

	"github.com/cptaffe/yatm/matrix/client/login"
	"github.com/cptaffe/yatm/matrix/client/version"
)

type Client struct {
	Login *login.Login
	Mux   *http.ServeMux
}

func New(l *login.Login) *Client {
	mux := http.NewServeMux()
	mux.Handle("/_matrix/client/r0/login", l)
	mux.Handle("/_matrix/client/versions", version.New("r0.1.0"))
	return &Client{
		Mux: mux,
	}
}

func (c *Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// CORS headers
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	c.Mux.ServeHTTP(w, r)
}
