package enkanetworkapigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/utils"
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
func (e *EnkaNetworkAPI) FetchGenshinUser(uid string, showCaseInfo bool, success utils.Consumer[*genshin.RawGenshinUser], failure utils.Consumer[error]) {
	go func(uid string, showCaseInfo bool, success func(*genshin.RawGenshinUser), failure func(error)) {
		user, err := e.FetchGenshinUserAndReturn(uid, showCaseInfo)
		if err != nil {
			if failure == nil {
				e.log.Warn("Provided failure call is nil, ignoring...")
				return
			}

			failure(err)
			return
		}
		if success == nil {
			e.log.Warn("Provided success call is nil, ignoring...")
			return
		}

		success(user)
	}(uid, showCaseInfo, success, failure)
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
//	The error, will be MaintenanceError if the API is in maintenance
//
// See FetchGenshinUser for an asynchronous version
func (e *EnkaNetworkAPI) FetchGenshinUserAndReturn(uid string, showCaseInfo bool) (*genshin.RawGenshinUser, error) {
	e.log.Debug("Fetching Genshin user with", "uid", uid)
	if _, err := strconv.Atoi(uid); err != nil || (len(uid) != 9 && len(uid) != 10) {
		return nil, errors.New("enka-network-api-go: UID must be a number, and 9 or 10 characters long")
	}

	cachedUser := e.cache.GetGenshinUser(uid)
	if cachedUser != nil {
		e.log.Debug("Returning from cache...", "uid", uid)
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

	if req.StatusCode == 424 {
		return nil, MaintenanceError
	}

	if req.StatusCode != 200 {
		e.log.Debug("Returned a non 200 status code. Got ", "status_code", req.StatusCode)

		var error string
		data, err := io.ReadAll(req.Body)
		if err != nil {
			e.log.Error("Failed to read body:", "error", err.Error())
			error = "Unknown: Failed to read body"
		} else {
			error = string(data)
		}
		return nil, fmt.Errorf("enka-network-api-go: Non 200 status code returned: %d\nBody: %s", req.StatusCode, error)
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

	e.cache.AddGenshinUser(&user)
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
	cardName := e.cache.GetNameCardName(id)
	if cardName == nil {
		return nil
	}

	return &genshin.NameCard{
		RawId: id,
		Url:   e.GetGenshinIcon(*cardName),
	}
}

// GetGenshinProfileIdentifier gets the profile identifier of a profile id
// returns an empty string if the id is invalid
func (e *EnkaNetworkAPI) GetGenshinProfileIdentifier(id int) string {
	return e.cache.GetProfileIcon(fmt.Sprintf("%d", id))
}

// GetGenshinCharacterData gets the character data of a character name
// Returns nil if the name is invalid
func (e *EnkaNetworkAPI) GetGenshinCharacterData(id string) *genshin.CharacterData {
	return e.cache.GetGenshinCharacterData(id)
}

// GetAllGenshinCharacterData gets all character data
func (e *EnkaNetworkAPI) GetAllGenshinCharacterData() []*genshin.CharacterData {
	return e.cache.GetAllGenshinCharacterData()
}

// GetGenshinMaterial gets the material data of a material id
// Use RawMaterial#ToMaterial to convert to a Material
// Returns nil if the id is invalid
func (e *EnkaNetworkAPI) GetGenshinMaterial(id int) *genshin.RawMaterial {
	return e.cache.GetGenshinMaterial(id)
}
