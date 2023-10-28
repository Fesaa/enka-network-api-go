package genshin

import (
	"errors"
	"strconv"
)

type UserCharacter struct {
	Id              int
	Constellation   int
	Experience      int
	Ascension       int
	Level           int
	FriendshipLevel int
	FightProperties map[FightProp]float64
	Artifacts       []Artifact
	Weapon          *Weapon
}

func userCharacterFromRaw(c RawAvatarInfo) UserCharacter {
	id := c.AvatarId
	var constellation int
	if c.TalentIdList == nil {
		constellation = 0
	} else {
		constellation = len(c.TalentIdList)
	}

	var fightProperties map[FightProp]float64 = make(map[FightProp]float64)
	var artifacts []Artifact = make([]Artifact, 0)
	var weapon *Weapon = nil

	for id, value := range c.FightPropMap {
		fightProp := fromId(id)
		if fightProp == nil {
			continue
		}
		fightProperties[*fightProp] = value
	}

	for _, equipmentData := range c.EquipList {
		artifactData := equipmentData.ArtifactData
		flatData := equipmentData.Flat
		if artifactData == nil {
			// ArtifacrData is nil if the equipment is a weapon
			temp := fromRawData(equipmentData, flatData)
			weapon = &temp
			continue
		}

		mainStat := flatData.ArtifactMainData
		subStats := make([]ArtifactStat, len(flatData.ReliquarySubStats))
		for i, stat := range flatData.ReliquarySubStats {
			subStats[i] = ArtifactStat{
				Stat:  stat.AppendPropId,
				Value: float64(stat.StatValue),
			}
		}

		artifactType := TypeFromId(flatData.EquipType)
		if artifactType == -1 {
			panic(errors.New("You tried to do an action that is not allowed on this library version. Either update or find a different way."))
		}

		artifacts = append(artifacts, Artifact{
			ArtifactType: artifactType,
			Level:        artifactData.Level - 1,
			MainStats: ArtifactStat{
				Stat:  mainStat.MainPropId,
				Value: float64(mainStat.StatValue),
			},
			SubStats:    subStats,
			SetNameHash: flatData.SetNameTextMapHash,
		})
	}

	levelMap := c.PropMap["4001"].(map[string]interface{})
	ascensionmap := c.PropMap["1002"].(map[string]interface{})
	experienceMap := c.PropMap["1001"].(map[string]interface{})

	var level int
	var ascension int
	var experience int

	if lvl, ok := levelMap["val"]; ok {
		level, _ = strconv.Atoi(lvl.(string))
	} else {
		level = 1
	}

	if asc, ok := ascensionmap["val"]; ok {
		ascension, _ = strconv.Atoi(asc.(string))
	} else {
		ascension = 1
	}

	if exp, ok := experienceMap["val"]; ok {
		experience, _ = strconv.Atoi(exp.(string))
	} else {
		experience = -1
	}

	return UserCharacter{
		Id:              id,
		Level:           level,
		Constellation:   constellation,
		Experience:      experience,
		Ascension:       ascension,
		FightProperties: fightProperties,
		Artifacts:       artifacts,
		FriendshipLevel: c.FetterInfo.ExpLevel,
		Weapon:          weapon,
	}
}
