package application

import "time"

const (
	MediaPath           = "media/"
	TemplatePath        = MediaPath + "template/"
	StaticPath          = MediaPath + "static/"
	AuthCookieExpire    = time.Hour * 24 * 30
	DefaultObjectsLimit = 12
	BasketCookie        = "basket"
	NextUrlCookie       = "next_url"
	AppCookieName       = "BAB"
	SecureKey           = "This should be change %Secure%"
	SMSAPIKey           = "6A34415A794E6D456C357047746856397335435353513D3D"
)
