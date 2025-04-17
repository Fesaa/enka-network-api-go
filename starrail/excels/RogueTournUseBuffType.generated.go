package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueTournUseBuffType struct {
	TournMode       string        `json:"TournMode"`
	UseBuffTypeList []json.Number `json:"UseBuffTypeList"`
}
type RogueTournUseBuffTypeAccessor struct {
	_data []RogueTournUseBuffType
}

// LoadData retrieves the data. Must be called before RogueTournUseBuffType.GroupData
func (a *RogueTournUseBuffTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournUseBuffType.json")
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
func (a *RogueTournUseBuffTypeAccessor) Raw() ([]RogueTournUseBuffType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournUseBuffType{}, err
		}
	}
	return a._data, nil
}
