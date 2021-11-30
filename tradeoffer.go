package gosteam

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
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

	ErrReceiptMatch        = errors.New("unable to match items in trade receipt")
	ErrCannotAcceptActive  = errors.New("unable to accept a non-active trade")
	ErrCannotFindOfferInfo = errors.New("unable to match data from trade offer url")
	ErrAccessDenied        = errors.New("access is denied")
)

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

func (s *session) GetTradeOffer(id uint64) (*TradeOffer, error) {
	data := url.Values{
		"key":          {s.apiKey},
		"tradeofferid": {strconv.FormatUint(id, 10)},
	}

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.SetRequestURI(apiGetTradeOffer + data.Encode())
	s.cookieClient.FillRequest(req)
	resp := fasthttp.AcquireResponse()

	if err := s.client.Do(req, resp); err != nil {
		return nil, errorText("GetTradeOffer | s.client.Do(req, resp)")
	}
	if resp.StatusCode() != 200 {
		return nil, errorStatusCode("GetTradeOffer", resp.StatusCode())
	}

	var response APIResponse
	if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&response); err != nil {
		return nil, errorText("GetTradeOffer | APIResponce - json.NewDecoder")
	}

	return response.Inner.Offer, nil
}

func (s *session) GetTradeOffers(options ...GetTradeOffersOption) (*TradeOfferResponse, error) {
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

	if err := s.client.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() == http.StatusForbidden {
		return nil, ErrAccessDenied
	}
	var response APIResponse
	if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&response); err != nil {
		return nil, err
	}
	return response.Inner, nil
}

func (s *session) GetMyTradeToken() (string, error) {
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

type EscrowSteamGuardInfo struct {
	MyDays   int64
	ThemDays int64
	ErrorMsg string
}

//func (session *Session) GetEscrowGuardInfo(sid SteamID, token string) (*EscrowSteamGuardInfo, error) {
//	return session.GetEscrow("https://steamcommunity.com/tradeoffer/new/?" + url.Values{
//		"partner": {strconv.FormatUint(uint64(sid.GetAccountID()), 10)},
//		"token":   {token},
//	}.Encode())
//}
//
//func (session *Session) GetEscrowGuardInfoForTrade(offerID uint64) (*EscrowSteamGuardInfo, error) {
//	return session.GetEscrow("https://steamcommunity.com/tradeoffer/" + strconv.FormatUint(offerID, 10))
//}
//
//func (session *Session) GetEscrow(url string) (*EscrowSteamGuardInfo, error) {
//	resp, err := session.client.Get(url)
//	if resp != nil {
//		defer resp.Body.Close()
//	}
//
//	if err != nil {
//		return nil, err
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("http error: %d", resp.StatusCode)
//	}
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	var my int64
//	var them int64
//	var errMsg string
//
//	m := myEscrowExp.FindStringSubmatch(string(body))
//	if len(m) == 2 {
//		my, _ = strconv.ParseInt(m[1], 10, 32)
//	}
//
//	m = themEscrowExp.FindStringSubmatch(string(body))
//	if len(m) == 2 {
//		them, _ = strconv.ParseInt(m[1], 10, 32)
//	}
//
//	m = errorMsgExp.FindStringSubmatch(string(body))
//	if len(m) == 2 {
//		errMsg = m[1]
//	}
//
//	return &EscrowSteamGuardInfo{
//		MyDays:   my,
//		ThemDays: them,
//		ErrorMsg: errMsg,
//	}, nil
//}
//

func (s *session) SendTradeOffer(offer *TradeOffer, sid SteamID, token string) (*SendTradeOfferResponse, error) {
	content := map[string]interface{}{
		"newversion": true,
		"version":    3,
		"me": map[string]interface{}{
			"assets":   offer.SendItems,
			"currency": make([]struct{}, 0),
			"ready":    false,
		},
		"them": map[string]interface{}{
			"assets":   offer.RecvItems,
			"currency": make([]struct{}, 0),
			"ready":    false,
		},
	}

	contentJSON, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(apiSendTradeOffer)
	req.SetBodyString(url.Values{
		"sessionid":                 {s.sessionID},
		"serverid":                  {"1"},
		"partner":                   {sid.ToString()},
		"tradeoffermessage":         {offer.Message},
		"json_tradeoffer":           {string(contentJSON)},
		"trade_offer_create_params": {"{\"trade_offer_access_token\":\"" + token + "\"}"},
	}.Encode())
	resp := fasthttp.AcquireResponse()

	req.Header.Add("Referer", "https://steamcommunity.com/tradeoffer/new/?"+url.Values{
		"partner": {strconv.FormatUint(uint64(sid.GetAccountID()), 10)},
		"token":   {token},
	}.Encode())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	if err := s.client.Do(req, resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errorStatusCode("SendTradeOffer", resp.StatusCode())
	}

	var response SendTradeOfferResponse
	if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&response); err != nil {
		return nil, errorText("SendTradeOffer | SendTradeOfferResponse - json.NewDecoder")
	}

	if len(response.ErrorMessage) != 0 {
		return nil, errors.New(response.ErrorMessage)
	}

	if response.ID == 0 {
		return nil, errors.New("no OfferID included")
	}

	//offer.ID = response.ID
	//offer.Created = time.Now().Unix()
	//offer.Updated = time.Now().Unix()
	//offer.Expires = offer.Created + 14*24*60*60
	//offer.RealTime = false
	//offer.IsOurOffer = true
	//
	//Just test mobile confirmation, email is deprecated
	//if response.MobileConfirmationRequired {
	//	offer.ConfirmationMethod = TradeConfirmationMobileApp
	//	offer.State = TradeStateCreatedNeedsConfirmation
	//} else {
	//	set state to active
	//offer.State = TradeStateActive
	//}

	return &response, nil
}

