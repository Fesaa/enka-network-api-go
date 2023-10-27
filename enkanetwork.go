package enkanetworkapigo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/withmandala/go-log"
)

const BASE_URL = "https://enka.network/api/"
const BASE_GENSHIN_UI_URL = "https://enka.network/ui/"
const BASE_SR_UI_URL = "https://enka.network/ui/hsr/"

type EnkaNetworkAPI struct {
	userAgent string

	log *log.Logger

	cache        EnkaCache
	localization Localization
}

// New creates a new EnkaNetworkAPI instance
// From the documentation:
//
//	Please set a custom User-Agent header with your requests so I can track them better and help you if needed.
//
// See https://api.enka.network/ for API docs
func New(userAgent string) *EnkaNetworkAPI {

	return &EnkaNetworkAPI{
		userAgent: userAgent,

		log:          log.New(os.Stdout).WithColor(),
		cache:        NewMemoryCache(),
		localization: NewLocalization(),
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

// Returns the localization used to get actual strings
func (e *EnkaNetworkAPI) Loc() *Localization {
	return &e.localization
}

func (e *EnkaNetworkAPI) SetCache(cache EnkaCache) {
	e.cache = cache
}

// FetchHonkaiUser fetches a Honkai User from the Enka Network API
// Or returns from (redis) cache if available and not expired
//
// Parameters:
//
//	uid: The UID of the user to fetch
//	success: A function to call on success, with the HonkaiUser as the first parameter
//	failure: A function to call on failure, with the error as the first parameter
//
// See FetchHonkaiUserAndReturn for a synchronous version
func (e *EnkaNetworkAPI) FetchHonkaiUser(uid string, success func(*RawHonkaiUser), failure func(error)) {
	go func() {
		user, err := e.FetchHonkaiUserAndReturn(uid)
		if err != nil {
			failure(err)
			return
		}

		success(user)
	}()
}

// FetchHonkaiUserAndReturn fetches a Honkai User from the Enka Network API
// Or returns from (redis) cache if available and not expired
//
// Parameters:
//
//	uid: The UID of the user to fetch
//
// Returns:
//
//	The HonkaiUser, or nil if an error occurred
//
// See FetchHonkaiUser for an asynchronous version
func (e *EnkaNetworkAPI) FetchHonkaiUserAndReturn(uid string) (*RawHonkaiUser, error) {
	e.log.Debugf("Fetching Honkai User with UID %s", uid)

	cachedUser := e.cache.GetHonkaiUser(uid)
	if cachedUser != nil {
		e.log.Debug("Returning from cache...")
		return cachedUser, nil
	}

	req, err := http.Get(BASE_URL + "hsr/uid/" + uid + "/")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		e.log.Debugf("Returned a non 200 status code. Got %d", req.StatusCode)
		return nil, errors.New("enka-network-api-go: Non 200 status code returned: " + req.Status)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var user RawHonkaiUser
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}

	e.cache.AddHonkaiUser(&user)
	return &user, nil
}
