package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type AetherDivideQuestType struct {
	ID            float64   `json:"ID"`
	TypeGroupList []float64 `json:"TypeGroupList"`
}
type AetherDivideQuestTypeAccessor struct {
	_data   []AetherDivideQuestType
	_dataID map[float64]AetherDivideQuestType
}

// LoadData retrieves the data. Must be called before AetherDivideQuestType.GroupData
func (a *AetherDivideQuestTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AetherDivideQuestType.json")
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
func (a *AetherDivideQuestTypeAccessor) Raw() ([]AetherDivideQuestType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AetherDivideQuestType{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AetherDivideQuestTypeAccessor.LoadData to preload everything
func (a *AetherDivideQuestTypeAccessor) GroupData() {
	a._dataID = map[float64]AetherDivideQuestType{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the AetherDivideQuestType uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideQuestTypeAccessor) ByID(identifier float64) (AetherDivideQuestType, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideQuestType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
