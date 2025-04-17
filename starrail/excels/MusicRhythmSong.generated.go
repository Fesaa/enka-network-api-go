package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MusicRhythmSong struct {
	BGMMenuState                string                 `json:"BGMMenuState"`
	BGMStageState               string                 `json:"BGMStageState"`
	GridNum                     json.Number            `json:"GridNum"`
	GridNumList                 []json.Number          `json:"GridNumList"`
	GridTransitionTime          json.Number            `json:"GridTransitionTime"`
	ID                          json.Number            `json:"ID"`
	MixingWaveMatPath           string                 `json:"MixingWaveMatPath"`
	PresetEndGrid               json.Number            `json:"PresetEndGrid"`
	PresetIDList                []json.Number          `json:"PresetIDList"`
	PresetStartGrid             json.Number            `json:"PresetStartGrid"`
	SongName                    map[string]json.Number `json:"SongName"`
	SoundEffectIDList           []json.Number          `json:"SoundEffectIDList"`
	SoundEffectUnlockSubMission json.Number            `json:"SoundEffectUnlockSubMission"`
	TrackIDList                 []json.Number          `json:"TrackIDList"`
	UnlockType                  json.Number            `json:"UnlockType"`
	UnlockTypeParam             json.Number            `json:"UnlockTypeParam"`
}
type MusicRhythmSongAccessor struct {
	_data                  []MusicRhythmSong
	_dataBGMMenuState      map[string]MusicRhythmSong
	_dataBGMStageState     map[string]MusicRhythmSong
	_dataMixingWaveMatPath map[string]MusicRhythmSong
}

// LoadData retrieves the data. Must be called before MusicRhythmSong.GroupData
func (a *MusicRhythmSongAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MusicRhythmSong.json")
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
func (a *MusicRhythmSongAccessor) Raw() ([]MusicRhythmSong, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MusicRhythmSong{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MusicRhythmSongAccessor.LoadData to preload everything
func (a *MusicRhythmSongAccessor) GroupData() {
	a._dataBGMMenuState = map[string]MusicRhythmSong{}
	a._dataBGMStageState = map[string]MusicRhythmSong{}
	a._dataMixingWaveMatPath = map[string]MusicRhythmSong{}
	for _, d := range a._data {
		a._dataBGMMenuState[d.BGMMenuState] = d
		a._dataBGMStageState[d.BGMStageState] = d
		a._dataMixingWaveMatPath[d.MixingWaveMatPath] = d
	}
}

// ByBGMMenuState returns the MusicRhythmSong uniquely identified by BGMMenuState
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmSongAccessor) ByBGMMenuState(identifier string) (MusicRhythmSong, error) {
	if a._dataBGMMenuState == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmSong{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBGMMenuState[identifier], nil
}

// ByBGMStageState returns the MusicRhythmSong uniquely identified by BGMStageState
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmSongAccessor) ByBGMStageState(identifier string) (MusicRhythmSong, error) {
	if a._dataBGMStageState == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmSong{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBGMStageState[identifier], nil
}

// ByMixingWaveMatPath returns the MusicRhythmSong uniquely identified by MixingWaveMatPath
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmSongAccessor) ByMixingWaveMatPath(identifier string) (MusicRhythmSong, error) {
	if a._dataMixingWaveMatPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmSong{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMixingWaveMatPath[identifier], nil
}
