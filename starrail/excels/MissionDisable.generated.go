package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MissionDisable struct {
	CompensateItemList             []interface{} `json:"CompensateItemList"`
	MainMissionIDList              []json.Number `json:"MainMissionIDList"`
	MainMissionIDListClientDisplay []json.Number `json:"MainMissionIDListClientDisplay"`
	RecycleItemList                []interface{} `json:"RecycleItemList"`
	SubMissionID                   json.Number   `json:"SubMissionID"`
}
type MissionDisableAccessor struct {
	_data []MissionDisable
}

// LoadData retrieves the data. Must be called before MissionDisable.GroupData
func (a *MissionDisableAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MissionDisable.json")
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
func (a *MissionDisableAccessor) Raw() ([]MissionDisable, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MissionDisable{}, err
		}
	}
	return a._data, nil
}
