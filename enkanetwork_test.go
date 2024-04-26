package enkanetworkapigo

import (
	"strings"
	"sync"
	"testing"

	"github.com/Fesaa/enka-network-api-go/cache"
	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

const OWN_UID = "714656501"

var api *EnkaNetworkAPI

func TestFetchGenshinUser(t *testing.T) {

	if api == nil {
		var e error

		api, e = New("enka-network-api-tests", cache.Default())
		if e != nil {
			t.Fatal(e)
		}
	}

	rgu, err := api.FetchGenshinUserAndReturn("618285856", true)
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

	material := api.GetGenshinMaterial(101)
	t.Log(material)
	if material == nil {
		t.Logf("Could not find material")
		t.FailNow()
	}

	if !strings.Contains(material.IconKey, "UI_ItemIcon") {
		t.Logf("Wanted UI_ItemIcon got %s", material.IconKey)
		t.Fail()
	}

	id := api.GetGenshinProfileIdentifier(user.ProfilePicture.AvatarId)
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	t.Logf("Pfp: %s", id)

}

func TestFetchHonkaiUser(t *testing.T) {

	if api == nil {
		var e error
		api, e = New("enka-network-api-tests", cache.Default())
		if e != nil {
			t.Fatal(e)
		}
	}

	hu, err := api.FetchHonkaiUserAndReturn(OWN_UID)
	if err == MaintenanceError {
		t.Logf("API in Maintenance skipping test")
		t.SkipNow()
	}

	if err != nil {
		api.log.Error(err.Error())
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
		name := lightCone.Name()

		characterData := api.GetStarRailCharacterData(&character)
		avatarIcon := api.GetStarRailIcon(characterData.AvatarSideIconPath)
		t.Logf("Character %s (%d) of path %s has lightCone %s", characterData.Name(), character.Level, characterData.Path, name)
		t.Logf("Avatar icon: %s", avatarIcon)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)

	api.FetchHonkaiUser(OWN_UID,
		func(rhu *starrail.RawHonkaiUser) {
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
