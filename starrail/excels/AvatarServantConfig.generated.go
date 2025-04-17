package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AvatarServantConfig struct {
	ActionServantHeadIconPath  string        `json:"ActionServantHeadIconPath"`
	Aggro                      hash.IntValue `json:"Aggro"`
	Config                     string        `json:"Config"`
	HPBase                     string        `json:"HPBase"`
	HPInherit                  string        `json:"HPInherit"`
	HPSkill                    float64       `json:"HPSkill"`
	HeadIcon                   string        `json:"HeadIcon"`
	ManikinJsonPath            string        `json:"ManikinJsonPath"`
	Prefab                     string        `json:"Prefab"`
	ServantID                  float64       `json:"ServantID"`
	ServantMiniIconPath        string        `json:"ServantMiniIconPath"`
	ServantName                hash.Hash     `json:"ServantName"`
	ServantSideIconPath        string        `json:"ServantSideIconPath"`
	SkillIDList                []float64     `json:"SkillIDList"`
	SpeedBase                  string        `json:"SpeedBase"`
	SpeedInherit               string        `json:"SpeedInherit"`
	SpeedSkill                 float64       `json:"SpeedSkill"`
	UIServantModelPath         string        `json:"UIServantModelPath"`
	UnCreateHeadIconPath       string        `json:"UnCreateHeadIconPath"`
	WaitingServantHeadIconPath string        `json:"WaitingServantHeadIconPath"`
}
type AvatarServantConfigAccessor struct {
	_data                           []AvatarServantConfig
	_dataActionServantHeadIconPath  map[string]AvatarServantConfig
	_dataConfig                     map[string]AvatarServantConfig
	_dataHPBase                     map[string]AvatarServantConfig
	_dataHPInherit                  map[string]AvatarServantConfig
	_dataHeadIcon                   map[string]AvatarServantConfig
	_dataManikinJsonPath            map[string]AvatarServantConfig
	_dataPrefab                     map[string]AvatarServantConfig
	_dataServantID                  map[float64]AvatarServantConfig
	_dataServantMiniIconPath        map[string]AvatarServantConfig
	_dataServantSideIconPath        map[string]AvatarServantConfig
	_dataSpeedSkill                 map[float64]AvatarServantConfig
	_dataUIServantModelPath         map[string]AvatarServantConfig
	_dataUnCreateHeadIconPath       map[string]AvatarServantConfig
	_dataWaitingServantHeadIconPath map[string]AvatarServantConfig
}

