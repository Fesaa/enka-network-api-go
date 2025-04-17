package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type AvatarConfigTrial struct {
	AIPath                    string                   `json:"AIPath"`
	ActionAvatarHeadIconPath  string                   `json:"ActionAvatarHeadIconPath"`
	AdventurePlayerID         json.Number              `json:"AdventurePlayerID"`
	AssistBgOffset            []json.Number            `json:"AssistBgOffset"`
	AssistOffset              []json.Number            `json:"AssistOffset"`
	AvatarBaseType            string                   `json:"AvatarBaseType"`
	AvatarCutinBgImgPath      string                   `json:"AvatarCutinBgImgPath"`
	AvatarCutinFrontImgPath   string                   `json:"AvatarCutinFrontImgPath"`
	AvatarCutinImgPath        string                   `json:"AvatarCutinImgPath"`
	AvatarCutinIntroText      map[string]json.Number   `json:"AvatarCutinIntroText"`
	AvatarDropOffset          []json.Number            `json:"AvatarDropOffset"`
	AvatarFullName            map[string]json.Number   `json:"AvatarFullName"`
	AvatarGachaResultImgPath  string                   `json:"AvatarGachaResultImgPath"`
	AvatarID                  json.Number              `json:"AvatarID"`
	AvatarMiniIconPath        string                   `json:"AvatarMiniIconPath"`
	AvatarName                map[string]json.Number   `json:"AvatarName"`
	AvatarSelfShowOffset      []json.Number            `json:"AvatarSelfShowOffset"`
	AvatarSideIconPath        string                   `json:"AvatarSideIconPath"`
	AvatarTrialOffset         []json.Number            `json:"AvatarTrialOffset"`
	AvatarVOTag               string                   `json:"AvatarVOTag"`
	DamageType                string                   `json:"DamageType"`
	DamageTypeResistance      []interface{}            `json:"DamageTypeResistance"`
	DefaultAvatarHeadIconPath string                   `json:"DefaultAvatarHeadIconPath"`
	DefaultAvatarModelPath    string                   `json:"DefaultAvatarModelPath"`
	ExpGroup                  json.Number              `json:"ExpGroup"`
	JsonPath                  string                   `json:"JsonPath"`
	ManikinJsonPath           string                   `json:"ManikinJsonPath"`
	MaxPromotion              json.Number              `json:"MaxPromotion"`
	MaxRank                   json.Number              `json:"MaxRank"`
	PlayerCardOffset          []json.Number            `json:"PlayerCardOffset"`
	RankIDList                []json.Number            `json:"RankIDList"`
	Rarity                    string                   `json:"Rarity"`
	Release                   bool                     `json:"Release"`
	RewardList                []map[string]json.Number `json:"RewardList"`
	SPNeed                    map[string]json.Number   `json:"SPNeed"`
	SideAvatarHeadIconPath    string                   `json:"SideAvatarHeadIconPath"`
	SkillList                 []json.Number            `json:"SkillList"`
	SkilltreePrefabPath       string                   `json:"SkilltreePrefabPath"`
	UIAvatarModelPath         string                   `json:"UIAvatarModelPath"`
	UltraSkillCutInPrefabPath string                   `json:"UltraSkillCutInPrefabPath"`
	WaitingAvatarHeadIconPath string                   `json:"WaitingAvatarHeadIconPath"`
}
type AvatarConfigTrialAccessor struct {
	_data                          []AvatarConfigTrial
	_dataActionAvatarHeadIconPath  map[string]AvatarConfigTrial
	_dataAvatarCutinBgImgPath      map[string]AvatarConfigTrial
	_dataAvatarCutinFrontImgPath   map[string]AvatarConfigTrial
	_dataAvatarCutinImgPath        map[string]AvatarConfigTrial
	_dataAvatarGachaResultImgPath  map[string]AvatarConfigTrial
	_dataAvatarMiniIconPath        map[string]AvatarConfigTrial
	_dataAvatarSideIconPath        map[string]AvatarConfigTrial
	_dataAvatarVOTag               map[string]AvatarConfigTrial
	_dataDamageType                map[string]AvatarConfigTrial
	_dataDefaultAvatarHeadIconPath map[string]AvatarConfigTrial
	_dataDefaultAvatarModelPath    map[string]AvatarConfigTrial
	_dataJsonPath                  map[string]AvatarConfigTrial
	_dataManikinJsonPath           map[string]AvatarConfigTrial
	_dataSideAvatarHeadIconPath    map[string]AvatarConfigTrial
	_dataUIAvatarModelPath         map[string]AvatarConfigTrial
	_dataUltraSkillCutInPrefabPath map[string]AvatarConfigTrial
	_dataWaitingAvatarHeadIconPath map[string]AvatarConfigTrial
}

