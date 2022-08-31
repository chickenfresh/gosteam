package gosteam

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
)

func (s *Session) doRequest(req *fasthttp.Request, resp *fasthttp.Response) error {
	if err := s.client.Do(req, resp); err != nil {
		return fmt.Errorf("steam request %s %s failed: %w",
			string(req.Header.Method()), req.URI().String(), err)
	}

	if logger != nil {
		logger.Info("steam request: %s | steam response: %d", req.URI(), resp.StatusCode())
	}

	switch resp.StatusCode() {
	case fasthttp.StatusUnauthorized:
		return ErrUnauthorized
	case fasthttp.StatusFound:
		l := resp.Header.Peek("Location")
		if len(l) > 0 && bytes.Contains(l, []byte("steammobile://lostauth")) {
			return ErrUnauthorized
		}
	case fasthttp.StatusForbidden:
		return ErrUnauthorized
	}

	s.takeCookie(resp)

	return nil
}

func (s *Session) takeCookie(resp *fasthttp.Response) {
	resp.Header.VisitAllCookie(func(key, value []byte) {
		processingcookie := bytes.Split(value, []byte(";"))
		cookie := bytes.Split(processingcookie[0], []byte("="))
		s.cookieClient.SetBytesKV(cookie[0], cookie[1])
	})
}

func (s *Session) enrichWithDefaultHeaders(req *fasthttp.Request) {
	req.Header.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36")
	req.Header.Set("Origin", steamDefault)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.8,en-US;q=0.6,en;q=0.4")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	s.cookieClient.FillRequest(req)
}

func plainTextBody(resp *fasthttp.Response) []byte {
	ce := resp.Header.Peek("Content-Encoding")
	if len(ce) > 0 && bytes.Contains(ce, []byte("gzip")) {
		r, err := ungzip(resp.Body())
		if err != nil {
			if logger != nil {
				logger.Error("have to return gziped body: unzip error: %v", err)
			}
			return resp.Body()
		}
		return r
	}

	return resp.Body()
}

func ungzip(origin []byte) ([]byte, error) {
	reader := bytes.NewReader(origin)
	gzreader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	output, err := io.ReadAll(gzreader)
	if err != nil {
		return nil, err
	}

	return output, nil
}
