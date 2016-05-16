package token

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
)

// Size specifies the Token size in bytes
const Size = 64

var encoding = base64.URLEncoding

// Token stores the user's session token
type Token struct {
	Token [Size]byte
}

// New creates a new Session with a random token
func New() (*Token, error) {
	s := &Token{}
	_, err := rand.Read(s.Token[:len(s.Token)])
	if err != nil {
		return nil, err
	}
	return s, nil
}

// MarshalJSON returns a URL-safe base64'd string encoded to JSON
func (s *Token) MarshalJSON() ([]byte, error) {
	b := make([]byte, Size)
	copy(b, s.Token[:len(s.Token)])
	return json.Marshal(encoding.EncodeToString(b))
}

// UnmarshalJSON decodes from a JSON encoded URL-safe base64 string
func (s *Token) UnmarshalJSON(b []byte) error {
	str := ""
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	b, err := encoding.DecodeString(str)
	if err != nil {
		return err
	}
	copy(s.Token[:len(s.Token)], b)
	return nil
}
