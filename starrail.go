package enkanetworkapigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"

	"github.com/Fesaa/enka-network-api-go/data"
	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

type starRailAPIImpl struct {
	api  EnkaNetworkAPI
	data data.StarRailData
	log  *slog.Logger
}

func newStarRail(api EnkaNetworkAPI, log *slog.Logger) StarRailAPI {
	return &starRailAPIImpl{
		api:  api,
		data: api.Data().StarRailData(),
		log:  log,
	}
}

func (sr *starRailAPIImpl) Fetch(uid string, success utils.Consumer[*starrail.RawHonkaiUser], failure utils.Consumer[error]) {
	go func(uid string, success func(*starrail.RawHonkaiUser), failure func(error)) {
		user, err := sr.FetchAndReturn(uid)
		if err != nil {
			if failure == nil {
				sr.log.Warn("Provided failure call is nil, ignoring...")
				return
			}

			failure(err)
			return
		}
		if success == nil {
			sr.log.Warn("Provided success call is nil, ignoring...")
			return
		}

		success(user)
	}(uid, success, failure)
}

func (sr *starRailAPIImpl) FetchAndReturn(uid string) (*starrail.RawHonkaiUser, error) {
	sr.log.Debug("Fetching Honkai User with UID ", "uid", uid)
	if _, err := strconv.Atoi(uid); err != nil || len(uid) != 9 {
		return nil, errors.New("enka-network-api-go: UID must be a number, and 9 characters long")
	}

	cachedUser := sr.api.Cache().GetHonkaiUser(uid)
	if cachedUser != nil {
		sr.log.Debug("Returning from cache...", "uid", uid)
		return cachedUser, nil
	}

	req, err := sr.api.HttpClient().Get(BASE_URL + "hsr/uid/" + uid + "/")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode == 424 {
		return nil, MaintenanceError
	}

	if req.StatusCode != 200 {
		sr.log.Debug("Returned a non 200 status code. Got ", "status_code", req.StatusCode)

		var e string
		var bytes []byte
		bytes, err = io.ReadAll(req.Body)
		if err != nil {
			sr.log.Error("Failed to read body:", "err", err.Error())
			e = "Unknown: Failed to read body"
		} else {
			e = string(bytes)
		}
		return nil, fmt.Errorf("enka-network-api-go: Non 200 status code returned: %d\nBody: %s", req.StatusCode, e)
	}

	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var user starrail.RawHonkaiUser
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}

	sr.api.Cache().AddHonkaiUser(&user)
	return &user, nil
}

func (sr *starRailAPIImpl) CharacterData(userCharacter *starrail.UserCharacter) *starrail.CharacterData {
	if userCharacter == nil {
		return nil
	}
	return sr.CharacterDataById(fmt.Sprint(userCharacter.AvatarId))
}

func (sr *starRailAPIImpl) CharacterDataById(uid string) *starrail.CharacterData {
	return sr.data.CharacterData(uid)
}

func (sr *starRailAPIImpl) Icon(key string) string {
	url := fmt.Sprintf("%s%s", BASE_SR_UI_URL, key)
	if strings.HasSuffix(url, ".png") {
		return url
	}

	return fmt.Sprintf("%s.png", url)
}

func (sr *starRailAPIImpl) AvatarKey(avatarId string) string {
	return sr.data.AvatarKey(avatarId)
}

func (sr *starRailAPIImpl) RelicData(relic *starrail.Relic) *starrail.RelicData {
	if relic == nil {
		return nil
	}
	return sr.RelicDataById(fmt.Sprint(relic.RelicID))
}

func (sr *starRailAPIImpl) RelicDataById(relicId string) *starrail.RelicData {
	return sr.data.RelicData(relicId)
}

func (sr *starRailAPIImpl) LightConeData(lightcone *starrail.LightCone) *starrail.LightConeData {
	if lightcone == nil {
		return nil
	}
	return sr.LightConeDataById(fmt.Sprint(lightcone.LightConeID))
}

func (sr *starRailAPIImpl) LightConeDataById(lightConeId string) *starrail.LightConeData {
	return sr.data.LightConeData(lightConeId)
}
