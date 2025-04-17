package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueUpgradeAvatarSubType struct {
	AvatarID     json.Number `json:"AvatarID"`
	SubRelicType string      `json:"SubRelicType"`
}
type RogueUpgradeAvatarSubTypeAccessor struct {
	_data         []RogueUpgradeAvatarSubType
	_dataAvatarID map[json.Number]RogueUpgradeAvatarSubType
}

// LoadData retrieves the data. Must be called before RogueUpgradeAvatarSubType.GroupData
func (a *RogueUpgradeAvatarSubTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueUpgradeAvatarSubType.json")
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
func (a *RogueUpgradeAvatarSubTypeAccessor) Raw() ([]RogueUpgradeAvatarSubType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueUpgradeAvatarSubType{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueUpgradeAvatarSubTypeAccessor.LoadData to preload everything
func (a *RogueUpgradeAvatarSubTypeAccessor) GroupData() {
	a._dataAvatarID = map[json.Number]RogueUpgradeAvatarSubType{}
	for _, d := range a._data {
		a._dataAvatarID[d.AvatarID] = d
	}
}

// ByAvatarID returns the RogueUpgradeAvatarSubType uniquely identified by AvatarID
//
// Error is only non-nil if the source errors out
func (a *RogueUpgradeAvatarSubTypeAccessor) ByAvatarID(identifier json.Number) (RogueUpgradeAvatarSubType, error) {
	if a._dataAvatarID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueUpgradeAvatarSubType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarID[identifier], nil
}
