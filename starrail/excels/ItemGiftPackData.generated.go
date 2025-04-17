package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ItemGiftPackData struct {
	GroupDesc  map[string]json.Number `json:"GroupDesc"`
	GroupID    json.Number            `json:"GroupID"`
	ID         json.Number            `json:"ID"`
	RewardList []json.Number          `json:"RewardList"`
}
type ItemGiftPackDataAccessor struct {
	_data []ItemGiftPackData
}

// LoadData retrieves the data. Must be called before ItemGiftPackData.GroupData
func (a *ItemGiftPackDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemGiftPackData.json")
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
func (a *ItemGiftPackDataAccessor) Raw() ([]ItemGiftPackData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemGiftPackData{}, err
		}
	}
	return a._data, nil
}
