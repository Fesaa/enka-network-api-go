package enkanetworkapigo

import (
	"testing"
)

const OWN_UID = "714656501"

func TestFetchHonkaiUser(t *testing.T) {

	api := New("enka-network-api-tests")
	api.SetDebug(true)

	hu, err := api.FetchHonkaiUserAndReturn(OWN_UID)
	if err != nil {
		api.log.Error(err)
		t.Fatal(err)
		t.FailNow()
	}
	user := StarRailUserFromRaw(hu)
	if user.NickName != "Amelia" {
		t.Logf("Wanted Amelia got %s", user.NickName)
		t.Fail()
	}
	t.Log(user.Characters[0].LightCone.Hash)
	t.Log(*api.localization.GetHonkaiLocale(user.Characters[0].LightCone))
	data := api.GetStarRailCharacterData(user.Characters[0])
	t.Log(data)

	_, _ = api.FetchHonkaiUserAndReturn(OWN_UID)
}
