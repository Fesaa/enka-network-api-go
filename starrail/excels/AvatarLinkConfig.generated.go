package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type AvatarLinkConfig struct {
	AvatarID   float64 `json:"AvatarID"`
	LinkAvatar float64 `json:"LinkAvatar"`
}
type AvatarLinkConfigAccessor struct {
	_data []AvatarLinkConfig
}

// LoadData retrieves the data. Must be called before AvatarLinkConfig.GroupData
func (a *AvatarLinkConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarLinkConfig.json")
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
func (a *AvatarLinkConfigAccessor) Raw() ([]AvatarLinkConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarLinkConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}
