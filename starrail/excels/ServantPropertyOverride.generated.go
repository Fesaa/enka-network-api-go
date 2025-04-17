package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ServantPropertyOverride struct {
	HidePropertyInBattleList []string    `json:"HidePropertyInBattleList"`
	HidePropertyList         []string    `json:"HidePropertyList"`
	SecretPropertyList       []string    `json:"SecretPropertyList"`
	ServantID                json.Number `json:"ServantID"`
}
type ServantPropertyOverrideAccessor struct {
	_data []ServantPropertyOverride
}

// LoadData retrieves the data. Must be called before ServantPropertyOverride.GroupData
func (a *ServantPropertyOverrideAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ServantPropertyOverride.json")
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
func (a *ServantPropertyOverrideAccessor) Raw() ([]ServantPropertyOverride, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ServantPropertyOverride{}, err
		}
	}
	return a._data, nil
}
