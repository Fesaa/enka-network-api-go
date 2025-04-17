package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MusicRhythmPhase struct {
	FinishMissionID float64   `json:"FinishMissionID"`
	LiveName        hash.Hash `json:"LiveName"`
	Phase           float64   `json:"Phase"`
	PostImgPath     string    `json:"PostImgPath"`
	SongID          float64   `json:"SongID"`
	TrackIDList     []float64 `json:"TrackIDList"`
}
type MusicRhythmPhaseAccessor struct {
	_data                []MusicRhythmPhase
	_dataFinishMissionID map[float64]MusicRhythmPhase
	_dataPhase           map[float64]MusicRhythmPhase
	_dataSongID          map[float64]MusicRhythmPhase
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
	for _, d := range a._data {
		a._dataFinishMissionID[d.FinishMissionID] = d
		a._dataPhase[d.Phase] = d
		a._dataSongID[d.SongID] = d
	}
}

// ByFinishMissionID returns the MusicRhythmPhase uniquely identified by FinishMissionID
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmPhaseAccessor) ByFinishMissionID(identifier float64) (MusicRhythmPhase, error) {
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
func (a *MusicRhythmPhaseAccessor) ByPhase(identifier float64) (MusicRhythmPhase, error) {
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
func (a *MusicRhythmPhaseAccessor) BySongID(identifier float64) (MusicRhythmPhase, error) {
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
