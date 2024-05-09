package enkanetworkapigo

import (
	"errors"
	"log/slog"

	"github.com/Fesaa/enka-network-api-go/cache"
	"github.com/Fesaa/enka-network-api-go/localization"
	"github.com/Fesaa/enka-network-api-go/utils"
)

const BASE_URL = "https://enka.network/api/"
const BASE_GENSHIN_UI_URL = "https://enka.network/ui/"
const BASE_SR_UI_URL = "https://enka.network/ui/hsr/"

var MaintenanceError error = errors.New("enka-network-api-go: The API is currently in maintenance")

type enkaNetworkAPIImpl struct {
	userAgent string
	cache     cache.EnkaCache

	log *slog.Logger

	starRailAPI StarRailAPI
	genshinApi  GenshinAPI
}

// New creates a new EnkaNetworkAPI instance
//
// Parameters:
//
//	userAgent: The User-Agent header for requests
//	cacheSupplier: A function that returns a cache instance
//	loggers: Optional custom loggers
//
// # Will also initialize the cache and localization
//
// From the documentation:
//
//	Please set a custom User-Agent header with your requests so I can track them better and help you if needed.
//
// See https://api.enka.network/ for API docs
func New(userAgent string, cacheSupplier utils.ErrorSupplier[cache.EnkaCache], loggers ...*slog.Logger) (EnkaNetworkAPI, error) {
	cache, err := cacheSupplier()
	if err != nil {
		return nil, err
	}

	var logger *slog.Logger
	if len(loggers) > 0 {
		logger = loggers[0]
	} else {
		logger = slog.Default()
	}

	localization.Init(logger)
	api := &enkaNetworkAPIImpl{
		userAgent: userAgent,
		log:       logger,
		cache:     cache,
	}
	api.starRailAPI = newStarRail(api, logger)
	api.genshinApi = newGenshinAPI(api, logger)
	return api, nil
}

// NewDefaultUserAgent creates a new EnkaNetworkAPI instance with the default User-Agent header
// Consider using New or SetUserAgent instead
func NewDefault() (EnkaNetworkAPI, error) {
	return New("enka-network-api-go (Unset User Agent)", cache.Default())
}

// SetUserAgent sets the User-Agent header for requests
func (e *enkaNetworkAPIImpl) SetUserAgent(userAgent string) {
	e.userAgent = userAgent
}

func (e *enkaNetworkAPIImpl) StarRail() StarRailAPI {
	return e.starRailAPI
}

func (e *enkaNetworkAPIImpl) Genshin() GenshinAPI {
	return e.genshinApi
}

func (e *enkaNetworkAPIImpl) Cache() cache.EnkaCache {
	return e.cache
}
