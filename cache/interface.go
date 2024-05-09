package cache

import (
	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

type EnkaHttpCache interface {
	AddGenshinUser(*genshin.RawGenshinUser)
	GetGenshinUser(string) *genshin.RawGenshinUser

	AddHonkaiUser(*starrail.RawHonkaiUser)
	GetHonkaiUser(string) *starrail.RawHonkaiUser
}
