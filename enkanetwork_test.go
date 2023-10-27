package enkanetworkapigo

import (
	"testing"
)

func TestFetchHonkaiUser(t *testing.T) {

	api := New("enka-network-api-tests")
	api.SetDebug(true)

	api.FetchHonkaiUser("714656501",
		func(hu *HonkaiUser) {
			t.Log(hu)
		},
		func(err error) {
			t.Fatal(err)
		})
}
