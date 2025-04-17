package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type TextJoinItem struct {
	TextJoinItemID json.Number            `json:"TextJoinItemID"`
	TextJoinText   map[string]json.Number `json:"TextJoinText"`
}
type TextJoinItemAccessor struct {
	_data []TextJoinItem
}

// LoadData retrieves the data. Must be called before TextJoinItem.GroupData
func (a *TextJoinItemAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TextJoinItem.json")
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
func (a *TextJoinItemAccessor) Raw() ([]TextJoinItem, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TextJoinItem{}, err
		}
	}
	return a._data, nil
}
