package enkanetworkapigo

import (
	"sync"
	"testing"

	"github.com/Fesaa/enka-network-api-go/starrail"
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
	user := starrail.UserFromRaw(hu)
	if user.NickName != "Amelia" {
		t.Logf("Wanted Amelia got %s", user.NickName)
		t.Fail()
	}

	FirstCharacter := user.Characters[1]

	lightCone := FirstCharacter.LightCone
	name := lightCone.Name()

	characterData := api.GetStarRailCharacterData(FirstCharacter)

	t.Logf("Character %s has lightCone %s", characterData.Name(), name)

	t.Log(characterData)

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
