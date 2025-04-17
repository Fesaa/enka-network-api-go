package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ItemCureInfoData struct {
	CureInfoDesc  map[string]json.Number `json:"CureInfoDesc"`
	CureInfoTitle map[string]json.Number `json:"CureInfoTitle"`
	ID            json.Number            `json:"ID"`
	ImgPath       string                 `json:"ImgPath"`
}
type ItemCureInfoDataAccessor struct {
	_data []ItemCureInfoData
}

// LoadData retrieves the data. Must be called before ItemCureInfoData.GroupData
func (a *ItemCureInfoDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemCureInfoData.json")
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
func (a *ItemCureInfoDataAccessor) Raw() ([]ItemCureInfoData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemCureInfoData{}, err
		}
	}
	return a._data, nil
}
