package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCAeon struct {
	AeonDiceID                  float64         `json:"AeonDiceID"`
	AeonID                      float64         `json:"AeonID"`
	BattleEventBuffGroup        float64         `json:"BattleEventBuffGroup"`
	BattleEventEnhanceBuffGroup float64         `json:"BattleEventEnhanceBuffGroup"`
	DescParam                   []hash.IntValue `json:"DescParam"`
	EffectDesc3                 hash.Hash       `json:"EffectDesc3"`
	EffectParam1                []float64       `json:"EffectParam1"`
	EffectParam2                []float64       `json:"EffectParam2"`
	EffectParam3                []float64       `json:"EffectParam3"`
	EffectParam4                []float64       `json:"EffectParam4"`
	EffectType1                 string          `json:"EffectType1"`
	EffectType3                 string          `json:"EffectType3"`
	EntrancePrefabPath          string          `json:"EntrancePrefabPath"`
	ExtraEffect                 []float64       `json:"ExtraEffect"`
	PlayShortDesc               hash.Hash       `json:"PlayShortDesc"`
	RogueAeonDisplayID          float64         `json:"RogueAeonDisplayID"`
	RogueBuffType               float64         `json:"RogueBuffType"`
	Sort                        float64         `json:"Sort"`
	UnlockID                    float64         `json:"UnlockID"`
}
type RogueDLCAeonAccessor struct {
	_data                            []RogueDLCAeon
	_dataAeonDiceID                  map[float64]RogueDLCAeon
	_dataAeonID                      map[float64]RogueDLCAeon
	_dataBattleEventBuffGroup        map[float64]RogueDLCAeon
	_dataBattleEventEnhanceBuffGroup map[float64]RogueDLCAeon
	_dataEffectType3                 map[string]RogueDLCAeon
	_dataEntrancePrefabPath          map[string]RogueDLCAeon
	_dataRogueAeonDisplayID          map[float64]RogueDLCAeon
	_dataRogueBuffType               map[float64]RogueDLCAeon
	_dataSort                        map[float64]RogueDLCAeon
}

// LoadData retrieves the data. Must be called before RogueDLCAeon.GroupData
func (a *RogueDLCAeonAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCAeon.json")
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
func (a *RogueDLCAeonAccessor) Raw() ([]RogueDLCAeon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCAeon{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCAeonAccessor.LoadData to preload everything
func (a *RogueDLCAeonAccessor) GroupData() {
	a._dataAeonDiceID = map[float64]RogueDLCAeon{}
	a._dataAeonID = map[float64]RogueDLCAeon{}
	a._dataBattleEventBuffGroup = map[float64]RogueDLCAeon{}
	a._dataBattleEventEnhanceBuffGroup = map[float64]RogueDLCAeon{}
	a._dataEffectType3 = map[string]RogueDLCAeon{}
	a._dataEntrancePrefabPath = map[string]RogueDLCAeon{}
	a._dataRogueAeonDisplayID = map[float64]RogueDLCAeon{}
	a._dataRogueBuffType = map[float64]RogueDLCAeon{}
	a._dataSort = map[float64]RogueDLCAeon{}
	for _, d := range a._data {
		a._dataAeonDiceID[d.AeonDiceID] = d
		a._dataAeonID[d.AeonID] = d
		a._dataBattleEventBuffGroup[d.BattleEventBuffGroup] = d
		a._dataBattleEventEnhanceBuffGroup[d.BattleEventEnhanceBuffGroup] = d
		a._dataEffectType3[d.EffectType3] = d
		a._dataEntrancePrefabPath[d.EntrancePrefabPath] = d
		a._dataRogueAeonDisplayID[d.RogueAeonDisplayID] = d
		a._dataRogueBuffType[d.RogueBuffType] = d
		a._dataSort[d.Sort] = d
	}
}

// ByAeonDiceID returns the RogueDLCAeon uniquely identified by AeonDiceID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonAccessor) ByAeonDiceID(identifier float64) (RogueDLCAeon, error) {
	if a._dataAeonDiceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAeonDiceID[identifier], nil
}

// ByAeonID returns the RogueDLCAeon uniquely identified by AeonID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonAccessor) ByAeonID(identifier float64) (RogueDLCAeon, error) {
	if a._dataAeonID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAeonID[identifier], nil
}

// ByBattleEventBuffGroup returns the RogueDLCAeon uniquely identified by BattleEventBuffGroup
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonAccessor) ByBattleEventBuffGroup(identifier float64) (RogueDLCAeon, error) {
	if a._dataBattleEventBuffGroup == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBattleEventBuffGroup[identifier], nil
}

// ByBattleEventEnhanceBuffGroup returns the RogueDLCAeon uniquely identified by BattleEventEnhanceBuffGroup
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonAccessor) ByBattleEventEnhanceBuffGroup(identifier float64) (RogueDLCAeon, error) {
	if a._dataBattleEventEnhanceBuffGroup == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBattleEventEnhanceBuffGroup[identifier], nil
}

// ByEffectType3 returns the RogueDLCAeon uniquely identified by EffectType3
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonAccessor) ByEffectType3(identifier string) (RogueDLCAeon, error) {
	if a._dataEffectType3 == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEffectType3[identifier], nil
}

// ByEntrancePrefabPath returns the RogueDLCAeon uniquely identified by EntrancePrefabPath
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonAccessor) ByEntrancePrefabPath(identifier string) (RogueDLCAeon, error) {
	if a._dataEntrancePrefabPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEntrancePrefabPath[identifier], nil
}

// ByRogueAeonDisplayID returns the RogueDLCAeon uniquely identified by RogueAeonDisplayID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonAccessor) ByRogueAeonDisplayID(identifier float64) (RogueDLCAeon, error) {
	if a._dataRogueAeonDisplayID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueAeonDisplayID[identifier], nil
}

// ByRogueBuffType returns the RogueDLCAeon uniquely identified by RogueBuffType
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonAccessor) ByRogueBuffType(identifier float64) (RogueDLCAeon, error) {
	if a._dataRogueBuffType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueBuffType[identifier], nil
}

// BySort returns the RogueDLCAeon uniquely identified by Sort
//
// Error is only non-nil if the source errors out
func (a *RogueDLCAeonAccessor) BySort(identifier float64) (RogueDLCAeon, error) {
	if a._dataSort == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSort[identifier], nil
}