// LoadData retrieves the data. Must be called before AvatarConfigTrial.GroupData
func (a *AvatarConfigTrialAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarConfigTrial.json")
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
func (a *AvatarConfigTrialAccessor) Raw() ([]AvatarConfigTrial, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarConfigTrial{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarConfigTrialAccessor.LoadData to preload everything
func (a *AvatarConfigTrialAccessor) GroupData() {
	a._dataActionAvatarHeadIconPath = map[string]AvatarConfigTrial{}
	a._dataAvatarCutinBgImgPath = map[string]AvatarConfigTrial{}
	a._dataAvatarCutinFrontImgPath = map[string]AvatarConfigTrial{}
	a._dataAvatarCutinImgPath = map[string]AvatarConfigTrial{}
	a._dataAvatarGachaResultImgPath = map[string]AvatarConfigTrial{}
	a._dataAvatarMiniIconPath = map[string]AvatarConfigTrial{}
	a._dataAvatarSideIconPath = map[string]AvatarConfigTrial{}
	a._dataAvatarVOTag = map[string]AvatarConfigTrial{}
	a._dataDamageType = map[string]AvatarConfigTrial{}
	a._dataDefaultAvatarHeadIconPath = map[string]AvatarConfigTrial{}
	a._dataDefaultAvatarModelPath = map[string]AvatarConfigTrial{}
	a._dataJsonPath = map[string]AvatarConfigTrial{}
	a._dataManikinJsonPath = map[string]AvatarConfigTrial{}
	a._dataSideAvatarHeadIconPath = map[string]AvatarConfigTrial{}
	a._dataUIAvatarModelPath = map[string]AvatarConfigTrial{}
	a._dataUltraSkillCutInPrefabPath = map[string]AvatarConfigTrial{}
	a._dataWaitingAvatarHeadIconPath = map[string]AvatarConfigTrial{}
	for _, d := range a._data {
		a._dataActionAvatarHeadIconPath[d.ActionAvatarHeadIconPath] = d
		a._dataAvatarCutinBgImgPath[d.AvatarCutinBgImgPath] = d
		a._dataAvatarCutinFrontImgPath[d.AvatarCutinFrontImgPath] = d
		a._dataAvatarCutinImgPath[d.AvatarCutinImgPath] = d
		a._dataAvatarGachaResultImgPath[d.AvatarGachaResultImgPath] = d
		a._dataAvatarMiniIconPath[d.AvatarMiniIconPath] = d
		a._dataAvatarSideIconPath[d.AvatarSideIconPath] = d
		a._dataAvatarVOTag[d.AvatarVOTag] = d
		a._dataDamageType[d.DamageType] = d
		a._dataDefaultAvatarHeadIconPath[d.DefaultAvatarHeadIconPath] = d
		a._dataDefaultAvatarModelPath[d.DefaultAvatarModelPath] = d
		a._dataJsonPath[d.JsonPath] = d
		a._dataManikinJsonPath[d.ManikinJsonPath] = d
		a._dataSideAvatarHeadIconPath[d.SideAvatarHeadIconPath] = d
		a._dataUIAvatarModelPath[d.UIAvatarModelPath] = d
		a._dataUltraSkillCutInPrefabPath[d.UltraSkillCutInPrefabPath] = d
		a._dataWaitingAvatarHeadIconPath[d.WaitingAvatarHeadIconPath] = d
	}
}

// ByActionAvatarHeadIconPath returns the AvatarConfigTrial uniquely identified by ActionAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByActionAvatarHeadIconPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataActionAvatarHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActionAvatarHeadIconPath[identifier], nil
}

// ByAvatarCutinBgImgPath returns the AvatarConfigTrial uniquely identified by AvatarCutinBgImgPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByAvatarCutinBgImgPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataAvatarCutinBgImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarCutinBgImgPath[identifier], nil
}

// ByAvatarCutinFrontImgPath returns the AvatarConfigTrial uniquely identified by AvatarCutinFrontImgPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByAvatarCutinFrontImgPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataAvatarCutinFrontImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarCutinFrontImgPath[identifier], nil
}

