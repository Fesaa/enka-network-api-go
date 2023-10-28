package genshin

import (
	"fmt"
	"strings"

	"github.com/Fesaa/enka-network-api-go/localization"
)

type CharacterData struct {
	Element          string         `json:"Element"`
	TalentsImageKeys []string       `json:"Consts"`
	SkillOrder       []int          `json:"SkillOrder"`
	Skills           map[int]string `json:"Skills"`
	NameHash         int            `json:"NameTextMapHash"`
	SideIconKey      string         `json:"SideIconName"`
	WeaponType       string
	RawWeaponType    string          `json:"WeaponType"`
	Costumes         map[int]Costume `json:"Costumes"`
}

func (c *CharacterData) GetHash() string {
	return fmt.Sprint(c.NameHash)
}

func (c *CharacterData) Name() string {
	return localization.GetGenshinLocaleOrHash(c)
}

// Returns the icon of the character when looking at the camera
// This is a key, and should be parsed with EnkaNetworkApi#GetGenshinIcon
//
// This is no longer needed after 4.1
func (c *CharacterData) GetSideIconKey() string {
	return strings.Replace(c.SideIconKey, "_Side", "", 1)
}

type Costume struct {
	SideIconKey string `json:"sideIconName"`
	IconKey     string `json:"icon"`
	Art         string `json:"art"`
}

func WrapWeaponType(key string) string {
	switch key {
	case "WEAPON_SWORD_ONE_HAND":
		return "Sword"
	case "WEAPON_CATALYST":
		return "Catalyst"
	case "WEAPON_CLAYMORE":
		return "Claymore"
	case "WEAPON_BOW":
		return "Bow"
	case "WEAPON_POLE":
		return "Polearm"
	default:
		return "Unknown"
	}
}
