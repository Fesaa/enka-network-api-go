package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ActivityVersionBanner struct {
	ActivityID json.Number `json:"ActivityID"`
	Type       string      `json:"Type"`
}
type ActivityVersionBannerAccessor struct {
	_data []ActivityVersionBanner
}

// LoadData retrieves the data. Must be called before ActivityVersionBanner.GroupData
func (a *ActivityVersionBannerAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityVersionBanner.json")
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
func (a *ActivityVersionBannerAccessor) Raw() ([]ActivityVersionBanner, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityVersionBanner{}, err
		}
	}
	return a._data, nil
}
