package enkanetworkapigo

import (
	"errors"
	"log/slog"

	"github.com/Fesaa/enka-network-api-go/cache"
	"github.com/Fesaa/enka-network-api-go/localization"
)

const BASE_URL = "https://enka.network/api/"
const BASE_GENSHIN_UI_URL = "https://enka.network/ui/"
const BASE_SR_UI_URL = "https://enka.network/ui/hsr/"

var MaintenanceError error = errors.New("enka-network-api-go: The API is currently in maintenance")

type EnkaNetworkAPI struct {
	userAgent string

	log *slog.Logger
}

// New creates a new EnkaNetworkAPI instance
//
// Parameters:
//
//	userAgent: The User-Agent header for requests
//	cacheType: The cache type to use
//	loggers: Optional custom loggers
//
// # Will also initialize the cache and localization
//
// From the documentation:
//
//	Please set a custom User-Agent header with your requests so I can track them better and help you if needed.
//
// See https://api.enka.network/ for API docs
func New(userAgent string, cacheType cache.CacheType, loggers ...*slog.Logger) (*EnkaNetworkAPI, error) {
	e := cache.Init(cacheType)
	if e != nil {
		return nil, e
	}

	var logger *slog.Logger
	if len(loggers) > 0 {
		logger = loggers[0]
	} else {
		logger = slog.Default()
	}

	localization.Init(logger)
	return &EnkaNetworkAPI{
		userAgent: userAgent,
		log:       logger,
	}, nil
}

// NewDefaultUserAgent creates a new EnkaNetworkAPI instance with the default User-Agent header
// Consider using New or SetUserAgent instead
func NewDefault() (*EnkaNetworkAPI, error) {
	return New("enka-network-api-go (Unset User Agent)", cache.MEMORY)
}

// SetUserAgent sets the User-Agent header for requests
func (e *EnkaNetworkAPI) SetUserAgent(userAgent string) {
	e.userAgent = userAgent
}
