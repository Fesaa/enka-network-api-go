package starrail

import "github.com/Fesaa/enka-network-api-go/hash"

type MultiplePathAvatarConfig struct {
	AvatarID         int         `json:"AvatarID"`
	Gender           string      `json:"Gender"`
	BaseAvatarID     int         `json:"BaseAvatarID"`
	Desc             hash.UInt64 `json:"Desc"`
	ChangeConfigPath string      `json:"ChangeConfigPath"`
}
