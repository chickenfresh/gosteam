package gosteam

const (
	steamDefault    = "https://steamcommunity.com"
	steamAPIDefault = "https://api.steampowered.com/IEconService"
	steamLogin      = steamDefault + "/login"
	steamLogout     = steamDefault + "/login/logout/"
	steamGetRSAkey  = steamDefault + "/login/getrsakey/"
	steamDoLogin    = steamDefault + "/login/dologin/"
	steamSetAvatar  = steamDefault + "/actions/FileUploader/"

	steamPurchaseHistory = "https://store.steampowered.com/account/history/"

	// tradeoffer
	apiSendTradeOffer  = steamDefault + "/tradeoffer/new/send"
	apiTradeOfferToken = steamDefault + "/my/tradeoffers/privacy"

	apiGetTradeOffer     = steamAPIDefault + "/GetTradeOffer/v1/?"
	apiGetTradeOffers    = steamAPIDefault + "/GetTradeOffers/v1/?"
	apiDeclineTradeOffer = steamAPIDefault + "/DeclineTradeOffer/v1/"
	apiCancelTradeOffer  = steamAPIDefault + "/CancelTradeOffer/v1/?"
)
