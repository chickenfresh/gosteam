package gosteam

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
)

func (s *Session) doRequest(req *fasthttp.Request, resp *fasthttp.Response) error {
	if err := s.client.Do(req, resp); err != nil {
		return wrappedError("steam request fail", err)
	}
	if logger != nil {
		logger.Info("steam request: %s | steam response: %d", req.URI(), resp.StatusCode())
		logger.Debug("STEAM RQ: %+v", req)
		logger.Debug("STEAM RS: %+v", resp)
	}
	s.takeCookie(resp)

	return nil
}

func (s *Session) getRedirect(req *fasthttp.Request, resp *fasthttp.Response, statuscode int, funcname string) error {
	for {
		if err := s.client.Do(req, resp); err != nil {
			logger.Error("Error: %v\n", err)
			return errorText(fmt.Sprintf("request error %s", funcname))
		}

		if logger != nil {
			logger.Info("steam request: %s | steam response: %d", req.URI(), resp.StatusCode())
			logger.Debug("STEAM RQ: %v", req)
			logger.Debug("STEAM RS: %v", resp)
		}

		s.takeCookie(resp)

		switch resp.StatusCode() {
		case statuscode:
			return nil
		case 200:
			return nil
		case 302:
			location := resp.Header.Peek("Location")
			req.SetRequestURIBytes(location)
			s.cookieClient.FillRequest(req)
			continue
		default:
			return errorStatusCode(funcname, resp.StatusCode())
		}
	}
}

func (s *Session) takeCookie(resp *fasthttp.Response) {
	resp.Header.VisitAllCookie(func(key, value []byte) {
		processingcookie := bytes.Split(value, []byte(";"))
		cookie := bytes.Split(processingcookie[0], []byte("="))
		s.cookieClient.SetBytesKV(cookie[0], cookie[1])
	})
}
