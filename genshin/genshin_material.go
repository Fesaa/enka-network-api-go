package genshin

import "github.com/Fesaa/enka-network-api-go/localization"

type Material struct {
	Id              int
	Name            string
	NameHash        localization.HashInt
	Description     string
	DescriptionHash localization.HashInt
	IconKey         string
	Pictures        []string
	ItemType        ItemType

	// Only avaible for ITEM_MATERIAL
	// Because I'm lazy it might be present for others
	// be sure to check ItemType
	MaterialType string
	Stars        int
}

type ItemType string

const (
	ITEM_VIRTUAL  ItemType = "ITEM_VIRTUAL"
	ITEM_MATERIAL ItemType = "ITEM_MATERIAL"

	UNIMPLEMENTED ItemType = "UNIMPLEMENTED"
)

type RawMaterial struct {
	Id              int      `json:"id"`
	NameHash        int      `json:"nameTextMapHash"`
	DescriptionHash int      `json:"descTextMapHash"`
	IconKey         string   `json:"icon"`
	Pictures        []string `json:"picPath"`
	ItemType        string   `json:"itemType"`
	MaterialType    string   `json:"materialType,omitempty"`
	Stars           int      `json:"rankLevel,omitempty"`
}

func (r RawMaterial) ToMaterial() Material {

	name := localization.HashInt{Hash: r.NameHash}
	description := localization.HashInt{Hash: r.DescriptionHash}

	return Material{
		Id:              r.Id,
		Name:            localization.GetGenshinLocaleOrHash(&name),
		NameHash:        name,
		Description:     localization.GetGenshinLocaleOrHash(&description),
		DescriptionHash: description,
		IconKey:         r.IconKey,
		Pictures:        r.Pictures,
		ItemType:        ItemType(r.ItemType),

		MaterialType: r.MaterialType,
		Stars:        r.Stars,
	}
}