// LoadData retrieves the data. Must be called before AvatarServantConfig.GroupData
func (a *AvatarServantConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarServantConfig.json")
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
func (a *AvatarServantConfigAccessor) Raw() ([]AvatarServantConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarServantConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarServantConfigAccessor.LoadData to preload everything
func (a *AvatarServantConfigAccessor) GroupData() {
	a._dataActionServantHeadIconPath = map[string]AvatarServantConfig{}
	a._dataConfig = map[string]AvatarServantConfig{}
	a._dataHPBase = map[string]AvatarServantConfig{}
	a._dataHPInherit = map[string]AvatarServantConfig{}
	a._dataHeadIcon = map[string]AvatarServantConfig{}
	a._dataManikinJsonPath = map[string]AvatarServantConfig{}
	a._dataPrefab = map[string]AvatarServantConfig{}
	a._dataServantID = map[float64]AvatarServantConfig{}
	a._dataServantMiniIconPath = map[string]AvatarServantConfig{}
	a._dataServantSideIconPath = map[string]AvatarServantConfig{}
	a._dataSpeedSkill = map[float64]AvatarServantConfig{}
	a._dataUIServantModelPath = map[string]AvatarServantConfig{}
	a._dataUnCreateHeadIconPath = map[string]AvatarServantConfig{}
	a._dataWaitingServantHeadIconPath = map[string]AvatarServantConfig{}
	for _, d := range a._data {
		a._dataActionServantHeadIconPath[d.ActionServantHeadIconPath] = d
		a._dataConfig[d.Config] = d
		a._dataHPBase[d.HPBase] = d
		a._dataHPInherit[d.HPInherit] = d
		a._dataHeadIcon[d.HeadIcon] = d
		a._dataManikinJsonPath[d.ManikinJsonPath] = d
		a._dataPrefab[d.Prefab] = d
		a._dataServantID[d.ServantID] = d
		a._dataServantMiniIconPath[d.ServantMiniIconPath] = d
		a._dataServantSideIconPath[d.ServantSideIconPath] = d
		a._dataSpeedSkill[d.SpeedSkill] = d
		a._dataUIServantModelPath[d.UIServantModelPath] = d
		a._dataUnCreateHeadIconPath[d.UnCreateHeadIconPath] = d
		a._dataWaitingServantHeadIconPath[d.WaitingServantHeadIconPath] = d
	}
}

// ByActionServantHeadIconPath returns the AvatarServantConfig uniquely identified by ActionServantHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByActionServantHeadIconPath(identifier string) (AvatarServantConfig, error) {
	if a._dataActionServantHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActionServantHeadIconPath[identifier], nil
}

// ByConfig returns the AvatarServantConfig uniquely identified by Config
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByConfig(identifier string) (AvatarServantConfig, error) {
	if a._dataConfig == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConfig[identifier], nil
}

// ByHPBase returns the AvatarServantConfig uniquely identified by HPBase
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByHPBase(identifier string) (AvatarServantConfig, error) {
	if a._dataHPBase == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataHPBase[identifier], nil
}

// ByHPInherit returns the AvatarServantConfig uniquely identified by HPInherit
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByHPInherit(identifier string) (AvatarServantConfig, error) {
	if a._dataHPInherit == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataHPInherit[identifier], nil
}

// ByHeadIcon returns the AvatarServantConfig uniquely identified by HeadIcon
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByHeadIcon(identifier string) (AvatarServantConfig, error) {
	if a._dataHeadIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataHeadIcon[identifier], nil
}

// ByManikinJsonPath returns the AvatarServantConfig uniquely identified by ManikinJsonPath
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByManikinJsonPath(identifier string) (AvatarServantConfig, error) {
	if a._dataManikinJsonPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataManikinJsonPath[identifier], nil
}

// ByPrefab returns the AvatarServantConfig uniquely identified by Prefab
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByPrefab(identifier string) (AvatarServantConfig, error) {
	if a._dataPrefab == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPrefab[identifier], nil
}

// ByServantID returns the AvatarServantConfig uniquely identified by ServantID
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByServantID(identifier float64) (AvatarServantConfig, error) {
	if a._dataServantID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataServantID[identifier], nil
}

// ByServantMiniIconPath returns the AvatarServantConfig uniquely identified by ServantMiniIconPath
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByServantMiniIconPath(identifier string) (AvatarServantConfig, error) {
	if a._dataServantMiniIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataServantMiniIconPath[identifier], nil
}

// ByServantSideIconPath returns the AvatarServantConfig uniquely identified by ServantSideIconPath
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByServantSideIconPath(identifier string) (AvatarServantConfig, error) {
	if a._dataServantSideIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataServantSideIconPath[identifier], nil
}

// BySpeedSkill returns the AvatarServantConfig uniquely identified by SpeedSkill
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) BySpeedSkill(identifier float64) (AvatarServantConfig, error) {
	if a._dataSpeedSkill == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSpeedSkill[identifier], nil
}

// ByUIServantModelPath returns the AvatarServantConfig uniquely identified by UIServantModelPath
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByUIServantModelPath(identifier string) (AvatarServantConfig, error) {
	if a._dataUIServantModelPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUIServantModelPath[identifier], nil
}

// ByUnCreateHeadIconPath returns the AvatarServantConfig uniquely identified by UnCreateHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByUnCreateHeadIconPath(identifier string) (AvatarServantConfig, error) {
	if a._dataUnCreateHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUnCreateHeadIconPath[identifier], nil
}

// ByWaitingServantHeadIconPath returns the AvatarServantConfig uniquely identified by WaitingServantHeadIconPath
//
// Error is only non-nil if the source errors out
func (a *AvatarServantConfigAccessor) ByWaitingServantHeadIconPath(identifier string) (AvatarServantConfig, error) {
	if a._dataWaitingServantHeadIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarServantConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataWaitingServantHeadIconPath[identifier], nil
}
