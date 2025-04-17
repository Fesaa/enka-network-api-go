package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MusicRhythmPhase struct {
	FinishMissionID json.Number            `json:"FinishMissionID"`
	LiveName        map[string]json.Number `json:"LiveName"`
	Phase           json.Number            `json:"Phase"`
	PostImgPath     string                 `json:"PostImgPath"`
	SongID          json.Number            `json:"SongID"`
	TrackIDList     []json.Number          `json:"TrackIDList"`
}
type MusicRhythmPhaseAccessor struct {
	_data []MusicRhythmPhase
}

// LoadData retrieves the data. Must be called before MusicRhythmPhase.GroupData
func (a *MusicRhythmPhaseAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MusicRhythmPhase.json")
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
func (a *MusicRhythmPhaseAccessor) Raw() ([]MusicRhythmPhase, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MusicRhythmPhase{}, err
		}
	}
	return a._data, nil
}
