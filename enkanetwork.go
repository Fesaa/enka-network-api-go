package enkanetworkapigo

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Fesaa/enka-network-api-go/cache"
	"github.com/Fesaa/enka-network-api-go/data"
	"github.com/Fesaa/enka-network-api-go/localization"
)

const BASE_URL = "https://enka.network/api/"
const BASE_GENSHIN_UI_URL = "https://enka.network/ui/"
const BASE_SR_UI_URL = "https://enka.network/ui/hsr/"

var MaintenanceError = errors.New("enka-network-api-go: The API is currently in maintenance")

type enkaNetworkAPIImpl struct {
	userAgent string
	data      data.EnkaData
	cache     cache.EnkaHttpCache

	http *http.Client

	log *slog.Logger

	starRailAPI StarRailAPI
	genshinApi  GenshinAPI
}

func WithCustomUserAgent(userAgent string) (EnkaNetworkAPI, error) {
	return New(userAgent, cache.Default(slog.Default()))
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
//	Please set a custom User-Agent header with your requests, so I can track them better and help you if needed.
//
// See https://api.enka.network/ for API docs
func New(userAgent string, httpCache cache.EnkaHttpCache, loggers ...*slog.Logger) (EnkaNetworkAPI, error) {
	var logger *slog.Logger
	if len(loggers) > 0 {
		logger = loggers[0]
	} else {
		logger = slog.Default()
	}

	d, err := data.New(logger)
	if err != nil {
		return nil, err
	}

	localization.Init(logger)
	api := &enkaNetworkAPIImpl{
		userAgent: userAgent,
		log:       logger,
		data:      d,
		cache:     httpCache,
	}

	api.http = &http.Client{Transport: &httpUserAgentInterceptor{
		api:         api,
		transporter: http.DefaultTransport,
	}}
	api.starRailAPI = newStarRail(api, logger)
	api.genshinApi = newGenshinAPI(api, logger)
	return api, nil
}

// NewDefault creates a new EnkaNetworkAPI instance with the default User-Agent header
// Consider using New or SetUserAgent instead
func NewDefault() (EnkaNetworkAPI, error) {
	httpCache := cache.Default(slog.Default())
	return New("enka-network-api-go (Unset User Agent)", httpCache)
}

// SetUserAgent sets the User-Agent header for requests
func (e *enkaNetworkAPIImpl) SetUserAgent(userAgent string) {
	e.userAgent = userAgent
}

// GetUserAgent returns the User-Agent header for requests
func (e *enkaNetworkAPIImpl) GetUserAgent() string {
	return e.userAgent
}

func (e *enkaNetworkAPIImpl) HttpClient() *http.Client {
	return e.http
}

func (e *enkaNetworkAPIImpl) StarRail() StarRailAPI {
	return e.starRailAPI
}

func (e *enkaNetworkAPIImpl) Genshin() GenshinAPI {
	return e.genshinApi
}

func (e *enkaNetworkAPIImpl) Data() data.EnkaData {
	return e.data
}

func (e *enkaNetworkAPIImpl) Cache() cache.EnkaHttpCache {
	return e.cache
}
