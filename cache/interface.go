package cache

import (
	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

type EnkaHttpCache interface {
	AddGenshinUser(*genshin.RawUser)
	GetGenshinUser(string) *genshin.RawUser

	AddHSRUser(*starrail.RawUser)
	GetHsrUser(string) *starrail.RawUser
}
