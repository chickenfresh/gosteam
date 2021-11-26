package gosteam

import (
	"github.com/valyala/fasthttp"
	"net/url"
)

func (s *session) Logout() ([]byte, error) {
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
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errorStatusCode("SendTradeOffer", resp.StatusCode())
	}

	return resp.Body(), nil
}
