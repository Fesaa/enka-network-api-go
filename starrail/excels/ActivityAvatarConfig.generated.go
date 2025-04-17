package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivityAvatarConfig struct {
	AIPath                    string        `json:"AIPath"`
	ActionAvatarHeadIconPath  string        `json:"ActionAvatarHeadIconPath"`
	AdventurePlayerID         float64       `json:"AdventurePlayerID"`
	AssistBgOffset            []interface{} `json:"AssistBgOffset"`
	AssistOffset              []float64     `json:"AssistOffset"`
	AvatarBaseType            string        `json:"AvatarBaseType"`
	AvatarCutinBgImgPath      string        `json:"AvatarCutinBgImgPath"`
	AvatarCutinFrontImgPath   string        `json:"AvatarCutinFrontImgPath"`
	AvatarCutinImgPath        string        `json:"AvatarCutinImgPath"`
	AvatarCutinIntroText      hash.Hash     `json:"AvatarCutinIntroText"`
	AvatarDropOffset          []float64     `json:"AvatarDropOffset"`
	AvatarFullName            hash.Hash     `json:"AvatarFullName"`
	AvatarGachaResultImgPath  string        `json:"AvatarGachaResultImgPath"`
	AvatarID                  float64       `json:"AvatarID"`
	AvatarMiniIconPath        string        `json:"AvatarMiniIconPath"`
	AvatarName                hash.Hash     `json:"AvatarName"`
	AvatarSelfShowOffset      []interface{} `json:"AvatarSelfShowOffset"`
	AvatarSideIconPath        string        `json:"AvatarSideIconPath"`
	AvatarTrialOffset         []interface{} `json:"AvatarTrialOffset"`
	AvatarVOTag               string        `json:"AvatarVOTag"`
	DamageType                string        `json:"DamageType"`
	DamageTypeResistance      []interface{} `json:"DamageTypeResistance"`
	DefaultAvatarHeadIconPath string        `json:"DefaultAvatarHeadIconPath"`
	DefaultAvatarModelPath    string        `json:"DefaultAvatarModelPath"`
	ExpGroup                  float64       `json:"ExpGroup"`
	JsonPath                  string        `json:"JsonPath"`
	ManikinJsonPath           string        `json:"ManikinJsonPath"`
	MaxPromotion              float64       `json:"MaxPromotion"`
	PlayerCardOffset          []float64     `json:"PlayerCardOffset"`
	RankIDList                []float64     `json:"RankIDList"`
	Rarity                    string        `json:"Rarity"`
	Release                   bool          `json:"Release"`
	RewardList                []interface{} `json:"RewardList"`
	SPNeed                    hash.IntValue `json:"SPNeed"`
	SideAvatarHeadIconPath    string        `json:"SideAvatarHeadIconPath"`
	SkillList                 []float64     `json:"SkillList"`
	SkilltreePrefabPath       string        `json:"SkilltreePrefabPath"`
	UIAvatarModelPath         string        `json:"UIAvatarModelPath"`
	UltraSkillCutInPrefabPath string        `json:"UltraSkillCutInPrefabPath"`
	WaitingAvatarHeadIconPath string        `json:"WaitingAvatarHeadIconPath"`
}
type ActivityAvatarConfigAccessor struct {
	_data                          []ActivityAvatarConfig
	_dataAvatarGachaResultImgPath  map[string]ActivityAvatarConfig
	_dataUltraSkillCutInPrefabPath map[string]ActivityAvatarConfig
	_dataAvatarID                  map[float64]ActivityAvatarConfig
	_dataDefaultAvatarHeadIconPath map[string]ActivityAvatarConfig
	_dataAvatarCutinImgPath        map[string]ActivityAvatarConfig
	_dataAdventurePlayerID         map[float64]ActivityAvatarConfig
	_dataAvatarCutinFrontImgPath   map[string]ActivityAvatarConfig
	_dataJsonPath                  map[string]ActivityAvatarConfig
	_dataAvatarSideIconPath        map[string]ActivityAvatarConfig
	_dataAvatarMiniIconPath        map[string]ActivityAvatarConfig
	_dataSideAvatarHeadIconPath    map[string]ActivityAvatarConfig
	_dataActionAvatarHeadIconPath  map[string]ActivityAvatarConfig
	_dataWaitingAvatarHeadIconPath map[string]ActivityAvatarConfig
	_dataAvatarCutinBgImgPath      map[string]ActivityAvatarConfig
	_dataDefaultAvatarModelPath    map[string]ActivityAvatarConfig
}

