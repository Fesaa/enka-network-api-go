package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type VoiceConfig struct {
	IsPlayerInvolved bool    `json:"IsPlayerInvolved"`
	VoiceID          float64 `json:"VoiceID"`
	VoicePath        string  `json:"VoicePath"`
	VoiceType        string  `json:"VoiceType"`
}
type VoiceConfigAccessor struct {
	_data        []VoiceConfig
	_dataVoiceID map[float64]VoiceConfig
}

// LoadData retrieves the data. Must be called before VoiceConfig.GroupData
func (a *VoiceConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/VoiceConfig.json")
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
func (a *VoiceConfigAccessor) Raw() ([]VoiceConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []VoiceConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with VoiceConfigAccessor.LoadData to preload everything
func (a *VoiceConfigAccessor) GroupData() {
	a._dataVoiceID = map[float64]VoiceConfig{}
	for _, d := range a._data {
		a._dataVoiceID[d.VoiceID] = d
	}
}

// ByVoiceID returns the VoiceConfig uniquely identified by VoiceID
//
// Error is only non-nil if the source errors out
func (a *VoiceConfigAccessor) ByVoiceID(identifier float64) (VoiceConfig, error) {
	if a._dataVoiceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return VoiceConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataVoiceID[identifier], nil
}
