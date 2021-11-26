package gosteam

const (
	steamDefault   = "https://steamcommunity.com"
	steamLogin     = steamDefault + "/login"
	steamLogout    = steamDefault + "/login/logout/"
	steamGetRSAkey = steamDefault + "/login/getrsakey/"
	steamDoLogin   = steamDefault + "/login/dologin/"
	steamSetAvatar = steamDefault + "/actions/FileUploader"

	// tradeoffer
	apiSendTradeOffer    = steamDefault + "/tradeoffer/new/send"
	apiGetTradeOffer     = "https://api.steampowered.com/IEconService/GetTradeOffer/v1/?"
	apiGetTradeOffers    = "https://api.steampowered.com/IEconService/GetTradeOffers/v1/?"
	apiDeclineTradeOffer = "https://api.steampowered.com/IEconService/DeclineTradeOffer/v1/"
	apiCancelTradeOffer  = "https://api.steampowered.com/IEconService/CancelTradeOffer/v1/"
	apiTradeOfferToken   = steamDefault + "/my/tradeoffers/privacy"
)
