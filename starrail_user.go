package enkanetworkapigo

type StarRailPlatform string

const (
	WINDOWS StarRailPlatform = "Windows"
	ANDROID                  = "Android"
	IOS                      = "IOS"
	PS4                      = "PS4"
	PS5                      = "PS5"
)

type StarRailUser struct {
	NickName string
	// Between 1 and 70
	Level            int
	Signature        string
	Uid              int
	FriendCount      int
	EquilibriumLevel int
	// The highest planet(?) the user has unlocked in the simulated universe
	SimulatedUniverse int
	Platform          StarRailPlatform
	Characters        []StarRailUserCharacter
}

func StarRailUserFromRaw(rawUser *RawHonkaiUser) *StarRailUser {
	var detailInfo RawDetailInfo = rawUser.DetailInfo
	var recordInfo RawRecordInfo = detailInfo.RecordInfo

	var user StarRailUser = StarRailUser{
		NickName:          detailInfo.NickName,
		Level:             detailInfo.Level,
		Signature:         detailInfo.Signature,
		Uid:               detailInfo.Uid,
		FriendCount:       detailInfo.FriendCount,
		EquilibriumLevel:  detailInfo.WordlLevel,
		SimulatedUniverse: recordInfo.MaxRogueChallengeScore,
		Platform:          StarRailPlatform(detailInfo.Platform),
	}

	if !detailInfo.IsDisplayAvatar {
		user.Characters = make([]StarRailUserCharacter, 0)
		return &user
	}

	var characters = []StarRailUserCharacter{}

	for _, avatar := range detailInfo.AvatarDetailList {

		var eidolon int = avatar.Rank
		var ascension int = avatar.Promotion
		var level int = avatar.Level
		var avatarId int = avatar.AvatarId
		var relics []StarRailRelic = []StarRailRelic{}
		var lightCone *StarRailLightCone = nil

		var equipment *RawEquipmentInfo = avatar.Equipment
		if equipment != nil {
			var stats []StarRailLightConeStat = []StarRailLightConeStat{}

			var flat RawEquipmentFlatData = equipment.EquipmentFlatData

			for _, subData := range flat.Props {
				stats = append(stats, StarRailLightConeStat{
					Stat:  subData.Type,
					Value: subData.Value,
				})
			}

			lightCone = &StarRailLightCone{
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
			var subStats []StarRailRelicStat = []StarRailRelicStat{}

			var props []RawRelicFlatProp = flat.Props
			if props == nil || len(props) == 0 {
				continue
			}

			for i, subData := range props {
				if i == 0 {
					// Main Stat is not a sub stats
					continue
				}
				subStats = append(subStats, StarRailRelicStat{
					Stat:  subData.Type,
					Value: subData.Value,
				})
			}

			var mainStat RawRelicFlatProp = props[0]
			var main StarRailRelicStat = StarRailRelicStat{
				Stat:  mainStat.Type,
				Value: mainStat.Value,
			}

			relics = append(relics, StarRailRelic{
				Level:    relicLevel,
				Hash:     flat.SetName,
				MainStat: main,
				SubStats: subStats,
				Type:     starRailRelicTypeFromId(relicData.Type),
			})
		}

		characters = append(characters, StarRailUserCharacter{
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
