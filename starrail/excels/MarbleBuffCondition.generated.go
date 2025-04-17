package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MarbleBuffCondition struct {
	DrawType          string        `json:"DrawType"`
	DrawTypeParameter json.Number   `json:"DrawTypeParameter"`
	ID                json.Number   `json:"ID"`
	OperationType     string        `json:"OperationType"`
	ParamList         []json.Number `json:"ParamList"`
}
type MarbleBuffConditionAccessor struct {
	_data []MarbleBuffCondition
}

// LoadData retrieves the data. Must be called before MarbleBuffCondition.GroupData
func (a *MarbleBuffConditionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MarbleBuffCondition.json")
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
func (a *MarbleBuffConditionAccessor) Raw() ([]MarbleBuffCondition, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MarbleBuffCondition{}, err
		}
	}
	return a._data, nil
}
