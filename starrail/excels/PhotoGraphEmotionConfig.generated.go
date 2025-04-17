package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PhotoGraphEmotionConfig struct {
	BrowClipName    string    `json:"BrowClipName"`
	EmotionClipPath string    `json:"EmotionClipPath"`
	EmotionID       float64   `json:"EmotionID"`
	EmotionIconPath string    `json:"EmotionIconPath"`
	EmotionName     hash.Hash `json:"EmotionName"`
	EyeClipName     string    `json:"EyeClipName"`
	MouthClipName   string    `json:"MouthClipName"`
}
type PhotoGraphEmotionConfigAccessor struct {
	_data []PhotoGraphEmotionConfig
}

// LoadData retrieves the data. Must be called before PhotoGraphEmotionConfig.GroupData
func (a *PhotoGraphEmotionConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PhotoGraphEmotionConfig.json")
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
func (a *PhotoGraphEmotionConfigAccessor) Raw() ([]PhotoGraphEmotionConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PhotoGraphEmotionConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}
