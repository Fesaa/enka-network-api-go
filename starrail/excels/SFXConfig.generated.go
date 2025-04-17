package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type SFXConfig struct {
	IsPlayerInvolved bool        `json:"IsPlayerInvolved"`
	SFXID            json.Number `json:"SFXID"`
	SFXPath          string      `json:"SFXPath"`
	SFXType          string      `json:"SFXType"`
}
type SFXConfigAccessor struct {
	_data []SFXConfig
}

// LoadData retrieves the data. Must be called before SFXConfig.GroupData
func (a *SFXConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SFXConfig.json")
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
func (a *SFXConfigAccessor) Raw() ([]SFXConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SFXConfig{}, err
		}
	}
	return a._data, nil
}