// LoadData retrieves the data. Must be called before ActivityAvatarConfig.GroupData
func (a *ActivityAvatarConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityAvatarConfig.json")
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
func (a *ActivityAvatarConfigAccessor) Raw() ([]ActivityAvatarConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityAvatarConfigAccessor.LoadData to preload everything
func (a *ActivityAvatarConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataAvatarGachaResultImgPath[d.AvatarGachaResultImgPath] = d
		a._dataUltraSkillCutInPrefabPath[d.UltraSkillCutInPrefabPath] = d
		a._dataAvatarID[d.AvatarID] = d
		a._dataDefaultAvatarHeadIconPath[d.DefaultAvatarHeadIconPath] = d
		a._dataAvatarCutinImgPath[d.AvatarCutinImgPath] = d
		a._dataAdventurePlayerID[d.AdventurePlayerID] = d
		a._dataAvatarCutinFrontImgPath[d.AvatarCutinFrontImgPath] = d
		a._dataJsonPath[d.JsonPath] = d
		a._dataAvatarSideIconPath[d.AvatarSideIconPath] = d
		a._dataAvatarMiniIconPath[d.AvatarMiniIconPath] = d
		a._dataSideAvatarHeadIconPath[d.SideAvatarHeadIconPath] = d
		a._dataActionAvatarHeadIconPath[d.ActionAvatarHeadIconPath] = d
		a._dataWaitingAvatarHeadIconPath[d.WaitingAvatarHeadIconPath] = d
		a._dataAvatarCutinBgImgPath[d.AvatarCutinBgImgPath] = d
		a._dataDefaultAvatarModelPath[d.DefaultAvatarModelPath] = d
	}
}

// ByAvatarGachaResultImgPath returns the ActivityAvatarConfig uniquely identified by AvatarGachaResultImgPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByAvatarGachaResultImgPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataAvatarGachaResultImgPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataAvatarGachaResultImgPath[identifier], nil
}

// ByUltraSkillCutInPrefabPath returns the ActivityAvatarConfig uniquely identified by UltraSkillCutInPrefabPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByUltraSkillCutInPrefabPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataUltraSkillCutInPrefabPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataUltraSkillCutInPrefabPath[identifier], nil
}

// ByAvatarID returns the ActivityAvatarConfig uniquely identified by AvatarID
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByAvatarID(identifier float64) (ActivityAvatarConfig, error) {
	if a._dataAvatarID == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataAvatarID[identifier], nil
}

// ByDefaultAvatarHeadIconPath returns the ActivityAvatarConfig uniquely identified by DefaultAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByDefaultAvatarHeadIconPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataDefaultAvatarHeadIconPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataDefaultAvatarHeadIconPath[identifier], nil
}

// ByAvatarCutinImgPath returns the ActivityAvatarConfig uniquely identified by AvatarCutinImgPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByAvatarCutinImgPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataAvatarCutinImgPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataAvatarCutinImgPath[identifier], nil
}

// ByAdventurePlayerID returns the ActivityAvatarConfig uniquely identified by AdventurePlayerID
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByAdventurePlayerID(identifier float64) (ActivityAvatarConfig, error) {
	if a._dataAdventurePlayerID == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataAdventurePlayerID[identifier], nil
}

// ByAvatarCutinFrontImgPath returns the ActivityAvatarConfig uniquely identified by AvatarCutinFrontImgPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByAvatarCutinFrontImgPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataAvatarCutinFrontImgPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataAvatarCutinFrontImgPath[identifier], nil
}

// ByJsonPath returns the ActivityAvatarConfig uniquely identified by JsonPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByJsonPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataJsonPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataJsonPath[identifier], nil
}

// ByAvatarSideIconPath returns the ActivityAvatarConfig uniquely identified by AvatarSideIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByAvatarSideIconPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataAvatarSideIconPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataAvatarSideIconPath[identifier], nil
}

// ByAvatarMiniIconPath returns the ActivityAvatarConfig uniquely identified by AvatarMiniIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByAvatarMiniIconPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataAvatarMiniIconPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataAvatarMiniIconPath[identifier], nil
}

// BySideAvatarHeadIconPath returns the ActivityAvatarConfig uniquely identified by SideAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) BySideAvatarHeadIconPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataSideAvatarHeadIconPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataSideAvatarHeadIconPath[identifier], nil
}

// ByActionAvatarHeadIconPath returns the ActivityAvatarConfig uniquely identified by ActionAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByActionAvatarHeadIconPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataActionAvatarHeadIconPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataActionAvatarHeadIconPath[identifier], nil
}

// ByWaitingAvatarHeadIconPath returns the ActivityAvatarConfig uniquely identified by WaitingAvatarHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByWaitingAvatarHeadIconPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataWaitingAvatarHeadIconPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataWaitingAvatarHeadIconPath[identifier], nil
}

// ByAvatarCutinBgImgPath returns the ActivityAvatarConfig uniquely identified by AvatarCutinBgImgPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByAvatarCutinBgImgPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataAvatarCutinBgImgPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataAvatarCutinBgImgPath[identifier], nil
}

// ByDefaultAvatarModelPath returns the ActivityAvatarConfig uniquely identified by DefaultAvatarModelPath
//
// Error is only non-nil if the source errors out
func (a *ActivityAvatarConfigAccessor) ByDefaultAvatarModelPath(identifier string) (ActivityAvatarConfig, error) {
	if a._dataDefaultAvatarModelPath == nil {
		err := a.LoadData()
		if err != nil {
			return ActivityAvatarConfig{}, err
		}
		a.GroupData()
	}
	return a._dataDefaultAvatarModelPath[identifier], nil
}
