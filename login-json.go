package gosteam

import (
	cookiejar "github.com/mihakyr/go-jar"
	"github.com/valyala/fasthttp"
)

type Session struct {
	cookieClient *cookiejar.CookieJar
	client       *fasthttp.Client
	oauth        OAuth
	sessionID    string
	apiKey       string
	deviceID     string
	umqID        string
	chatMessage  int
	captchaGid   string
	captchaText  string
	language     string
}

type LoginResponse struct {
	Success      bool   `json:"success"`
	PublicKeyMod string `json:"publickey_mod"`
	PublicKeyExp string `json:"publickey_exp"`
	Timestamp    string
	TokenGID     string
}

type LoginSession struct {
	Success           bool   `json:"success"`
	LoginComplete     bool   `json:"login_complete"`
	RequiresTwoFactor bool   `json:"requires_twofactor"`
	EmailNeeded       bool   `json:"emailauth_needed"`
	CaptchaNeeded     bool   `json:"captcha_needed"`
	CaptchaGid        string `json:"captcha_gid"`
	Message           string `json:"message"`
	RedirectURI       string `json:"redirect_uri"`
	OAuthInfo         string `json:"oauth"`
}

type OAuth struct {
	SteamID       SteamID `json:"steamid,string"`
	Token         string  `json:"oauth_token"`
	WGToken       string  `json:"wgtoken"`
	WGTokenSecure string  `json:"wgtoken_secure"`
	WebCookie     string  `json:"webcookie"`
}
