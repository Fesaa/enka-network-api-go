package enkanetworkapigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Fesaa/enka-network-api-go/cache"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

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
func (e *EnkaNetworkAPI) FetchHonkaiUser(uid string, success func(*starrail.RawHonkaiUser), failure func(error)) {
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
func (e *EnkaNetworkAPI) FetchHonkaiUserAndReturn(uid string) (*starrail.RawHonkaiUser, error) {
	e.log.Debugf("Fetching Honkai User with UID %s", uid)

	cachedUser := cache.Get().GetHonkaiUser(uid)
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

	var user starrail.RawHonkaiUser
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}

	cache.Get().AddHonkaiUser(&user)
	return &user, nil
}

func (e *EnkaNetworkAPI) GetStarRailCharacterData(userCharacter starrail.UserCharacter) *starrail.CharacterData {
	return e.GetStarRailCharacterDataById(fmt.Sprint(userCharacter.AvatarId))
}

func (e *EnkaNetworkAPI) GetStarRailCharacterDataById(uid string) *starrail.CharacterData {
	return cache.Get().GetStarRailCharacterData(uid)
}
