package gosteam

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	TradeStateNone = iota
	TradeStateInvalid
	TradeStateActive
	TradeStateAccepted
	TradeStateCountered
	TradeStateExpired
	TradeStateCanceled
	TradeStateDeclined
	TradeStateInvalidItems
	TradeStateCreatedNeedsConfirmation
	TradeStateCanceledByTwoFactor
	TradeStateInEscrow
)

const (
	TradeConfirmationNone = iota
	TradeConfirmationEmail
	TradeConfirmationMobileApp
	TradeConfirmationMobile
)

const (
	TradeFilterNone             = iota
	TradeFilterSentOffers       = 1 << 0
	TradeFilterRecvOffers       = 1 << 1
	TradeFilterActiveOnly       = 1 << 3
	TradeFilterHistoricalOnly   = 1 << 4
	TradeFilterItemDescriptions = 1 << 5
)

var (
	// receiptExp matches JSON in the following form:
	//	oItem = {"id":"...",...}; (Javascript code)
	receiptExp    = regexp.MustCompile("oItem =\\s(.+?});")
	myEscrowExp   = regexp.MustCompile("var g_daysMyEscrow = (\\d+);")
	themEscrowExp = regexp.MustCompile("var g_daysTheirEscrow = (\\d+);")
	errorMsgExp   = regexp.MustCompile("<div id=\"error_msg\">\\s*([^<]+)\\s*</div>")
	offerInfoExp  = regexp.MustCompile("token=([a-zA-Z0-9-_]+)")

	ErrReceiptMatch         = errors.New("unable to match items in trade receipt")
	ErrCannotAcceptActive   = errors.New("unable to accept a non-active trade")
	ErrCannotFindOfferInfo  = errors.New("unable to match data from trade offer url")
	ErrAccessDenied         = errors.New("access is denied")
	TradeOfferStateMappings = map[uint8]string{
		TradeStateNone:                     "none",
		TradeStateInvalid:                  "invalid",
		TradeStateActive:                   "active",
		TradeStateAccepted:                 "accepted",
		TradeStateCountered:                "countered",
		TradeStateExpired:                  "expired",
		TradeStateCanceled:                 "cancelled",
		TradeStateDeclined:                 "declined",
		TradeStateInvalidItems:             "invalid_items",
		TradeStateCreatedNeedsConfirmation: "created_needs_confirmation",
		TradeStateCanceledByTwoFactor:      "cancelled_by_two_factor",
		TradeStateInEscrow:                 "in_escrow",
	}
)

func (t *TradeOffer) Status() string {
	return TradeOfferStateMappings[t.State]
}

type GetTradeOffersOption struct {
	Key   string
	Value string
}

func GetSentOffers() GetTradeOffersOption {
	return GetTradeOffersOption{
		Key:   "get_sent_offers",
		Value: "1",
	}
}

func GetReceivedOffers() GetTradeOffersOption {
	return GetTradeOffersOption{
		Key:   "get_received_offers",
		Value: "1",
	}
}
func GetDescriptions() GetTradeOffersOption {
	return GetTradeOffersOption{
		Key:   "get_descriptions",
		Value: "1",
	}
}
func ActiveOnly() GetTradeOffersOption {
	return GetTradeOffersOption{
		Key:   "active_only",
		Value: "1",
	}
}

func HistoricalOnly() GetTradeOffersOption {
	return GetTradeOffersOption{
		Key:   "historical_only",
		Value: "1",
	}
}

func Language(value string) GetTradeOffersOption {
	return GetTradeOffersOption{
		Key:   "language",
		Value: value,
	}
}

func TimeHistoricalCutoff(value uint32) GetTradeOffersOption {
	return GetTradeOffersOption{
		Key:   "time_historical_cutoff",
		Value: fmt.Sprintf("%d", value),
	}
}

func (s *Session) GetTradeStatus(id int) ([]byte, error) {
	data := url.Values{
		"key":              {s.apiKey},
		"tradeid":          {strconv.Itoa(id)},
		"get_descriptions": {"0"},
	}

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.SetRequestURI(apiTradeStatus + data.Encode())
	s.cookieClient.FillRequest(req)
	resp := fasthttp.AcquireResponse()

	if err := s.doRequest(req, resp); err != nil {
		return nil, wrappedError("GetTradeStatus | s.doRequest(req, resp)", err)
	}
	//if resp.StatusCode() != 200 {
	//	return nil, errorStatusCode("GetTradeStatus", resp.StatusCode())
	//}
	//
	//var response APIResponse
	//if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&response); err != nil {
	//	return nil, errorText("GetTradeStatus | APIResponce - json.NewDecoder")
	//}
	//
	//return response.Inner.Offer, nil

	return resp.Body(), nil
}

