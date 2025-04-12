package data

import (
	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

type starRailData struct {
	StarRailCharacterData *utils.Map[string, *starrail.CharacterData]
	StarRailAvatars       *utils.Map[string, string]
	StarRailRelics        *utils.Map[string, *starrail.RelicData]
	StarRailLightCones    *utils.Map[string, *starrail.LightConeData]

	excels HSRExcels
}

func (m *starRailData) Excels() HSRExcels {
	return m.excels
}

func (m *starRailData) CharacterData(uid string) *starrail.CharacterData {
	if data, ok := m.StarRailCharacterData.Get(uid); ok {
		data.Path = starrail.PathFromRaw(data.RawPath)
		return data
	}

	return nil
}

func (m *starRailData) Characters() []*starrail.CharacterData {
	s := make([]*starrail.CharacterData, 0, m.StarRailCharacterData.Len())
	m.StarRailCharacterData.ForEach(func(_ string, c *starrail.CharacterData) {
		s = append(s, c)
	})
	return s
}

func (m *starRailData) AvatarKey(id string) string {
	if avatar, ok := m.StarRailAvatars.Get(id); ok {
		return avatar
	}
	return id
}

func (m *starRailData) RelicData(id string) *starrail.RelicData {
	if relic, ok := m.StarRailRelics.Get(id); ok {
		return relic
	}
	return nil
}

func (m *starRailData) LightConeData(id string) *starrail.LightConeData {
	if lightCone, ok := m.StarRailLightCones.Get(id); ok {
		return lightCone
	}
	return nil
}
