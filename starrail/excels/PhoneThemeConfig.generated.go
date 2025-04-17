package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type PhoneThemeConfig struct {
	ID             json.Number `json:"ID"`
	PhoneThemeApp  string      `json:"PhoneThemeApp"`
	PhoneThemeItem string      `json:"PhoneThemeItem"`
	PhoneThemeMain string      `json:"PhoneThemeMain"`
	ShowParam      json.Number `json:"ShowParam"`
	ShowType       string      `json:"ShowType"`
}
type PhoneThemeConfigAccessor struct {
	_data               []PhoneThemeConfig
	_dataPhoneThemeApp  map[string]PhoneThemeConfig
	_dataPhoneThemeItem map[string]PhoneThemeConfig
	_dataPhoneThemeMain map[string]PhoneThemeConfig
}

// LoadData retrieves the data. Must be called before PhoneThemeConfig.GroupData
func (a *PhoneThemeConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PhoneThemeConfig.json")
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
func (a *PhoneThemeConfigAccessor) Raw() ([]PhoneThemeConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PhoneThemeConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PhoneThemeConfigAccessor.LoadData to preload everything
func (a *PhoneThemeConfigAccessor) GroupData() {
	a._dataPhoneThemeApp = map[string]PhoneThemeConfig{}
	a._dataPhoneThemeItem = map[string]PhoneThemeConfig{}
	a._dataPhoneThemeMain = map[string]PhoneThemeConfig{}
	for _, d := range a._data {
		a._dataPhoneThemeApp[d.PhoneThemeApp] = d
		a._dataPhoneThemeItem[d.PhoneThemeItem] = d
		a._dataPhoneThemeMain[d.PhoneThemeMain] = d
	}
}

// ByPhoneThemeApp returns the PhoneThemeConfig uniquely identified by PhoneThemeApp
//
// Error is only non-nil if the source errors out
func (a *PhoneThemeConfigAccessor) ByPhoneThemeApp(identifier string) (PhoneThemeConfig, error) {
	if a._dataPhoneThemeApp == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PhoneThemeConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPhoneThemeApp[identifier], nil
}

// ByPhoneThemeItem returns the PhoneThemeConfig uniquely identified by PhoneThemeItem
//
// Error is only non-nil if the source errors out
func (a *PhoneThemeConfigAccessor) ByPhoneThemeItem(identifier string) (PhoneThemeConfig, error) {
	if a._dataPhoneThemeItem == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PhoneThemeConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPhoneThemeItem[identifier], nil
}

// ByPhoneThemeMain returns the PhoneThemeConfig uniquely identified by PhoneThemeMain
//
// Error is only non-nil if the source errors out
func (a *PhoneThemeConfigAccessor) ByPhoneThemeMain(identifier string) (PhoneThemeConfig, error) {
	if a._dataPhoneThemeMain == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PhoneThemeConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPhoneThemeMain[identifier], nil
}
