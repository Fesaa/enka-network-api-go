package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MusicRhythmTrack struct {
	EmptyGridList      []float64 `json:"EmptyGridList"`
	ID                 float64   `json:"ID"`
	IconPath           string    `json:"IconPath"`
	TrackName          hash.Hash `json:"TrackName"`
	UnlockSubMissionID float64   `json:"UnlockSubMissionID"`
}
type MusicRhythmTrackAccessor struct {
	_data   []MusicRhythmTrack
	_dataID map[float64]MusicRhythmTrack
}

// LoadData retrieves the data. Must be called before MusicRhythmTrack.GroupData
func (a *MusicRhythmTrackAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MusicRhythmTrack.json")
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
func (a *MusicRhythmTrackAccessor) Raw() ([]MusicRhythmTrack, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MusicRhythmTrack{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MusicRhythmTrackAccessor.LoadData to preload everything
func (a *MusicRhythmTrackAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MusicRhythmTrack uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmTrackAccessor) ByID(identifier float64) (MusicRhythmTrack, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmTrack{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
