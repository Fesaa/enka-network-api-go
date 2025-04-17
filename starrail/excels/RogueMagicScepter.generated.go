package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueMagicScepter struct {
	EffectTypeList    []string                 `json:"EffectTypeList"`
	FuncType          string                   `json:"FuncType"`
	LimitRangeType    string                   `json:"LimitRangeType"`
	LockMagicUnit     []map[string]json.Number `json:"LockMagicUnit"`
	ScepterBasicPower map[string]json.Number   `json:"ScepterBasicPower"`
	ScepterID         json.Number              `json:"ScepterID"`
	ScepterLevel      json.Number              `json:"ScepterLevel"`
	StaffMazeBuffID   json.Number              `json:"StaffMazeBuffID"`
	StyleType         string                   `json:"StyleType"`
	TrenchCount       map[string]json.Number   `json:"TrenchCount"`
	UnlockID          json.Number              `json:"UnlockID"`
}
type RogueMagicScepterAccessor struct {
	_data []RogueMagicScepter
}

// LoadData retrieves the data. Must be called before RogueMagicScepter.GroupData
func (a *RogueMagicScepterAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicScepter.json")
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
func (a *RogueMagicScepterAccessor) Raw() ([]RogueMagicScepter, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicScepter{}, err
		}
	}
	return a._data, nil
}
