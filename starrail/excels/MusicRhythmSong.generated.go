package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
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
	_data                   []MusicRhythmSong
	_dataBGMMenuState       map[string]MusicRhythmSong
	_dataBGMStageState      map[string]MusicRhythmSong
	_dataGridTransitionTime map[json.Number]MusicRhythmSong
	_dataID                 map[json.Number]MusicRhythmSong
	_dataMixingWaveMatPath  map[string]MusicRhythmSong
	_dataPresetEndGrid      map[json.Number]MusicRhythmSong
	_dataUnlockTypeParam    map[json.Number]MusicRhythmSong
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
	a._dataGridTransitionTime = map[json.Number]MusicRhythmSong{}
	a._dataID = map[json.Number]MusicRhythmSong{}
	a._dataMixingWaveMatPath = map[string]MusicRhythmSong{}
	a._dataPresetEndGrid = map[json.Number]MusicRhythmSong{}
	a._dataUnlockTypeParam = map[json.Number]MusicRhythmSong{}
	for _, d := range a._data {
		a._dataBGMMenuState[d.BGMMenuState] = d
		a._dataBGMStageState[d.BGMStageState] = d
		a._dataGridTransitionTime[d.GridTransitionTime] = d
		a._dataID[d.ID] = d
		a._dataMixingWaveMatPath[d.MixingWaveMatPath] = d
		a._dataPresetEndGrid[d.PresetEndGrid] = d
		a._dataUnlockTypeParam[d.UnlockTypeParam] = d
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

// ByGridTransitionTime returns the MusicRhythmSong uniquely identified by GridTransitionTime
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmSongAccessor) ByGridTransitionTime(identifier json.Number) (MusicRhythmSong, error) {
	if a._dataGridTransitionTime == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmSong{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGridTransitionTime[identifier], nil
}

// ByID returns the MusicRhythmSong uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmSongAccessor) ByID(identifier json.Number) (MusicRhythmSong, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmSong{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
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

// ByPresetEndGrid returns the MusicRhythmSong uniquely identified by PresetEndGrid
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmSongAccessor) ByPresetEndGrid(identifier json.Number) (MusicRhythmSong, error) {
	if a._dataPresetEndGrid == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmSong{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPresetEndGrid[identifier], nil
}

// ByUnlockTypeParam returns the MusicRhythmSong uniquely identified by UnlockTypeParam
//
// Error is only non-nil if the source errors out
func (a *MusicRhythmSongAccessor) ByUnlockTypeParam(identifier json.Number) (MusicRhythmSong, error) {
	if a._dataUnlockTypeParam == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MusicRhythmSong{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUnlockTypeParam[identifier], nil
}
