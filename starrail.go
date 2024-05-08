package enkanetworkapigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
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
func (e *EnkaNetworkAPI) FetchHonkaiUser(uid string, success utils.Consumer[*starrail.RawHonkaiUser], failure utils.Consumer[error]) {
	go func(uid string, success func(*starrail.RawHonkaiUser), failure func(error)) {
		user, err := e.FetchHonkaiUserAndReturn(uid)
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
	}(uid, success, failure)
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
//	The error, will be MaintenanceError if the API is in maintenance
//
// See FetchHonkaiUser for an asynchronous version
func (e *EnkaNetworkAPI) FetchHonkaiUserAndReturn(uid string) (*starrail.RawHonkaiUser, error) {
	e.log.Debug("Fetching Honkai User with UID ", "uid", uid)
	if _, err := strconv.Atoi(uid); err != nil || len(uid) != 9 {
		return nil, errors.New("enka-network-api-go: UID must be a number, and 9 characters long")
	}

	cachedUser := e.cache.GetHonkaiUser(uid)
	if cachedUser != nil {
		e.log.Debug("Returning from cache...", "uid", uid)
		return cachedUser, nil
	}

	req, err := http.Get(BASE_URL + "hsr/uid/" + uid + "/")
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

	var user starrail.RawHonkaiUser
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}

	e.cache.AddHonkaiUser(&user)
	return &user, nil
}

// GetStarRailCharacterData returns the CharacterData for the given UserCharacter
//
// Convenience function for GetStarRailCharacterDataById
func (e *EnkaNetworkAPI) GetStarRailCharacterData(userCharacter *starrail.UserCharacter) *starrail.CharacterData {
	if userCharacter == nil {
		return nil
	}
	return e.GetStarRailCharacterDataById(fmt.Sprint(userCharacter.AvatarId))
}

// GetStarRailCharacterDataById returns the CharacterData for the given UID
//
// Parameters:
//
//	uid: The UID of the character
//
// Returns:
//
//	The CharacterData, or nil if not found
func (e *EnkaNetworkAPI) GetStarRailCharacterDataById(uid string) *starrail.CharacterData {
	return e.cache.GetStarRailCharacterData(uid)
}

// GetStarRailIcon returns the URL of the StarRail icon for the given key
//
// You can change the BASE_SR_UI_URL if you need images from a different source
//
// Parameters:
//
//	key: The key of the icon
//
// Returns:
//
//	The URL of the icon
func (e *EnkaNetworkAPI) GetStarRailIcon(key string) string {
	url := fmt.Sprintf("%s%s", BASE_SR_UI_URL, key)
	if strings.HasSuffix(url, ".png") {
		return url
	}

	return fmt.Sprintf("%s.png", url)
}

// GetAllStarRailCharacters returns all StarRail characters
//
// Returns:
//
//	A slice of all StarRail characters
func (e *EnkaNetworkAPI) GetAllStarRailCharacters() []*starrail.CharacterData {
	return e.cache.GetAllStarRailCharacters()
}

// GetStarRailAvatarKey returns the avatar key for the given avatar ID
//
// Parameters:
//
//	avatarId: The avatar ID
//
// Returns:
//
//	The avatar key or id if not found
func (e *EnkaNetworkAPI) GetStarRailAvatarKey(avatarId string) string {
	return e.cache.GetStarRailAvatarKey(avatarId)
}

// GetStarRailRelicData returns the RelicData for the given Relic
//
// Parameters:
//
//	relic: The relic to get the data for
//
// Returns:
//
//	The RelicData, or nil if not found (or if the relic was nil)
func (e *EnkaNetworkAPI) GetStarRailRelicData(relic *starrail.Relic) *starrail.RelicData {
	if relic == nil {
		return nil
	}
	return e.cache.GetStarRailRelicData(fmt.Sprintf("%d", relic.RelicID))
}
