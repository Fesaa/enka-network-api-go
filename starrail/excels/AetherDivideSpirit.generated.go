package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type AetherDivideSpirit struct {
	AIPath                    string                 `json:"AIPath"`
	ActionAvatarHeadIconPath  string                 `json:"ActionAvatarHeadIconPath"`
	AtlasAvatarHeadIconPath   string                 `json:"AtlasAvatarHeadIconPath"`
	AvatarID                  json.Number            `json:"AvatarID"`
	AvatarName                map[string]json.Number `json:"AvatarName"`
	AvatarSideIconPath        string                 `json:"AvatarSideIconPath"`
	AvatarVOTag               string                 `json:"AvatarVOTag"`
	DamageType                string                 `json:"DamageType"`
	DefaultAvatarHeadIconPath string                 `json:"DefaultAvatarHeadIconPath"`
	DefaultAvatarModelPath    string                 `json:"DefaultAvatarModelPath"`
	ExpItemID                 json.Number            `json:"ExpItemID"`
	GymLocation               json.Number            `json:"GymLocation"`
	JsonPath                  string                 `json:"JsonPath"`
	ManikinAvatarModelPath    string                 `json:"ManikinAvatarModelPath"`
	ManikinJsonPath           string                 `json:"ManikinJsonPath"`
	MaxPromotion              json.Number            `json:"MaxPromotion"`
	MiddleAvatarHeadIconPath  string                 `json:"MiddleAvatarHeadIconPath"`
	PassiveSkillSlotList      []string               `json:"PassiveSkillSlotList"`
	Rarity                    string                 `json:"Rarity"`
	RecommendPassiveSkillList []json.Number          `json:"RecommendPassiveSkillList"`
	SPMax                     map[string]json.Number `json:"SPMax"`
	SideAvatarHeadIconPath    string                 `json:"SideAvatarHeadIconPath"`
	SkillList                 []json.Number          `json:"SkillList"`
	SpiritDescription         map[string]json.Number `json:"SpiritDescription"`
	SpiritType                string                 `json:"SpiritType"`
	SpiritUnlockDescription   map[string]json.Number `json:"SpiritUnlockDescription"`
	TeamLeftPrefabPath        string                 `json:"TeamLeftPrefabPath"`
	TeamRightPrefabPath       string                 `json:"TeamRightPrefabPath"`
	UltraSkillCutInPrefabPath string                 `json:"UltraSkillCutInPrefabPath"`
	WaitingAvatarHeadIconPath string                 `json:"WaitingAvatarHeadIconPath"`
}
type AetherDivideSpiritAccessor struct {
	_data                          []AetherDivideSpirit
	_dataActionAvatarHeadIconPath  map[string]AetherDivideSpirit
	_dataAtlasAvatarHeadIconPath   map[string]AetherDivideSpirit
	_dataAvatarID                  map[json.Number]AetherDivideSpirit
	_dataAvatarSideIconPath        map[string]AetherDivideSpirit
	_dataDefaultAvatarHeadIconPath map[string]AetherDivideSpirit
	_dataDefaultAvatarModelPath    map[string]AetherDivideSpirit
	_dataExpItemID                 map[json.Number]AetherDivideSpirit
	_dataJsonPath                  map[string]AetherDivideSpirit
	_dataManikinAvatarModelPath    map[string]AetherDivideSpirit
	_dataMiddleAvatarHeadIconPath  map[string]AetherDivideSpirit
	_dataSideAvatarHeadIconPath    map[string]AetherDivideSpirit
	_dataTeamLeftPrefabPath        map[string]AetherDivideSpirit
	_dataTeamRightPrefabPath       map[string]AetherDivideSpirit
	_dataUltraSkillCutInPrefabPath map[string]AetherDivideSpirit
	_dataWaitingAvatarHeadIconPath map[string]AetherDivideSpirit
}

// LoadData retrieves the data. Must be called before AetherDivideSpirit.GroupData
func (a *AetherDivideSpiritAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AetherDivideSpirit.json")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &a._data)
}