//
//
//func (session *Session) GetTradeReceivedItems(receiptID uint64) ([]*InventoryItem, error) {
//	resp, err := session.client.Get(fmt.Sprintf("https://steamcommunity.com/trade/%d/receipt", receiptID))
//	if resp != nil {
//		defer resp.Body.Close()
//	}
//
//	if err != nil {
//		return nil, err
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("http error: %d", resp.StatusCode)
//	}
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	m := receiptExp.FindAllSubmatch(body, -1)
//	if m == nil {
//		return nil, ErrReceiptMatch
//	}
//
//	items := make([]*InventoryItem, len(m))
//	for k := range m {
//		item := &InventoryItem{}
//		if err = json.Unmarshal(m[k][1], item); err != nil {
//			return nil, err
//		}
//
//		items[k] = item
//	}
//
//	return items, nil
//}
//
func (s *session) DeclineTradeOffer(id uint64) error {
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

func (s *session) CancelTradeOffer(id uint64) error {
	data := url.Values{
		"key":          {s.apiKey},
		"tradeofferid": {strconv.FormatUint(id, 10)},
	}.Encode()

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(apiCancelTradeOffer)
	req.SetBodyString(data)
	resp := fasthttp.AcquireResponse()

	if err := s.client.Do(req, resp); err != nil {
		return errorText("fasthttp.Do(req, resp) | CancelTradeOffer")
	}
	if resp.StatusCode() != 200 {
		return errorStatusCode("CancelTradeOffer", resp.StatusCode())
	}

	result := string(resp.Header.Peek("x-eresult"))
	if result != "1" {
		return fmt.Errorf("cannot cancel trade: %s", result)
	}

	return nil
}

//
//func (session *Session) AcceptTradeOffer(id uint64) error {
//	tid := strconv.FormatUint(id, 10)
//	postURL := "https://steamcommunity.com/tradeoffer/" + tid
//
//	req, err := http.NewRequest(
//		http.MethodPost,
//		postURL+"/accept",
//		strings.NewReader(url.Values{
//			"sessionid":    {session.sessionID},
//			"serverid":     {"1"},
//			"tradeofferid": {tid},
//		}.Encode()),
//	)
//	if err != nil {
//		return err
//	}
//
//	req.Header.Add("Referer", postURL)
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//
//	resp, err := session.client.Do(req)
//	if resp != nil {
//		defer resp.Body.Close()
//	}
//
//	if err != nil {
//		return err
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return fmt.Errorf("http error: %d", resp.StatusCode)
//	}
//
//	type Response struct {
//		ErrorMessage string `json:"strError"`
//	}
//
//	var response Response
//	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
//		return err
//	}
//
//	if len(response.ErrorMessage) != 0 {
//		return errors.New(response.ErrorMessage)
//	}
//
//	return nil
//}
//
//func (offer *TradeOffer) Send(session *Session, sid SteamID, token string) error {
//	return session.SendTradeOffer(offer, sid, token)
//}
//
//func (offer *TradeOffer) Accept(session *Session) error {
//	return session.AcceptTradeOffer(offer.ID)
//}
//
//func (offer *TradeOffer) Cancel(session *Session) error {
//	if offer.IsOurOffer {
//		return session.CancelTradeOffer(offer.ID)
//	}
//
//	return session.DeclineTradeOffer(offer.ID)
//}
