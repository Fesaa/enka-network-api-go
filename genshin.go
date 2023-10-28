package enkanetworkapigo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Fesaa/enka-network-api-go/cache"
	"github.com/Fesaa/enka-network-api-go/genshin"
)

// FetchGenshinUser fetches a Genshin User from the Enka Network API
// or returns from (redis) cache if available and not expired
//
// Parameters:
//
// uid: The UID of the user to fetch
//
// success: The callback to call when the user is fetched successfully
//
// failure: The callback to call when the user could not be fetched
//
// See FetchGenshinUserAndReturn for a synchronous version
func (e *EnkaNetworkAPI) FetchGenshinUser(uid string, success func(*genshin.RawGenshinUser), failure func(error)) {
	go func() {
		user, err := e.FetchGenshinUserAndReturn(uid)
		if err != nil {
			failure(err)
			return
		}
		success(user)
	}()
}

// FetchGenshinUserAndReturn fetches a Genshin User from the Enka Network API
// or returns from (redis) cache if available and not expired
//
// Parameters:
//
// uid: The UID of the user to fetch
//
// Returns:
//
// *RawGenshinUser: The fetched user
//
// See FetchGenshinUser for an asynchronous version
func (e *EnkaNetworkAPI) FetchGenshinUserAndReturn(uid string) (*genshin.RawGenshinUser, error) {
	e.log.Debugf("Fetching Genshin user with uid %s", uid)

	cachedUser := cache.Get().GetGenshinUser(uid)
	if cachedUser != nil {
		e.log.Debug("Returning from cache...")
		return cachedUser, nil
	}

	req, err := http.Get(BASE_URL + "uid/" + uid + "/")
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

	var user genshin.RawGenshinUser
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}

	cache.Get().AddGenshinUser(&user)
	return &user, nil
}

func (e *EnkaNetworkAPI) GetGenshinIcon(key string) string {
	return BASE_GENSHIN_UI_URL + key + ".png"
}

func (e *EnkaNetworkAPI) GetGenshinNameCard(id int) *genshin.NameCard {
	cardName := cache.Get().GetNameCardName(id)
	if cardName == nil {
		return nil
	}

	return &genshin.NameCard{
		RawId: id,
		Url:   e.GetGenshinIcon(*cardName),
	}

}
