package gosteam

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/valyala/fasthttp"
	"net"
	"strings"
)

func proxyHttpDialer(proxy string) fasthttp.DialFunc {
	var auth string
	if strings.Contains(proxy, "@") {
		split := strings.Split(proxy, "@")
		auth = base64.StdEncoding.EncodeToString([]byte(split[0]))
		proxy = split[1]
	}

	return func(addr string) (net.Conn, error) {
		conn, err := fasthttp.Dial(proxy)
		if err != nil {
			return nil, err
		}

		req := "CONNECT " + addr + " HTTP/1.1\r\n"
		if auth != "" {
			req += "Proxy-Authorization: Basic " + auth + "\r\n"
		}
		req += "\r\n"

		if _, err := conn.Write([]byte(req)); err != nil {
			return nil, err
		}

		res := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(res)

		res.SkipBody = true

		if err := res.Read(bufio.NewReader(conn)); err != nil {
			conn.Close()
			return nil, err
		}
		if res.Header.StatusCode() != 200 {
			conn.Close()
			return nil, fmt.Errorf("could not connect to proxy")
		}
		return conn, nil
	}
}

func (s *Session) SetProxy(address string) {
	// If it has schema
	if strings.Contains(address, "//") {
		address = strings.Split(address, "//")[1]
	}
	s.client.Dial = proxyHttpDialer(address)
}
