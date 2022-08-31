package gosteam

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
)

const (
	PrivacyStatePrivate     = 1
	PrivacyStateFriendsOnly = 2
	PrivacyStatePublic      = 3
)

const (
	CommentSettingSelf    = "commentselfonly"
	CommentSettingFriends = "commentfriendsonly"
	CommentSettingPublic  = "commentanyone"
)

const (
	apiGetPlayerSummaries = "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?"
	apiGetOwnedGames      = "https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?"
	apiGetPlayerBans      = "https://api.steampowered.com/ISteamUser/GetPlayerBans/v1/?"
	apiGetPlayerFriends   = "https://api.steampowered.com/ISteamUser/GetFriendList/v1/?"
	apiResolveVanityURL   = "https://api.steampowered.com/ISteamUser/ResolveVanityURL/v1/?"
)

var ErrCannotFindVanityMatch = errors.New("no match for the vanity URL")

type PlayerSummary struct {
	SteamID           SteamID `json:"steamid,string"`
	VisibilityState   uint32  `json:"communityvisibilitystate"`
	ProfileState      uint32  `json:"profilestate"`
	PersonaName       string  `json:"personaname"`
	PersonaState      uint32  `json:"personastate"`
	PersonaStateFlags uint32  `json:"personastateflags"`
	RealName          string  `json:"realname"`
	LastLogoff        int64   `json:"lastlogoff"`
	ProfileURL        string  `json:"profileurl"`
	AvatarURL         string  `json:"avatar"`
	AvatarMediumURL   string  `json:"avatarmedium"`
	AvatarFullURL     string  `json:"avatarfull"`
	PrimaryClanID     uint64  `json:"primaryclanid,string"`
	TimeCreated       int64   `json:"timecreated"`
	LocCountryCode    string  `json:"loccountrycode"`
	LocStateCode      string  `json:"locstatecode"`
	LocCityID         uint32  `json:"loccityid"`
	GameID            uint64  `json:"gameid,string"`
	GameServerIP      string  `json:"gameserverip"`
	GameExtraInfo     string  `json:"gameextrainfo"`
}

type Game struct {
	AppID           uint32 `json:"appid"`
	PlaytimeForever int64  `json:"playtime_forever"`
	Playtime2Weeks  int64  `json:"playtime_2weeks"`
}

type OwnedGamesResponse struct {
	Count uint32  `json:"game_count"`
	Games []*Game `json:"games"`
}

type PlayerBan struct {
	SteamID          uint64 `json:"SteamId,string"`
	CommunityBanned  bool   `json:"CommunityBanned"`
	VACBanned        bool   `json:"VACBanned"`
	NumberOfVACBans  int    `json:"NumberOfVACBans"`
	DaysSinceLastBan int    `json:"DaysSinceLastBan"`
	NumberOfGameBans int    `json:"NumberOfGameBans"`
	EconomyBan       string `json:"EconomyBan"`
}

type Friend struct {
	SteamID      uint64 `json:"steamid,string"`
	Relationship string `json:"relationship"`
	FriendSince  int64  `json:"friend_since"`
}

func (s *Session) GetProfileURL() (string, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://steamcommunity.com/my")
	s.cookieClient.FillRequest(req)
	resp := fasthttp.AcquireResponse()

	if err := s.doRequest(req, resp); err != nil {
		return "", err
	}

	return string(resp.Header.Peek("Location")), nil
}

type ProfileDataOption struct {
	Key   string
	Value string
}

func ProfileName(name string) ProfileDataOption {
	return ProfileDataOption{
		Key:   "personaName",
		Value: name,
	}
}

func ProfileCustomURL(customURL string) ProfileDataOption {
	return ProfileDataOption{
		Key:   "customURL",
		Value: customURL,
	}
}

func ProfileCountry(country string) ProfileDataOption {
	return ProfileDataOption{
		Key:   "country",
		Value: country,
	}
}

func ProfileSummary(summary string) ProfileDataOption {
	return ProfileDataOption{
		Key:   "summary",
		Value: summary,
	}
}

var NoDataToUpdate = errors.New("no data for update")

func (s *Session) SetProfileInfo(options ...ProfileDataOption) error {
	if len(options) == 0 {
		return NoDataToUpdate
	}
	data := url.Values{
		"sessionID": {s.sessionID},
		"type":      {"profileSave"},
		"json":      {"1"},
	}
	for _, option := range options {
		data.Set(option.Key, option.Value)
	}
	urlString := fmt.Sprintf("https://steamcommunity.com/profiles/%s/edit", s.oauth.SteamID.ToString())
	req := fasthttp.AcquireRequest()
	s.cookieClient.FillRequest(req)
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(urlString)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://steamcommunity.com")
	req.Header.Set("Referer", urlString+"/info")
	req.SetBodyString(data.Encode())
	resp := fasthttp.AcquireResponse()

	if err := s.client.Do(req, resp); err != nil {
		return errorText("fasthttp.Do(req, resp) | SetProfileInfo")
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("http error: %d", resp.StatusCode())
	}
	return nil
}
