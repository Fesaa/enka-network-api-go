package data

import (
	_ "embed"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/utils"
)

type genshinData struct {
	GenshinNameCards     *utils.Map[int, string]
	GenshinProfileIcons  *utils.Map[string, *genshin.ProfilePicture]
	GenshinCharacterData *utils.Map[string, *genshin.CharacterData]
	GenshinMaterials     *utils.Map[int, *genshin.RawMaterial]
}

func (m *genshinData) NameCardName(id int) *string {
	if name, ok := m.GenshinNameCards.Get(id); ok {
		return &name
	}
	return nil
}

func (m *genshinData) HasNameCard(id int) bool {
	_, ok := m.GenshinNameCards.Get(id)
	return ok
}

// If this returns an error, I have the update the lib.
// You can mostly assume it'll return succesfully.
func (m *genshinData) ProfileIcon(id string) string {
	icon, ok := m.GenshinProfileIcons.Get(id)
	if !ok {
		return ""
	}
	return icon.IconPath
}

func (m *genshinData) CharacterData(name string) *genshin.CharacterData {
	if character, ok := m.GenshinCharacterData.Get(name); ok {
		return character
	}
	return nil
}

func (m *genshinData) Characters() []*genshin.CharacterData {
	var characters []*genshin.CharacterData = make([]*genshin.CharacterData, 0, m.GenshinCharacterData.Len())
	m.GenshinCharacterData.ForEach(func(_ string, character *genshin.CharacterData) {
		characters = append(characters, character)
	})
	return characters
}

func (m *genshinData) Material(id int) *genshin.RawMaterial {
	if material, ok := m.GenshinMaterials.Get(id); ok {
		return material
	}
	return nil
}
