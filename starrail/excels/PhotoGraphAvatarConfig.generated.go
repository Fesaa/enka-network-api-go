package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PhotoGraphAvatarConfig struct {
	AvatarID          float64   `json:"AvatarID"`
	EmotionConfigList []float64 `json:"EmotionConfigList"`
}
type PhotoGraphAvatarConfigAccessor struct {
	_data         []PhotoGraphAvatarConfig
	_dataAvatarID map[float64]PhotoGraphAvatarConfig
}

// LoadData retrieves the data. Must be called before PhotoGraphAvatarConfig.GroupData
func (a *PhotoGraphAvatarConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PhotoGraphAvatarConfig.json")
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
func (a *PhotoGraphAvatarConfigAccessor) Raw() ([]PhotoGraphAvatarConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PhotoGraphAvatarConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PhotoGraphAvatarConfigAccessor.LoadData to preload everything
func (a *PhotoGraphAvatarConfigAccessor) GroupData() {
	a._dataAvatarID = map[float64]PhotoGraphAvatarConfig{}
	for _, d := range a._data {
		a._dataAvatarID[d.AvatarID] = d
	}
}

// ByAvatarID returns the PhotoGraphAvatarConfig uniquely identified by AvatarID
//
// Error is only non-nil if the source errors out
func (a *PhotoGraphAvatarConfigAccessor) ByAvatarID(identifier float64) (PhotoGraphAvatarConfig, error) {
	if a._dataAvatarID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PhotoGraphAvatarConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarID[identifier], nil
}
