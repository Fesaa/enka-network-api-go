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

type Option func(e *enkaNetworkAPI)

type enkaNetworkAPI struct {
	userAgent string
	data      data.EnkaData
	cache     cache.EnkaHttpCache

	http *http.Client

	log *slog.Logger

	starRailAPI StarRailAPI
	genshinApi  GenshinAPI
}

func WithHttpClient(httpClient *http.Client) Option {
	return func(e *enkaNetworkAPI) {
		e.http = httpClient
	}
}

func WithHttpCache(cache cache.EnkaHttpCache) Option {
	return func(e *enkaNetworkAPI) {
		e.cache = cache
	}
}

func WithCustomUserAgent(userAgent string) Option {
	return func(e *enkaNetworkAPI) {
		e.userAgent = userAgent
	}
}

// New creates a new EnkaNetworkAPI instance
// Will also initialize the cache and localization
//
// From the documentation:
//
//	Please set a custom User-Agent header with your requests, so I can track them better and help you if needed.
//
// See https://api.enka.network/ for API docs
func New(opts ...Option) (EnkaNetworkAPI, error) {
	api := &enkaNetworkAPI{}

	for _, opt := range opts {
		opt(api)
	}

	if api.log == nil {
		api.log = slog.Default()
	}

	if api.cache == nil {
		api.cache = cache.Default(api.log)
	}

	if api.http == nil {
		api.http = &http.Client{Transport: &HttpUserAgentInterceptor{
			api:         api,
			transporter: http.DefaultTransport,
		}}
	}

	localization.Init(api.log)

	d, err := data.New(api.log)
	if err != nil {
		return nil, err
	}

	api.data = d

	api.starRailAPI = newStarRail(api, api.log)
	api.genshinApi = newGenshinAPI(api, api.log)
	return api, nil
}

// SetUserAgent sets the User-Agent header for requests
func (e *enkaNetworkAPI) SetUserAgent(userAgent string) {
	e.userAgent = userAgent
}

// GetUserAgent returns the User-Agent header for requests
func (e *enkaNetworkAPI) GetUserAgent() string {
	return e.userAgent
}

func (e *enkaNetworkAPI) HttpClient() *http.Client {
	return e.http
}

func (e *enkaNetworkAPI) StarRail() StarRailAPI {
	return e.starRailAPI
}

func (e *enkaNetworkAPI) Genshin() GenshinAPI {
	return e.genshinApi
}

func (e *enkaNetworkAPI) Data() data.EnkaData {
	return e.data
}

func (e *enkaNetworkAPI) Cache() cache.EnkaHttpCache {
	return e.cache
}
