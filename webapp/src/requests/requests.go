package requests

import (
	"io"
	"net/http"
)

func MakeARequestWithAuthentication(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {

}
