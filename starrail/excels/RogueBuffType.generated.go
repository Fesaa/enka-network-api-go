package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueBuffType struct {
	HintDesc                     map[string]json.Number `json:"HintDesc"`
	RogueBuffType                json.Number            `json:"RogueBuffType"`
	RogueBuffTypeIcon            string                 `json:"RogueBuffTypeIcon"`
	RogueBuffTypeSubTitle        map[string]json.Number `json:"RogueBuffTypeSubTitle"`
	RogueBuffTypeTextmapID       map[string]json.Number `json:"RogueBuffTypeTextmapID"`
	RogueBuffTypeTitle           map[string]json.Number `json:"RogueBuffTypeTitle"`
	RugueBuffTypeRewardQuestList []json.Number          `json:"RugueBuffTypeRewardQuestList"`
}
type RogueBuffTypeAccessor struct {
	_data                  []RogueBuffType
	_dataRogueBuffType     map[json.Number]RogueBuffType
	_dataRogueBuffTypeIcon map[string]RogueBuffType
}

// LoadData retrieves the data. Must be called before RogueBuffType.GroupData
func (a *RogueBuffTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueBuffType.json")
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
func (a *RogueBuffTypeAccessor) Raw() ([]RogueBuffType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueBuffType{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueBuffTypeAccessor.LoadData to preload everything
func (a *RogueBuffTypeAccessor) GroupData() {
	a._dataRogueBuffType = map[json.Number]RogueBuffType{}
	a._dataRogueBuffTypeIcon = map[string]RogueBuffType{}
	for _, d := range a._data {
		a._dataRogueBuffType[d.RogueBuffType] = d
		a._dataRogueBuffTypeIcon[d.RogueBuffTypeIcon] = d
	}
}

// ByRogueBuffType returns the RogueBuffType uniquely identified by RogueBuffType
//
// Error is only non-nil if the source errors out
func (a *RogueBuffTypeAccessor) ByRogueBuffType(identifier json.Number) (RogueBuffType, error) {
	if a._dataRogueBuffType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueBuffType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueBuffType[identifier], nil
}

// ByRogueBuffTypeIcon returns the RogueBuffType uniquely identified by RogueBuffTypeIcon
//
// Error is only non-nil if the source errors out
func (a *RogueBuffTypeAccessor) ByRogueBuffTypeIcon(identifier string) (RogueBuffType, error) {
	if a._dataRogueBuffTypeIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueBuffType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueBuffTypeIcon[identifier], nil
}
