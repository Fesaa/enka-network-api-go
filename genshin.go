package enkanetworkapigo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

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
// showCaseInfo: Whether to show the showcase info or not
//
//	Consider not showing the showcase info if you don't need it
//
// success: The callback to call when the user is fetched successfully
//
// failure: The callback to call when the user could not be fetched
//
// See FetchGenshinUserAndReturn for a synchronous version
func (e *EnkaNetworkAPI) FetchGenshinUser(uid string, showCaseInfo bool, success func(*genshin.RawGenshinUser), failure func(error)) {
	go func() {
		user, err := e.FetchGenshinUserAndReturn(uid, showCaseInfo)
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
// showCaseInfo: Whether to show the showcase info or not
//
//	Consider not showing the showcase info if you don't need it
//
// Returns:
//
// *RawGenshinUser: The fetched user
//
// See FetchGenshinUser for an asynchronous version
func (e *EnkaNetworkAPI) FetchGenshinUserAndReturn(uid string, showCaseInfo bool) (*genshin.RawGenshinUser, error) {
	e.log.Debugf("Fetching Genshin user with uid %s", uid)

	cachedUser := cache.Get().GetGenshinUser(uid)
	if cachedUser != nil {
		e.log.Debug("Returning from cache...")
		return cachedUser, nil
	}

	var suffix string = ""
	if !showCaseInfo {
		suffix = "?info"
	}

	req, err := http.Get(BASE_URL + "uid/" + uid + suffix)
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

// GetGenshinIcon retrieves the URL of a Genshin Impact icon
func (e *EnkaNetworkAPI) GetGenshinIcon(key string) string {
	url := BASE_GENSHIN_UI_URL + key

	if strings.HasSuffix(url, ".png") {
		return url
	}

	return url + ".png"
}

// GetGenshinNameCard gets the NameCard of an id
// Returns nil if the id is invalid
//
// You can access the enka-network url with the Url field
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

// GetGenshinProfileIdentifier gets the profile identifier of a profile id
// Returns nil if the id is invalid
// Falls back to 4.0 data, if not changed > 4.1
func (e *EnkaNetworkAPI) GetGenshinProfileIdentifier(profileId int) string {
	return cache.Get().GetProfileIcon(profileId)
}

// GetGenshinCharacterData gets the character data of a character name
// Returns nil if the name is invalid
func (e *EnkaNetworkAPI) GetGenshinCharacterData(id string) *genshin.CharacterData {
	return cache.Get().GetGenshinCharacterData(id)
}

// GetAllGenshinCharacterData gets all character data
func (e *EnkaNetworkAPI) GetAllGenshinCharacterData() []*genshin.CharacterData {
	return cache.Get().GetAllGenshinCharacterData()
}

// GetGenshinMaterial gets the material data of a material id
// Use RawMaterial#ToMaterial to convert to a Material
// Returns nil if the id is invalid
func (e *EnkaNetworkAPI) GetGenshinMaterial(id int) *genshin.RawMaterial {
	return cache.Get().GetGenshinMaterial(id)
}
