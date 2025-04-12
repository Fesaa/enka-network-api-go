package enkanetworkapigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"strconv"
	"strings"

	"github.com/Fesaa/enka-network-api-go/data"
	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/utils"
)

type genshinAPIImpl struct {
	api  EnkaNetworkAPI
	data data.GenshinData
	log  zerolog.Logger
}

func newGenshinAPI(api EnkaNetworkAPI, log zerolog.Logger) GenshinAPI {
	return &genshinAPIImpl{
		api:  api,
		data: api.Data().GenshinData(),
		log:  log.With().Str("handler", "genshin-api").Logger(),
	}
}

func (g *genshinAPIImpl) Fetch(uid string, showCaseInfo bool, success utils.Consumer[*genshin.RawUser], failure utils.Consumer[error]) {
	go func(uid string, showCaseInfo bool, success func(*genshin.RawUser), failure func(error)) {
		user, err := g.FetchAndReturn(uid, showCaseInfo)
		if err != nil {
			if failure == nil {
				g.log.Warn().Msg("Provided failure call is nil, ignoring...")
				return
			}

			failure(err)
			return
		}
		if success == nil {
			g.log.Warn().Msg("Provided success call is nil, ignoring...")
			return
		}

		success(user)
	}(uid, showCaseInfo, success, failure)
}

func (g *genshinAPIImpl) FetchAndReturn(uid string, showCaseInfo bool) (*genshin.RawUser, error) {
	g.log.Debug().Str("uid", uid).Msg("Fetching Genshin user with")
	if _, err := strconv.Atoi(uid); err != nil || (len(uid) != 9 && len(uid) != 10) {
		return nil, errors.New("enka-network-api-go: UID must be a number, and 9 or 10 characters long")
	}

	cachedUser := g.api.Cache().GetGenshinUser(uid)
	if cachedUser != nil {
		g.log.Debug().Str("uid", uid).Msg("Returning from cache...")
		return cachedUser, nil
	}

	var suffix string = ""
	if !showCaseInfo {
		suffix = "?info"
	}

	req, err := g.api.HttpClient().Get(BASE_URL + "uid/" + uid + suffix)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode == 424 {
		return nil, MaintenanceError
	}

	if req.StatusCode != 200 {
		g.log.Debug().Int("status_code", req.StatusCode).Msg("Returned a non 200 status code. Got ")

		var e string
		var bytes []byte
		bytes, err = io.ReadAll(req.Body)
		if err != nil {
			g.log.Error().Err(err).Msg("Failed to read body:")
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

	var user genshin.RawUser
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}

	g.api.Cache().AddGenshinUser(&user)
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
	cardName := g.data.NameCardName(id)
	if cardName == nil {
		return nil
	}

	return &genshin.NameCard{
		RawId: id,
		Url:   g.Icon(*cardName),
	}
}

func (g *genshinAPIImpl) ProfileId(id int) string {
	return g.data.ProfileIcon(fmt.Sprintf("%d", id))
}

func (g *genshinAPIImpl) CharacterData(character *genshin.UserCharacter) *genshin.CharacterData {
	return g.data.CharacterData(fmt.Sprint(character.Id))
}

func (g *genshinAPIImpl) CharacterDataById(id string) *genshin.CharacterData {
	return g.data.CharacterData(id)
}

func (g *genshinAPIImpl) Material(id int) *genshin.RawMaterial {
	return g.data.Material(id)
}
