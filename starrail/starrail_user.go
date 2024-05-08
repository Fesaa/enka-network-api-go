package starrail

import "fmt"

type Platform string

const (
	WINDOWS Platform = "Windows"
	ANDROID          = "Android"
	IOS              = "IOS"
	PS4              = "PS4"
	PS5              = "PS5"
)

type User struct {
	NickName string
	// Between 1 and 70
	Level     int
	Signature string
	// This has to be parsed with EnkaNetworkApi#GetStarRailAvatarKey & EnkaiNetworkApi#GetStarRailIcon
	AvatarIconId     string
	Uid              int
	FriendCount      int
	EquilibriumLevel int
	// The highest planet(?) the user has unlocked in the simulated universe
	SimulatedUniverse int
	Platform          Platform
	Characters        []UserCharacter
}

func UserFromRaw(rawUser *RawHonkaiUser) *User {
	var detailInfo RawDetailInfo = rawUser.DetailInfo
	var recordInfo RawRecordInfo = detailInfo.RecordInfo

	var user User = User{
		NickName:          detailInfo.NickName,
		Level:             detailInfo.Level,
		Signature:         detailInfo.Signature,
		AvatarIconId:      fmt.Sprint(detailInfo.HeadIcon),
		Uid:               detailInfo.Uid,
		FriendCount:       detailInfo.FriendCount,
		EquilibriumLevel:  detailInfo.WordlLevel,
		SimulatedUniverse: recordInfo.MaxRogueChallengeScore,
		Platform:          Platform(detailInfo.Platform),
	}

	if !detailInfo.IsDisplayAvatar {
		user.Characters = make([]UserCharacter, 0)
		return &user
	}

	var characters = []UserCharacter{}

	for _, avatar := range detailInfo.AvatarDetailList {

		var eidolon int = avatar.Rank
		var ascension int = avatar.Promotion
		var level int = avatar.Level
		var avatarId int = avatar.AvatarId
		var relics []Relic = []Relic{}
		var lightCone *LightCone = nil

		var equipment *RawEquipmentInfo = avatar.Equipment
		if equipment != nil {
			var stats []LightConeStat = []LightConeStat{}

			var flat RawEquipmentFlatData = equipment.EquipmentFlatData

			for _, subData := range flat.Props {
				stats = append(stats, LightConeStat{
					Stat:  subData.Type,
					Value: subData.Value,
				})
			}

			lightCone = &LightCone{
				SuperImposion: equipment.Rank,
				Promotion:     equipment.Promotion,
				Level:         equipment.Level,
				Stats:         stats,
				Hash:          flat.Name,
			}
		}

		for _, relicData := range avatar.RelicList {
			var relicLevel int = relicData.Level
			var flat RawRelicFlatData = relicData.RelicFlatData
			var subStats []RelicStat = []RelicStat{}

			var props []RawRelicFlatProp = flat.Props
			if props == nil || len(props) == 0 {
				continue
			}

			for i, subData := range props {
				if i == 0 {
					// Main Stat is not a sub stats
					continue
				}
				subStats = append(subStats, RelicStat{
					Stat:  subData.Type,
					Value: subData.Value,
				})
			}

			var mainStat RawRelicFlatProp = props[0]
			var main RelicStat = RelicStat{
				Stat:  mainStat.Type,
				Value: mainStat.Value,
			}

			relics = append(relics, Relic{
				RelicID:  relicData.Tid,
				Level:    relicLevel,
				Hash:     flat.SetName,
				MainStat: main,
				SubStats: subStats,
				Type:     starRailRelicTypeFromId(relicData.Type),
			})
		}

		characters = append(characters, UserCharacter{
			Eidolon:   eidolon,
			Ascension: ascension,
			Relics:    relics,
			Level:     level,
			AvatarId:  avatarId,
			LightCone: lightCone,
		})
	}

	user.Characters = characters
	return &user
}
