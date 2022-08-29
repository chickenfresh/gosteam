package gosteam

import (
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
)

func (s *Session) Logout() error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(steamLogout)
	req.SetBodyString(url.Values{"sessionid": {s.sessionID}}.Encode())
	resp := fasthttp.AcquireResponse()
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8,ru;q=0.6")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Host", "steamcommunity.com")
	req.Header.Set("Origin", "https://steamcommunity.com")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	if err := s.client.Do(req, resp); err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusFound {
		return errorStatusCode("Logout", resp.StatusCode())
	}

	return nil
}
