package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type SubMission struct {
	DescrptionText map[string]json.Number `json:"DescrptionText"`
	SubMissionID   json.Number            `json:"SubMissionID"`
	TargetText     map[string]json.Number `json:"TargetText"`
}
type SubMissionAccessor struct {
	_data []SubMission
}

// LoadData retrieves the data. Must be called before SubMission.GroupData
func (a *SubMissionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SubMission.json")
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
func (a *SubMissionAccessor) Raw() ([]SubMission, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SubMission{}, err
		}
	}
	return a._data, nil
}
