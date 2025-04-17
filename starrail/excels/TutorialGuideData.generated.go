package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type TutorialGuideData struct {
	DescText     map[string]json.Number `json:"DescText"`
	ID           json.Number            `json:"ID"`
	ImagePath    string                 `json:"ImagePath"`
	PlatformType string                 `json:"PlatformType"`
}
type TutorialGuideDataAccessor struct {
	_data []TutorialGuideData
}

// LoadData retrieves the data. Must be called before TutorialGuideData.GroupData
func (a *TutorialGuideDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TutorialGuideData.json")
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
func (a *TutorialGuideDataAccessor) Raw() ([]TutorialGuideData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TutorialGuideData{}, err
		}
	}
	return a._data, nil
}
