package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MusicRhythmSong struct {
	BGMMenuState                string    `json:"BGMMenuState"`
	BGMStageState               string    `json:"BGMStageState"`
	GridNum                     float64   `json:"GridNum"`
	GridNumList                 []float64 `json:"GridNumList"`
	GridTransitionTime          float64   `json:"GridTransitionTime"`
	ID                          float64   `json:"ID"`
	MixingWaveMatPath           string    `json:"MixingWaveMatPath"`
	PresetEndGrid               float64   `json:"PresetEndGrid"`
	PresetIDList                []float64 `json:"PresetIDList"`
	PresetStartGrid             float64   `json:"PresetStartGrid"`
	SongName                    hash.Hash `json:"SongName"`
	SoundEffectIDList           []float64 `json:"SoundEffectIDList"`
	SoundEffectUnlockSubMission float64   `json:"SoundEffectUnlockSubMission"`
	TrackIDList                 []float64 `json:"TrackIDList"`
	UnlockType                  float64   `json:"UnlockType"`
	UnlockTypeParam             float64   `json:"UnlockTypeParam"`
}
type MusicRhythmSongAccessor struct {
	_data                   []MusicRhythmSong
	_dataBGMMenuState       map[string]MusicRhythmSong
	_dataBGMStageState      map[string]MusicRhythmSong
	_dataGridTransitionTime map[float64]MusicRhythmSong
	_dataID                 map[float64]MusicRhythmSong
	_dataMixingWaveMatPath  map[string]MusicRhythmSong
	_dataPresetEndGrid      map[float64]MusicRhythmSong
	_dataUnlockTypeParam    map[float64]MusicRhythmSong
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
func (a *MusicRhythmSongAccessor) ByGridTransitionTime(identifier float64) (MusicRhythmSong, error) {
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
func (a *MusicRhythmSongAccessor) ByID(identifier float64) (MusicRhythmSong, error) {
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
func (a *MusicRhythmSongAccessor) ByPresetEndGrid(identifier float64) (MusicRhythmSong, error) {
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
func (a *MusicRhythmSongAccessor) ByUnlockTypeParam(identifier float64) (MusicRhythmSong, error) {
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
