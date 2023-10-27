package enkanetworkapigo

import (
	"os"

	"github.com/Fesaa/enka-network-api-go/cache"
	"github.com/Fesaa/enka-network-api-go/localization"
	"github.com/withmandala/go-log"
)

const BASE_URL = "https://enka.network/api/"
const BASE_GENSHIN_UI_URL = "https://enka.network/ui/"
const BASE_SR_UI_URL = "https://enka.network/ui/hsr/"

type EnkaNetworkAPI struct {
	userAgent string

	log *log.Logger
}

// New creates a new EnkaNetworkAPI instance
//
// # Will also initialize the cache and localization
//
// From the documentation:
//
//	Please set a custom User-Agent header with your requests so I can track them better and help you if needed.
//
// See https://api.enka.network/ for API docs
func New(userAgent string) *EnkaNetworkAPI {
	localization.Init()
	cache.Init(cache.MEMORY)
	return &EnkaNetworkAPI{
		userAgent: userAgent,

		log: log.New(os.Stdout).WithColor(),
	}
}

// NewDefaultUserAgent creates a new EnkaNetworkAPI instance with the default User-Agent header
//
// # This is not recommended, as it does not help the developer track requests
//
// Consider using New or SetUserAgent instead
func NewDefaultUserAgent() *EnkaNetworkAPI {
	return New("enka-network-api-go (Unset User Agent)")
}

// SetUserAgent sets the User-Agent header for requests
func (e *EnkaNetworkAPI) SetUserAgent(userAgent string) {
	e.userAgent = userAgent
}

// SetLogOut sets the log output for the EnkaNetworkAPI instance
//
// Parameters:
//
//	w: The log output
//
// See github.com/withmandala/go-log
func (e *EnkaNetworkAPI) SetLogOut(w log.FdWriter) {
	e.log = log.New(w).WithColor()
}

// SetDebug sets the debug mode for the EnkaNetworkAPI instance
func (e *EnkaNetworkAPI) SetDebug(debug bool) {
	if debug {
		e.log.WithDebug()
	} else {
		e.log.WithoutDebug()
	}
}
