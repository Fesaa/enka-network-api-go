package data

import (
	_ "embed"

	"github.com/Fesaa/enka-network-api-go/genshin"
)

func (m *memoryCache) GetNameCardName(id int) *string {
	if name, ok := m.GenshinNameCards.Get(id); ok {
		return &name
	}
	return nil
}

func (m *memoryCache) HasNameCard(id int) bool {
	_, ok := m.GenshinNameCards.Get(id)
	return ok
}

// If this returns an error, I have the update the lib.
// You can mostly assume it'll return succesfully.
func (m *memoryCache) GetProfileIcon(id string) string {
	icon, ok := m.GenshinProfileIcons.Get(id)
	if !ok {
		return ""
	}
	return icon.IconPath
}

func (m *memoryCache) GetGenshinCharacterData(name string) *genshin.CharacterData {
	if character, ok := m.GenshinCharacterData.Get(name); ok {
		return character
	}
	return nil
}

func (m *memoryCache) GetAllGenshinCharacterData() []*genshin.CharacterData {
	var characters []*genshin.CharacterData = make([]*genshin.CharacterData, 0, m.GenshinCharacterData.Len())
	m.GenshinCharacterData.ForEach(func(_ string, character *genshin.CharacterData) {
		characters = append(characters, character)
	})
	return characters
}

func (m *memoryCache) GetGenshinMaterial(id int) *genshin.RawMaterial {
	if material, ok := m.GenshinMaterials.Get(id); ok {
		return material
	}
	return nil
}
