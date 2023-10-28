package genshin

import "github.com/Fesaa/enka-network-api-go/localization"

type Material struct {
	Id              int
	Name            string
	NameHash        string
	Description     string
	DescriptionHash string
	IconKey         string
	Pictures        []string
	ItemType        ItemType

	// Only avaible for ITEM_MATERIAL
	// Because I'm lazy it might be present for others
	// be sure to check ItemType
	MaterialType string
	Stars        string
}

type ItemType string

const (
	ITEM_VIRTUAL  ItemType = "ITEM_VIRTUAL"
	ITEM_MATERIAL ItemType = "ITEM_MATERIAL"

	UNIMPLEMENTED ItemType = "UNIMPLEMENTED"
)

type RawMaterial struct {
	Id              int      `json:"id"`
	NameHash        string   `json:"nameTextMapHash"`
	DescriptionHash string   `json:"descTextMapHash"`
	IconKey         string   `json:"icon"`
	Pictures        []string `json:"picPath"`
	ItemType        string   `json:"itemType"`
	MaterialType    string   `json:"materialType,omitempty"`
	Stars           string   `json:"rankLevel,omitempty"`
}

func (r RawMaterial) ToMaterial() Material {
	return Material{
		Id:              r.Id,
		Name:            localization.GetGenshinLocaleOrHash(&localization.HashString{Hash: r.NameHash}),
		NameHash:        r.NameHash,
		Description:     localization.GetGenshinLocaleOrHash(&localization.HashString{Hash: r.DescriptionHash}),
		DescriptionHash: r.DescriptionHash,
		IconKey:         r.IconKey,
		Pictures:        r.Pictures,
		ItemType:        ItemType(r.ItemType),

		MaterialType: r.MaterialType,
		Stars:        r.Stars,
	}
}