func (s *Session) GetTradeOffer(id uint64) ([]byte, error) {
	data := url.Values{
		"key":          {s.apiKey},
		"tradeofferid": {strconv.FormatUint(id, 10)},
	}

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.SetRequestURI(apiGetTradeOffer + data.Encode())
	s.cookieClient.FillRequest(req)
	resp := fasthttp.AcquireResponse()

	if err := s.doRequest(req, resp); err != nil {
		return nil, wrappedError("GetTradeOffer | s.doRequest(req, resp)", err)
	}

	//if resp.StatusCode() != 200 {
	//	return nil, errorStatusCode("GetTradeOffer", resp.StatusCode())
	//}
	//
	//var response APIResponse
	//if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&response); err != nil {
	//	return nil, errorText("GetTradeOffer | APIResponse - json.NewDecoder")
	//}
	//
	//return response.Inner.Offer, nil
	return resp.Body(), nil
}

func (s *Session) GetTradeOffers(options ...GetTradeOffersOption) ([]byte, error) {
	data := url.Values{}

	if s.apiKey != "" {
		data.Set("key", s.apiKey)
	} else {
		return nil, errorApiKey
	}
	for _, option := range options {
		data.Set(option.Key, option.Value)
	}

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.SetRequestURI(apiGetTradeOffers + data.Encode())
	resp := fasthttp.AcquireResponse()

	if err := s.doRequest(req, resp); err != nil {
		return nil, wrappedError("GetTradeOffers | s.doRequest(req, resp)", err)
	}
	//if resp.StatusCode() == http.StatusForbidden {
	//	return nil, ErrAccessDenied
	//}
	//var response APIResponse
	//if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&response); err != nil {
	//	return nil, err
	//}
	//return response.Inner, nil
	return resp.Body(), nil
}

func (s *Session) GetMyTradeToken() (string, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.SetRequestURI(apiTradeOfferToken)
	s.cookieClient.FillRequest(req)
	resp := fasthttp.AcquireResponse()

	if err := s.getRedirect(req, resp, 200, "GetMyTradeToken"); err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(bytes.NewReader(resp.Body()))
	if err != nil {
		return "", err
	}

	m := offerInfoExp.FindStringSubmatch(string(body))
	if m == nil || len(m) != 2 {
		return "", ErrCannotFindOfferInfo
	}

	return m[1], nil
}

func (s *Session) SendTradeOffer(tradeOfferJSON, message, tradeLink, steamId string) ([]byte, error) {
	// https://steamcommunity.com/tradeoffer/new/?partner=472082445&token=wb7wG5LI
	token := strings.Split(tradeLink, "token=")[1]

	req := fasthttp.AcquireRequest()

	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(apiSendTradeOffer)
	req.Header.SetReferer(tradeLink)
	s.enrichWithDefaultHeaders(req)
	req.SetBodyString(url.Values{
		"sessionid":                 {s.sessionID},
		"serverid":                  {"1"},
		"partner":                   {steamId},
		"tradeoffermessage":         {message},
		"json_tradeoffer":           {tradeOfferJSON},
		"trade_offer_create_params": {"{\"trade_offer_access_token\":\"" + token + "\"}"},
	}.Encode())

	resp := fasthttp.AcquireResponse()

	if err := s.doRequest(req, resp); err != nil {
		return nil, wrappedError("SendTradeOffer | s.doRequest(req, resp)", err)
	}

	return resp.Body(), nil

	//if resp.StatusCode() != 200 {
	//	return nil, errorStatusCode("SendTradeOffer", resp.StatusCode())
	//}
	//
	//var response SendTradeOfferResponse
	//if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&response); err != nil {
	//	return nil, errorText("SendTradeOffer | SendTradeOfferResponse - json.NewDecoder")
	//}
	//
	//if len(response.ErrorMessage) != 0 {
	//	return nil, errors.New(response.ErrorMessage)
	//}
	//
	//if response.ID == 0 {
	//	return nil, errors.New("no OfferID included")
	//}
	//
	//return &response, nil
}

func (s *Session) DeclineTradeOffer(id uint64) error {
	data := url.Values{
		"key":          {s.apiKey},
		"tradeofferid": {strconv.FormatUint(id, 10)},
	}.Encode()

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(apiDeclineTradeOffer)
	req.SetBodyString(data)
	resp := fasthttp.AcquireResponse()

	if err := s.client.Do(req, resp); err != nil {
		return errorText("fasthttp.Do(req, resp) | DeclineTradeOffer")
	}
	if resp.StatusCode() != 200 {
		return errorStatusCode("DeclineTradeOffer", resp.StatusCode())
	}

	result := string(resp.Header.Peek("x-eresult"))
	if result != "1" {
		return fmt.Errorf("cannot decline trade: %s", result)
	}

	return nil
}

func (s *Session) CancelTradeOffer(id uint64) ([]byte, error) {
	req := fasthttp.AcquireRequest()

	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(fmt.Sprintf(apiCancelTradeOfferTempl, id))
	s.enrichWithDefaultHeaders(req)
	req.SetBodyString(url.Values{
		"sessionid": {s.sessionID},
	}.Encode())

	resp := fasthttp.AcquireResponse()

	if err := s.doRequest(req, resp); err != nil {
		return nil, wrappedError("CancelTradeOffer | doRequest(req, resp)", err)
	}

	return resp.Body(), nil
}
