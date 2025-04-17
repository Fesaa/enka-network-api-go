package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueAeon struct {
	AeonID                      float64   `json:"AeonID"`
	ArrivedTalkDialogueGroupID  float64   `json:"ArrivedTalkDialogueGroupID"`
	BattleEventBuffGroup        float64   `json:"BattleEventBuffGroup"`
	BattleEventEnhanceBuffGroup float64   `json:"BattleEventEnhanceBuffGroup"`
	DisplayID                   float64   `json:"DisplayID"`
	EffectDesc1                 hash.Hash `json:"EffectDesc1"`
	EffectDesc2                 hash.Hash `json:"EffectDesc2"`
	RogueBuffType               float64   `json:"RogueBuffType"`
	RogueVersion                float64   `json:"RogueVersion"`
	Sort                        float64   `json:"Sort"`
	UnlockID                    float64   `json:"UnlockID"`
}
type RogueAeonAccessor struct {
	_data                            []RogueAeon
	_dataAeonID                      map[float64]RogueAeon
	_dataBattleEventBuffGroup        map[float64]RogueAeon
	_dataBattleEventEnhanceBuffGroup map[float64]RogueAeon
	_dataDisplayID                   map[float64]RogueAeon
	_dataRogueBuffType               map[float64]RogueAeon
	_dataSort                        map[float64]RogueAeon
}

// LoadData retrieves the data. Must be called before RogueAeon.GroupData
func (a *RogueAeonAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueAeon.json")
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
func (a *RogueAeonAccessor) Raw() ([]RogueAeon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueAeon{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueAeonAccessor.LoadData to preload everything
func (a *RogueAeonAccessor) GroupData() {
	a._dataAeonID = map[float64]RogueAeon{}
	a._dataBattleEventBuffGroup = map[float64]RogueAeon{}
	a._dataBattleEventEnhanceBuffGroup = map[float64]RogueAeon{}
	a._dataDisplayID = map[float64]RogueAeon{}
	a._dataRogueBuffType = map[float64]RogueAeon{}
	a._dataSort = map[float64]RogueAeon{}
	for _, d := range a._data {
		a._dataAeonID[d.AeonID] = d
		a._dataBattleEventBuffGroup[d.BattleEventBuffGroup] = d
		a._dataBattleEventEnhanceBuffGroup[d.BattleEventEnhanceBuffGroup] = d
		a._dataDisplayID[d.DisplayID] = d
		a._dataRogueBuffType[d.RogueBuffType] = d
		a._dataSort[d.Sort] = d
	}
}

// ByAeonID returns the RogueAeon uniquely identified by AeonID
//
// Error is only non-nil if the source errors out
func (a *RogueAeonAccessor) ByAeonID(identifier float64) (RogueAeon, error) {
	if a._dataAeonID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAeonID[identifier], nil
}

// ByBattleEventBuffGroup returns the RogueAeon uniquely identified by BattleEventBuffGroup
//
// Error is only non-nil if the source errors out
func (a *RogueAeonAccessor) ByBattleEventBuffGroup(identifier float64) (RogueAeon, error) {
	if a._dataBattleEventBuffGroup == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBattleEventBuffGroup[identifier], nil
}

// ByBattleEventEnhanceBuffGroup returns the RogueAeon uniquely identified by BattleEventEnhanceBuffGroup
//
// Error is only non-nil if the source errors out
func (a *RogueAeonAccessor) ByBattleEventEnhanceBuffGroup(identifier float64) (RogueAeon, error) {
	if a._dataBattleEventEnhanceBuffGroup == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBattleEventEnhanceBuffGroup[identifier], nil
}

// ByDisplayID returns the RogueAeon uniquely identified by DisplayID
//
// Error is only non-nil if the source errors out
func (a *RogueAeonAccessor) ByDisplayID(identifier float64) (RogueAeon, error) {
	if a._dataDisplayID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDisplayID[identifier], nil
}

// ByRogueBuffType returns the RogueAeon uniquely identified by RogueBuffType
//
// Error is only non-nil if the source errors out
func (a *RogueAeonAccessor) ByRogueBuffType(identifier float64) (RogueAeon, error) {
	if a._dataRogueBuffType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueBuffType[identifier], nil
}

// BySort returns the RogueAeon uniquely identified by Sort
//
// Error is only non-nil if the source errors out
func (a *RogueAeonAccessor) BySort(identifier float64) (RogueAeon, error) {
	if a._dataSort == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueAeon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSort[identifier], nil
}
