package version

import (
	"encoding/json"
	"net/http"
)

type Version struct {
	Versions []string
}

func New(v ...string) *Version {
	return &Version{
		Versions: v,
	}
}

type Resp struct {
	Versions []string `json:"versions"`
}

func (v *Version) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "GET" {
		if err := json.NewEncoder(w).Encode(&Resp{
			Versions: v.Versions,
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Allow", "GET")
	}
}
