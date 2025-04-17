package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueWolfGunMiracleTarget struct {
	Basement    json.Number `json:"Basement"`
	GameMode    string      `json:"GameMode"`
	LayerMiddle json.Number `json:"LayerMiddle"`
	MiracleID   json.Number `json:"MiracleID"`
	MiraclePic  string      `json:"MiraclePic"`
}
type RogueWolfGunMiracleTargetAccessor struct {
	_data []RogueWolfGunMiracleTarget
}

// LoadData retrieves the data. Must be called before RogueWolfGunMiracleTarget.GroupData
func (a *RogueWolfGunMiracleTargetAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueWolfGunMiracleTarget.json")
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
func (a *RogueWolfGunMiracleTargetAccessor) Raw() ([]RogueWolfGunMiracleTarget, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueWolfGunMiracleTarget{}, err
		}
	}
	return a._data, nil
}
