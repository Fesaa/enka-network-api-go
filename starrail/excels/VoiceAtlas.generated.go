package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type VoiceAtlas struct {
	AudioEvent    string                 `json:"AudioEvent"`
	AudioID       json.Number            `json:"AudioID"`
	AvatarID      json.Number            `json:"AvatarID"`
	IsBattleVoice bool                   `json:"IsBattleVoice"`
	ReplaceID     json.Number            `json:"ReplaceID"`
	SortID        json.Number            `json:"SortID"`
	Unlock        json.Number            `json:"Unlock"`
	VoiceID       json.Number            `json:"VoiceID"`
	VoiceTitle    map[string]json.Number `json:"VoiceTitle"`
	VoiceM        map[string]json.Number `json:"Voice_M"`
}
type VoiceAtlasAccessor struct {
	_data []VoiceAtlas
}

// LoadData retrieves the data. Must be called before VoiceAtlas.GroupData
func (a *VoiceAtlasAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/VoiceAtlas.json")
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
func (a *VoiceAtlasAccessor) Raw() ([]VoiceAtlas, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []VoiceAtlas{}, err
		}
	}
	return a._data, nil
}
