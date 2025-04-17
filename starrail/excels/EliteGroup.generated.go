package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type EliteGroup struct {
	AttackRatio  map[string]json.Number `json:"AttackRatio"`
	DefenceRatio map[string]json.Number `json:"DefenceRatio"`
	EliteGroup   json.Number            `json:"EliteGroup"`
	HPRatio      map[string]json.Number `json:"HPRatio"`
	SpeedRatio   map[string]json.Number `json:"SpeedRatio"`
	StanceRatio  map[string]json.Number `json:"StanceRatio"`
}
type EliteGroupAccessor struct {
	_data []EliteGroup
}

// LoadData retrieves the data. Must be called before EliteGroup.GroupData
func (a *EliteGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EliteGroup.json")
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
func (a *EliteGroupAccessor) Raw() ([]EliteGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EliteGroup{}, err
		}
	}
	return a._data, nil
}
