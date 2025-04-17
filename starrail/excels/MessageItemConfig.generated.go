package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MessageItemConfig struct {
	ContactsID     json.Number            `json:"ContactsID"`
	ID             json.Number            `json:"ID"`
	ItemContentID  json.Number            `json:"ItemContentID"`
	ItemType       string                 `json:"ItemType"`
	MainText       map[string]json.Number `json:"MainText"`
	NextItemIDList []json.Number          `json:"NextItemIDList"`
	OptionText     map[string]json.Number `json:"OptionText"`
	SectionID      json.Number            `json:"SectionID"`
	Sender         string                 `json:"Sender"`
}
type MessageItemConfigAccessor struct {
	_data []MessageItemConfig
}

// LoadData retrieves the data. Must be called before MessageItemConfig.GroupData
func (a *MessageItemConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MessageItemConfig.json")
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
func (a *MessageItemConfigAccessor) Raw() ([]MessageItemConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MessageItemConfig{}, err
		}
	}
	return a._data, nil
}
