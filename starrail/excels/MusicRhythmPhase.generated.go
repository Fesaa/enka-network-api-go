package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
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
	_data                []MusicRhythmPhase
	_dataFinishMissionID map[json.Number]MusicRhythmPhase
	_dataPhase           map[json.Number]MusicRhythmPhase
	_dataSongID          map[json.Number]MusicRhythmPhase
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

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MusicRhythmPhaseAccessor.LoadData to preload everything
func (a *MusicRhythmPhaseAccessor) GroupData() {
	a._dataFinishMissionID = map[json.Number]MusicRhythmPhase{}
	a._dataPhase = map[json.Number]MusicRhythmPhase{}
	a._dataSongID = map[json.Number]MusicRhythmPhase{}
	for _, d := range a._data {
		a._dataFinishMissionID[d.FinishMissionID] = d
		a._dataPhase[d.Phase] = d
		a._dataSongID[d.SongID] = d
	}
}

// ByFinishMissionID returns the MusicRhythmPhase uniquely identified by FinishMissionID
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmPhaseAccessor) ByFinishMissionID(identifier json.Number) (MusicRhythmPhase, error) {
	if a._dataFinishMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFinishMissionID[identifier], nil
}

// ByPhase returns the MusicRhythmPhase uniquely identified by Phase
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmPhaseAccessor) ByPhase(identifier json.Number) (MusicRhythmPhase, error) {
	if a._dataPhase == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPhase[identifier], nil
}

// BySongID returns the MusicRhythmPhase uniquely identified by SongID
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmPhaseAccessor) BySongID(identifier json.Number) (MusicRhythmPhase, error) {
	if a._dataSongID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSongID[identifier], nil
}