// ByAvatarCutinImgPath returns the AvatarConfigTrial uniquely identified by AvatarCutinImgPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByAvatarCutinImgPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataAvatarCutinImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarCutinImgPath[identifier], nil
}

// ByAvatarGachaResultImgPath returns the AvatarConfigTrial uniquely identified by AvatarGachaResultImgPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByAvatarGachaResultImgPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataAvatarGachaResultImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarGachaResultImgPath[identifier], nil
}

// ByAvatarMiniIconPath returns the AvatarConfigTrial uniquely identified by AvatarMiniIconPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByAvatarMiniIconPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataAvatarMiniIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarMiniIconPath[identifier], nil
}

// ByAvatarSideIconPath returns the AvatarConfigTrial uniquely identified by AvatarSideIconPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByAvatarSideIconPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataAvatarSideIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarSideIconPath[identifier], nil
}

// ByAvatarVOTag returns the AvatarConfigTrial uniquely identified by AvatarVOTag
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByAvatarVOTag(identifier string) (AvatarConfigTrial, error) {
	if a._dataAvatarVOTag == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarVOTag[identifier], nil
}

// ByDamageType returns the AvatarConfigTrial uniquely identified by DamageType
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByDamageType(identifier string) (AvatarConfigTrial, error) {
	if a._dataDamageType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDamageType[identifier], nil
}

// ByDefaultAvatarHeadIconPath returns the AvatarConfigTrial uniquely identified by DefaultAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByDefaultAvatarHeadIconPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataDefaultAvatarHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDefaultAvatarHeadIconPath[identifier], nil
}

// ByDefaultAvatarModelPath returns the AvatarConfigTrial uniquely identified by DefaultAvatarModelPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByDefaultAvatarModelPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataDefaultAvatarModelPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDefaultAvatarModelPath[identifier], nil
}

// ByJsonPath returns the AvatarConfigTrial uniquely identified by JsonPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByJsonPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataJsonPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataJsonPath[identifier], nil
}

// ByManikinJsonPath returns the AvatarConfigTrial uniquely identified by ManikinJsonPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByManikinJsonPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataManikinJsonPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataManikinJsonPath[identifier], nil
}

// BySideAvatarHeadIconPath returns the AvatarConfigTrial uniquely identified by SideAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) BySideAvatarHeadIconPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataSideAvatarHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSideAvatarHeadIconPath[identifier], nil
}

// ByUIAvatarModelPath returns the AvatarConfigTrial uniquely identified by UIAvatarModelPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByUIAvatarModelPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataUIAvatarModelPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUIAvatarModelPath[identifier], nil
}

// ByUltraSkillCutInPrefabPath returns the AvatarConfigTrial uniquely identified by UltraSkillCutInPrefabPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByUltraSkillCutInPrefabPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataUltraSkillCutInPrefabPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUltraSkillCutInPrefabPath[identifier], nil
}

// ByWaitingAvatarHeadIconPath returns the AvatarConfigTrial uniquely identified by WaitingAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AvatarConfigTrialAccessor) ByWaitingAvatarHeadIconPath(identifier string) (AvatarConfigTrial, error) {
	if a._dataWaitingAvatarHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarConfigTrial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataWaitingAvatarHeadIconPath[identifier], nil
}
