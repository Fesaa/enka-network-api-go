package enkanetworkapigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/utils"
)

type genshinAPIImpl struct {
	api *EnkaNetworkAPI
	log *slog.Logger
}

func newGenshinAPI(api *EnkaNetworkAPI, log *slog.Logger) GenshinAPI {
	return &genshinAPIImpl{
		api: api,
		log: log,
	}
}

func (g *genshinAPIImpl) Fetch(uid string, showCaseInfo bool, success utils.Consumer[*genshin.RawGenshinUser], failure utils.Consumer[error]) {
	go func(uid string, showCaseInfo bool, success func(*genshin.RawGenshinUser), failure func(error)) {
		user, err := g.FetchAndReturn(uid, showCaseInfo)
		if err != nil {
			if failure == nil {
				g.log.Warn("Provided failure call is nil, ignoring...")
				return
			}

			failure(err)
			return
		}
		if success == nil {
			g.log.Warn("Provided success call is nil, ignoring...")
			return
		}

		success(user)
	}(uid, showCaseInfo, success, failure)
}

func (g *genshinAPIImpl) FetchAndReturn(uid string, showCaseInfo bool) (*genshin.RawGenshinUser, error) {
	g.log.Debug("Fetching Genshin user with", "uid", uid)
	if _, err := strconv.Atoi(uid); err != nil || (len(uid) != 9 && len(uid) != 10) {
		return nil, errors.New("enka-network-api-go: UID must be a number, and 9 or 10 characters long")
	}

	cachedUser := g.api.cache.GetGenshinUser(uid)
	if cachedUser != nil {
		g.log.Debug("Returning from cache...", "uid", uid)
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
		g.log.Debug("Returned a non 200 status code. Got ", "status_code", req.StatusCode)

		var error string
		data, err := io.ReadAll(req.Body)
		if err != nil {
			g.log.Error("Failed to read body:", "error", err.Error())
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

	g.api.cache.AddGenshinUser(&user)
	return &user, nil
}

func (g *genshinAPIImpl) Icon(key string) string {
	url := BASE_GENSHIN_UI_URL + key

	if strings.HasSuffix(url, ".png") {
		return url
	}

	return url + ".png"
}

func (g *genshinAPIImpl) NameCard(id int) *genshin.NameCard {
	cardName := g.api.cache.GetNameCardName(id)
	if cardName == nil {
		return nil
	}

	return &genshin.NameCard{
		RawId: id,
		Url:   g.Icon(*cardName),
	}
}

func (g *genshinAPIImpl) ProfileId(id int) string {
	return g.api.cache.GetProfileIcon(fmt.Sprintf("%d", id))
}

func (g *genshinAPIImpl) CharacterData(character *genshin.UserCharacter) *genshin.CharacterData {
	return g.api.cache.GetGenshinCharacterData(fmt.Sprint(character.Id))
}

func (g *genshinAPIImpl) CharacterDataById(id string) *genshin.CharacterData {
	return g.api.cache.GetGenshinCharacterData(id)
}

func (g *genshinAPIImpl) Material(id int) *genshin.RawMaterial {
	return g.api.cache.GetGenshinMaterial(id)
}