// Raw returns the raw data.
func (a *AetherDivideSpiritAccessor) Raw() ([]AetherDivideSpirit, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AetherDivideSpirit{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AetherDivideSpiritAccessor.LoadData to preload everything
func (a *AetherDivideSpiritAccessor) GroupData() {
	a._dataActionAvatarHeadIconPath = map[string]AetherDivideSpirit{}
	a._dataAtlasAvatarHeadIconPath = map[string]AetherDivideSpirit{}
	a._dataAvatarID = map[json.Number]AetherDivideSpirit{}
	a._dataAvatarSideIconPath = map[string]AetherDivideSpirit{}
	a._dataDefaultAvatarHeadIconPath = map[string]AetherDivideSpirit{}
	a._dataDefaultAvatarModelPath = map[string]AetherDivideSpirit{}
	a._dataExpItemID = map[json.Number]AetherDivideSpirit{}
	a._dataJsonPath = map[string]AetherDivideSpirit{}
	a._dataManikinAvatarModelPath = map[string]AetherDivideSpirit{}
	a._dataMiddleAvatarHeadIconPath = map[string]AetherDivideSpirit{}
	a._dataSideAvatarHeadIconPath = map[string]AetherDivideSpirit{}
	a._dataTeamLeftPrefabPath = map[string]AetherDivideSpirit{}
	a._dataTeamRightPrefabPath = map[string]AetherDivideSpirit{}
	a._dataUltraSkillCutInPrefabPath = map[string]AetherDivideSpirit{}
	a._dataWaitingAvatarHeadIconPath = map[string]AetherDivideSpirit{}
	for _, d := range a._data {
		a._dataActionAvatarHeadIconPath[d.ActionAvatarHeadIconPath] = d
		a._dataAtlasAvatarHeadIconPath[d.AtlasAvatarHeadIconPath] = d
		a._dataAvatarID[d.AvatarID] = d
		a._dataAvatarSideIconPath[d.AvatarSideIconPath] = d
		a._dataDefaultAvatarHeadIconPath[d.DefaultAvatarHeadIconPath] = d
		a._dataDefaultAvatarModelPath[d.DefaultAvatarModelPath] = d
		a._dataExpItemID[d.ExpItemID] = d
		a._dataJsonPath[d.JsonPath] = d
		a._dataManikinAvatarModelPath[d.ManikinAvatarModelPath] = d
		a._dataMiddleAvatarHeadIconPath[d.MiddleAvatarHeadIconPath] = d
		a._dataSideAvatarHeadIconPath[d.SideAvatarHeadIconPath] = d
		a._dataTeamLeftPrefabPath[d.TeamLeftPrefabPath] = d
		a._dataTeamRightPrefabPath[d.TeamRightPrefabPath] = d
		a._dataUltraSkillCutInPrefabPath[d.UltraSkillCutInPrefabPath] = d
		a._dataWaitingAvatarHeadIconPath[d.WaitingAvatarHeadIconPath] = d
	}
}

// ByActionAvatarHeadIconPath returns the AetherDivideSpirit uniquely identified by ActionAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByActionAvatarHeadIconPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataActionAvatarHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActionAvatarHeadIconPath[identifier], nil
}

// ByAtlasAvatarHeadIconPath returns the AetherDivideSpirit uniquely identified by AtlasAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByAtlasAvatarHeadIconPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataAtlasAvatarHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAtlasAvatarHeadIconPath[identifier], nil
}

// ByAvatarID returns the AetherDivideSpirit uniquely identified by AvatarID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByAvatarID(identifier json.Number) (AetherDivideSpirit, error) {
	if a._dataAvatarID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarID[identifier], nil
}

// ByAvatarSideIconPath returns the AetherDivideSpirit uniquely identified by AvatarSideIconPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByAvatarSideIconPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataAvatarSideIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarSideIconPath[identifier], nil
}

// ByDefaultAvatarHeadIconPath returns the AetherDivideSpirit uniquely identified by DefaultAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByDefaultAvatarHeadIconPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataDefaultAvatarHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDefaultAvatarHeadIconPath[identifier], nil
}

// ByDefaultAvatarModelPath returns the AetherDivideSpirit uniquely identified by DefaultAvatarModelPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByDefaultAvatarModelPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataDefaultAvatarModelPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDefaultAvatarModelPath[identifier], nil
}

// ByExpItemID returns the AetherDivideSpirit uniquely identified by ExpItemID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByExpItemID(identifier json.Number) (AetherDivideSpirit, error) {
	if a._dataExpItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataExpItemID[identifier], nil
}

// ByJsonPath returns the AetherDivideSpirit uniquely identified by JsonPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByJsonPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataJsonPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataJsonPath[identifier], nil
}

// ByManikinAvatarModelPath returns the AetherDivideSpirit uniquely identified by ManikinAvatarModelPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByManikinAvatarModelPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataManikinAvatarModelPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataManikinAvatarModelPath[identifier], nil
}

// ByMiddleAvatarHeadIconPath returns the AetherDivideSpirit uniquely identified by MiddleAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByMiddleAvatarHeadIconPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataMiddleAvatarHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMiddleAvatarHeadIconPath[identifier], nil
}

// BySideAvatarHeadIconPath returns the AetherDivideSpirit uniquely identified by SideAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) BySideAvatarHeadIconPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataSideAvatarHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSideAvatarHeadIconPath[identifier], nil
}

// ByTeamLeftPrefabPath returns the AetherDivideSpirit uniquely identified by TeamLeftPrefabPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByTeamLeftPrefabPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataTeamLeftPrefabPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTeamLeftPrefabPath[identifier], nil
}

// ByTeamRightPrefabPath returns the AetherDivideSpirit uniquely identified by TeamRightPrefabPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByTeamRightPrefabPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataTeamRightPrefabPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTeamRightPrefabPath[identifier], nil
}

// ByUltraSkillCutInPrefabPath returns the AetherDivideSpirit uniquely identified by UltraSkillCutInPrefabPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByUltraSkillCutInPrefabPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataUltraSkillCutInPrefabPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUltraSkillCutInPrefabPath[identifier], nil
}

// ByWaitingAvatarHeadIconPath returns the AetherDivideSpirit uniquely identified by WaitingAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AetherDivideSpiritAccessor) ByWaitingAvatarHeadIconPath(identifier string) (AetherDivideSpirit, error) {
	if a._dataWaitingAvatarHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideSpirit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataWaitingAvatarHeadIconPath[identifier], nil
}
