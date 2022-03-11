package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func ConfigCookie() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func SaveCookie(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encodedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func ReadCookie(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	if err = s.Decode("data", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}
