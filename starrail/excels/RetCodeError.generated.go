package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RetCodeError struct {
	ErrorID           json.Number            `json:"ErrorID"`
	IsPileToastCenter bool                   `json:"IsPileToastCenter"`
	Text              map[string]json.Number `json:"Text"`
}
type RetCodeErrorAccessor struct {
	_data []RetCodeError
}

// LoadData retrieves the data. Must be called before RetCodeError.GroupData
func (a *RetCodeErrorAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RetCodeError.json")
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
func (a *RetCodeErrorAccessor) Raw() ([]RetCodeError, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RetCodeError{}, err
		}
	}
	return a._data, nil
}
