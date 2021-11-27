package gosteam

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/valyala/fasthttp"
	"io"
	"strings"
)

func (s *session) GetPurchaseHistory() ([]Purchase, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.SetRequestURI(steamPurchaseHistory)
	for k, v := range map[string]string{
		//"Accept": "text/html",
		//"Accept-Encoding": "gzip,deflate,br",
		"Accept-Language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7",
		"Cache-Control":   "no-cache",
		"Connection":      "keep-alive",
		//"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		"Origin":                    "https://store.steamcommunity.com",
		"Pragma":                    "no-cache",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "none",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; Google Nexus 4 - 4.1.1 - API 16 - 768x1280 Build/JRO03S) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
	} {
		req.Header.Set(k, v)
	}
	s.cookieClient.FillRequest(req)
	resp := fasthttp.AcquireResponse()

	if err := s.getRedirect(req, resp, 200, "GetPurchaseHistory"); err != nil {
		return nil, err
	}
	reader := &bytes.Buffer{}
	reader.Write(resp.Body())
	return parseHTML(reader), nil
}

func formatString(s string) string {
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	s = strings.TrimSpace(s)
	return s
}

type Purchase struct {
	Date  string
	Items string
	Type  string
	Total string
}

func parseHTML(r io.Reader) []Purchase {
	tokenizer, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		panic(err)
	}
	class := "wallet_table_row"
	var purchases = make([]Purchase, 0)
	tokenizer.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		if s.HasClass(class) {
			purchase := Purchase{}
			s.Find("td").Each(func(j int, td *goquery.Selection) {
				switch {
				case td.HasClass("wht_date"):
					purchase.Date = formatString(td.Text())
				case td.HasClass("wht_items"):
					purchase.Items = formatString(td.Text())
				case td.HasClass("wht_type"):
					purchase.Type = formatString(td.Text())
				case td.HasClass("wht_total"):
					purchase.Total = formatString(td.Text())
				default:
				}
			})
			purchases = append(purchases, purchase)
		}
	})
	return purchases
}
