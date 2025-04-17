package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type CharacterNatureConfig struct {
	NatureID   float64 `json:"NatureID"`
	NatureType string  `json:"NatureType"`
	SpritePath string  `json:"SpritePath"`
}
type CharacterNatureConfigAccessor struct {
	_data         []CharacterNatureConfig
	_dataNatureID map[float64]CharacterNatureConfig
}

// LoadData retrieves the data. Must be called before CharacterNatureConfig.GroupData
func (a *CharacterNatureConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/CharacterNatureConfig.json")
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
func (a *CharacterNatureConfigAccessor) Raw() ([]CharacterNatureConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []CharacterNatureConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with CharacterNatureConfigAccessor.LoadData to preload everything
func (a *CharacterNatureConfigAccessor) GroupData() {
	a._dataNatureID = map[float64]CharacterNatureConfig{}
	for _, d := range a._data {
		a._dataNatureID[d.NatureID] = d
	}
}

// ByNatureID returns the CharacterNatureConfig uniquely identified by NatureID
//
// Error is only non-nil if the source errors out
func (a *CharacterNatureConfigAccessor) ByNatureID(identifier float64) (CharacterNatureConfig, error) {
	if a._dataNatureID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return CharacterNatureConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataNatureID[identifier], nil
}
