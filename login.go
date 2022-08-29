package gosteam

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mihakyr/go-jar"
	"github.com/valyala/fasthttp"
	"math/big"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	badCredentials = "The account name or password that you have entered is incorrect"
)

func (s *Session) Login(accountName, password, sharedSecret string) error {
	err := s.generateLoginCookie()
	if err != nil {
		return err
	}

	response, err := s.generateRSAkey(accountName)
	if err != nil {
		return err
	}

	twoFactorCode, err := GenerateTwoFactorCode(sharedSecret, time.Now().Unix())
	if err != nil {
		return err
	}

	return s.loginInAccount(response, accountName, password, twoFactorCode)
}

func (s *Session) generateRSAkey(username string) (*LoginResponse, error) {
	data := url.Values{
		"username":   {username},
		"donotcache": {strconv.FormatInt(time.Now().Unix()*1000, 10)},
	}.Encode()

	req := fasthttp.AcquireRequest()

	req.SetBodyString(data)
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(steamGetRSAkey)
	req.Header.SetUserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36")
	req.Header.SetContentLength(len(req.Body()))
	req.Header.SetContentType("application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.SetReferer(steamDefault)
	req.Header.Set("Origin", steamLogin)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Accept", "*/*")
	s.cookieClient.FillRequest(req)

	resp := fasthttp.AcquireResponse()

	if err := s.getRedirect(req, resp, 200, "generateRSAkey"); err != nil {
		return nil, err
	}

	var response LoginResponse
	if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&response); err != nil {
		return nil, wrappedError("generateRSAkey | LoginResponse-json.NewDecoder", err)
	}

	if !response.Success {
		return nil, ErrRSA
	}

	return &response, nil
}

// start login, setup cookie
func (s *Session) generateLoginCookie() error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.Header.SetRequestURI(steamLogin)
	resp := fasthttp.AcquireResponse()

	// TODO: with retry && error handling
	if err := s.getRedirect(req, resp, 200, "generateLoginCookie"); err != nil {
		return err
	}

	_, timezoneOffset := time.Now().Zone()

	s.cookieClient.Set("timezoneOffset", fmt.Sprintf("%d,0", timezoneOffset))
	s.cookieClient.Set("mobileClient", "android")
	s.cookieClient.Set("mobileClientVersion", "0 (2.1.3)")
	s.cookieClient.Set("Steam_Language", "english")

	return nil
}

func (s *Session) loginInAccount(response *LoginResponse, accountName, password, twoFactorCode string) error {
	var n big.Int
	n.SetString(response.PublicKeyMod, 16)

	exp, err := strconv.ParseInt(response.PublicKeyExp, 16, 32)
	if err != nil {
		return errorText("loginInAccount | strconv.ParseInt")
	}

	pub := rsa.PublicKey{N: &n, E: int(exp)}
	rsaOut, err := rsa.EncryptPKCS1v15(rand.Reader, &pub, []byte(password))
	if err != nil {
		return errorText("loginInAccount | rsa.EncryptPKCS1v15")
	}

	req := fasthttp.AcquireRequest()

	req.SetRequestURI(steamDoLogin)
	req.Header.SetMethod("POST")
	req.Header.SetUserAgent("Mozilla/5.0")
	req.Header.SetContentType("application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.SetReferer(steamDefault + "/login?oauth_client_id=DE45CD61&oauth_scope=read_profile%20write_profile%20read_client%20write_client")
	req.Header.Set("Accept", "text/javascript, text/html, application/xml, text/xml, */*")
	req.Header.Set("X-Requested-With", "com.valvesoftware.android.gosteam.community")
	req.Header.Set("Origin", steamDefault)
	s.cookieClient.FillRequest(req)
	req.SetBodyString(url.Values{
		"captcha_text":      {s.captchaText},
		"captchagid":        {s.captchaGid},
		"emailauth":         {""},
		"emailsteamid":      {""},
		"username":          {accountName},
		"password":          {base64.StdEncoding.EncodeToString(rsaOut)},
		"remember_login":    {"false"},
		"rsatimestamp":      {response.Timestamp},
		"twofactorcode":     {twoFactorCode},
		"donotcache":        {strconv.FormatInt(time.Now().Unix()*1000, 10)},
		"loginfriendlyname": {""},
		"oauth_client_id":   {"DE45CD61"},
		"oauth_scope":       {"read_profile write_profile read_client write_client"},
	}.Encode())

	resp := fasthttp.AcquireResponse()

	if err = s.getRedirect(req, resp, 200, "loginInAccount"); err != nil {
		return err
	}

	var loginSession LoginSession
	if err = json.NewDecoder(bytes.NewBuffer(resp.Body())).Decode(&loginSession); err != nil {
		return errorText("loginInAccount | loginSession-json.NewDecoder")
	}
	if !loginSession.Success {
		if logger != nil {
			logger.Error("login failed: %s", loginSession.Message)
		}
		switch {
		case strings.Contains(loginSession.Message, badCredentials):
			return ErrBadCredentials
		case loginSession.CaptchaNeeded:
			return newCaptchaNeededError(loginSession.CaptchaGid)
		case loginSession.EmailNeeded:
			return ErrNoAuthenticator
		case !loginSession.LoginComplete:
			return ErrBadCredentials
		case loginSession.RequiresTwoFactor:
			return ErrNeedTwoFactor
		}
		return errorText("dologin unknown error")
	}

	var oauthSession OAuth
	if err = json.NewDecoder(bytes.NewBufferString(loginSession.OAuthInfo)).Decode(&oauthSession); err != nil {
		return errorText("loginInAccount | oauthSession-json.NewDecoder")
	}

	randomBytes := make([]byte, 6)
	if _, err = rand.Read(randomBytes); err != nil {
		return errorText("loginInAccount | rand.Read(randomBytes)")
	}

	sessionID := make([]byte, hex.EncodedLen(len(randomBytes)))
	hex.Encode(sessionID, randomBytes)

	// Why Del?
	for name := range *s.cookieClient {
		if name == "mobileClient" || name == "mobileClientVersion" || name == "steamCountry" || strings.Contains(name, "steamMachineAuth") {
			s.cookieClient.Del(name)
		}
	}

	sum := md5.Sum([]byte(accountName + password))
	s.oauth = oauthSession
	s.deviceID = fmt.Sprintf("android:%x-%x-%x-%x-%x", sum[:2], sum[2:4], sum[4:6], sum[6:8], sum[8:10])
	s.sessionID = string(sessionID)
	s.cookieClient.Set("sessionid", s.sessionID)

	return nil
}

func NewSessionWithAPIKey(apiKey string) *Session {
	return &Session{
		cookieClient: &cookiejar.CookieJar{},
		client:       &fasthttp.Client{},
		apiKey:       apiKey,
		language:     "english",
		captchaGid:   "-1",
	}
}

func FromCache(apiKey, deviceId, sessionId string, oauth OAuth, cookies map[string]string) *Session {
	cj := &cookiejar.CookieJar{}
	for k, v := range cookies {
		cj.Set(k, v)
	}
	return &Session{
		cookieClient: cj,
		client:       &fasthttp.Client{},
		apiKey:       apiKey,
		oauth:        oauth,
		deviceID:     deviceId,
		sessionID:    sessionId,
	}
}

func (s *Session) OAuth() OAuth {
	return s.oauth
}

func (s *Session) DeviceId() string {
	return s.deviceID
}

func (s *Session) SessionId() string {
	return s.sessionID
}

func (s *Session) Cookies() map[string]string {
	cookies := make(map[string]string)
	if s.cookieClient != nil {
		for _, cookie := range *s.cookieClient {
			cookies[string(cookie.Key())] = string(cookie.Value())
		}
	}
	return cookies
}

func (s *Session) SetCaptchaText(text string) {
	s.captchaText = text
}

func (s *Session) SetCaptchaGid(gid string) {
	s.captchaGid = gid
}
