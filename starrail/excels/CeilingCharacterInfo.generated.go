package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type CeilingCharacterInfo struct {
	CeilingDesc map[string]json.Number `json:"CeilingDesc"`
	CharacterID json.Number            `json:"CharacterID"`
}
type CeilingCharacterInfoAccessor struct {
	_data            []CeilingCharacterInfo
	_dataCharacterID map[json.Number]CeilingCharacterInfo
}

// LoadData retrieves the data. Must be called before CeilingCharacterInfo.GroupData
func (a *CeilingCharacterInfoAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/CeilingCharacterInfo.json")
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
func (a *CeilingCharacterInfoAccessor) Raw() ([]CeilingCharacterInfo, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []CeilingCharacterInfo{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with CeilingCharacterInfoAccessor.LoadData to preload everything
func (a *CeilingCharacterInfoAccessor) GroupData() {
	a._dataCharacterID = map[json.Number]CeilingCharacterInfo{}
	for _, d := range a._data {
		a._dataCharacterID[d.CharacterID] = d
	}
}

// ByCharacterID returns the CeilingCharacterInfo uniquely identified by CharacterID
//
// Error is only non-nil if the source errors out
func (a *CeilingCharacterInfoAccessor) ByCharacterID(identifier json.Number) (CeilingCharacterInfo, error) {
	if a._dataCharacterID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return CeilingCharacterInfo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCharacterID[identifier], nil
}
