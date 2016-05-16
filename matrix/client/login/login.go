package login

import (
	"encoding/json"
	"net/http"

	"github.com/cptaffe/yatm/errors"
	"github.com/cptaffe/yatm/token"
)

const (
	PasswordStage  string = "m.login.password"
	RecaptchaStage        = "m.login.recaptcha"
	Oauth2Stage           = "m.login.oath2"
	EmailIDStage          = "m.login.email.identity"
	TokenStage            = "m.login.token"
	DummyStage            = "m.login.dummy"
)

// Login is the config for login
type Login struct {
}

func New() *Login {
	return &Login{}
}

type Flow struct {
	Stage string `json:"type"`
}

type AuthResp struct {
	Flows []Flow `json:"flows"`
}

type Req struct {
	Type     string `json:"type"`
	User     string `json:"user"`
	Medium   string `json:"medium"`
	Addr     string `json:"address"`
	Password string `json:"password"`
}

type Resp struct {
	UserID       string       `json:"user_id"`
	HomeServer   string       `json:"home_server"`
	Token        *token.Token `json:"access_token"`
	RefreshToken *token.Token `json:"refresh_token"`
}

func (l *Login) Authenticate(req *Req, w http.ResponseWriter) {
	switch req.Type {
	case PasswordStage:
		// HACK: shim to login ficticious user cpt
		if req.User == "cpt" && req.Password == "bob" {
			tkn, err := token.New()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			rtkn, err := token.New()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err := json.NewEncoder(w).Encode(&Resp{
				UserID:       req.User,
				HomeServer:   "fuck.u",
				Token:        tkn,
				RefreshToken: rtkn,
			}); err != nil {
				// internal server error
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			// Credentials failed
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(&errors.Error{Code: errors.Forbidden})
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errors.Error{Code: errors.BadJSON})
	}
}

func (l *Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: rate limiting
	defer r.Body.Close()
	if r.Method == "POST" {
		req := &Req{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			// 401 response, bad json
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errors.New(errors.BadJSON, err))
			return
		}
		l.Authenticate(req, w)
	} else if r.Method == "GET" {
		if err := json.NewEncoder(w).Encode(&AuthResp{
			Flows: []Flow{
				Flow{
					Stage: PasswordStage,
				},
			},
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if r.Method == "OPTIONS" {
		// preflight header
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Allow", "GET, PUT, OPTIONS")
	}
}
