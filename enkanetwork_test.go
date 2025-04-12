package enkanetworkapigo

import (
	"errors"
	"strings"
	"sync"
	"testing"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

const OWN_UID = "714656501"

var api EnkaNetworkAPI

func TestFetchGenshinUser(t *testing.T) {

	if api == nil {
		var e error

		api, e = New(WithCustomUserAgent("enka-network-api-tests"))
		if e != nil {
			t.Fatal(e)
		}
	}

	g := api.Genshin()
	rgu, err := g.FetchAndReturn("618285856", true)
	if err == MaintenanceError {
		t.Logf("API in Maintenance skipping test")
		t.SkipNow()
	}

	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}

	user := genshin.UserFromRaw(rgu)
	if user.NickName != "Algoinde" {
		t.Logf("Wanted Algoinde got %s", user.NickName)
		t.Fail()
	}

	if user.Level != 57 {
		t.Logf("Wanted level 57 got %d", user.Level)
		t.Fail()
	}

	material := g.Material(101)
	t.Log(material)
	if material == nil {
		t.Logf("Could not find material")
		t.FailNow()
	}

	if !strings.Contains(material.IconKey, "UI_ItemIcon") {
		t.Logf("Wanted UI_ItemIcon got %s", material.IconKey)
		t.Fail()
	}

	id := g.ProfileId(user.ProfilePicture.AvatarId)
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	t.Logf("Pfp: %s", id)

}

func TestFetchHonkaiUser(t *testing.T) {

	if api == nil {
		var e error
		api, e = New(WithCustomUserAgent("enka-network-api-tests"))
		if e != nil {
			t.Fatal(e)
		}
	}

	sr := api.StarRail()
	hu, err := sr.FetchAndReturn(OWN_UID)
	if errors.Is(err, MaintenanceError) {
		t.Logf("API in Maintenance skipping test")
		t.SkipNow()
	}

	if err != nil {
		t.Log(err.Error())
		t.Fatal(err)
		t.FailNow()
	}
	user := starrail.UserFromRaw(hu)
	if user.NickName != "Amelia" {
		t.Logf("Wanted Amelia got %s", user.NickName)
		t.Fail()
	}

	if len(user.Characters) < 2 {
		t.Logf("Wanted at least 2 characters got %d", len(user.Characters))
		t.Fail()
	}

	for _, character := range user.Characters {
		lightCone := character.LightCone

		var name, lightConeIcon string
		if lightCone != nil {
			lightConeData := sr.LightConeData(lightCone)
			name = lightCone.Name()
			lightConeIcon = sr.Icon(lightConeData.ImagePath)
		}

		characterData := sr.CharacterData(&character)
		avatarIcon := sr.Icon(characterData.AvatarSideIconPath)
		t.Logf("Character %s (%d) of path %s has lightCone %s with icon %s", characterData.Name(), character.Level, characterData.Path, name, lightConeIcon)
		t.Logf("Avatar icon: %s", avatarIcon)

		if len(character.Relics) == 0 {
			continue
		}

		relic := character.Relics[0]
		relicData := sr.RelicData(&relic)
		icon := sr.Icon(relicData.Icon)
		t.Logf("Relic %s with icon %s", relic.Name(), icon)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)

	sr.Fetch(OWN_UID,
		func(rhu *starrail.RawUser) {
			user := starrail.UserFromRaw(rhu)

			t.Logf("Found user %s with level %d", user.NickName, user.Level)

			wg.Add(-1)

		},
		func(err error) {
			t.Error(err)
			t.Fail()
			wg.Add(-1)
		})

	wg.Wait()

}
