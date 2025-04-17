package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TitanAtlasVoicePool struct {
	AudioEvent       string  `json:"AudioEvent"`
	TitanVoiceID     float64 `json:"TitanVoiceID"`
	TitanVoicePoolID float64 `json:"TitanVoicePoolID"`
	Weight           float64 `json:"Weight"`
}
type TitanAtlasVoicePoolAccessor struct {
	_data                 []TitanAtlasVoicePool
	_dataAudioEvent       map[string]TitanAtlasVoicePool
	_dataTitanVoiceID     map[float64]TitanAtlasVoicePool
	_dataTitanVoicePoolID map[float64]TitanAtlasVoicePool
}

// LoadData retrieves the data. Must be called before TitanAtlasVoicePool.GroupData
func (a *TitanAtlasVoicePoolAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TitanAtlasVoicePool.json")
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
func (a *TitanAtlasVoicePoolAccessor) Raw() ([]TitanAtlasVoicePool, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TitanAtlasVoicePool{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TitanAtlasVoicePoolAccessor.LoadData to preload everything
func (a *TitanAtlasVoicePoolAccessor) GroupData() {
	for _, d := range a._data {
		a._dataAudioEvent[d.AudioEvent] = d
		a._dataTitanVoiceID[d.TitanVoiceID] = d
		a._dataTitanVoicePoolID[d.TitanVoicePoolID] = d
	}
}

// ByAudioEvent returns the TitanAtlasVoicePool uniquely identified by AudioEvent
//
// Error is only non-nil if the source errors out
func (a *TitanAtlasVoicePoolAccessor) ByAudioEvent(identifier string) (TitanAtlasVoicePool, error) {
	if a._dataAudioEvent == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TitanAtlasVoicePool{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAudioEvent[identifier], nil
}

// ByTitanVoiceID returns the TitanAtlasVoicePool uniquely identified by TitanVoiceID
//
// Error is only non-nil if the source errors out
func (a *TitanAtlasVoicePoolAccessor) ByTitanVoiceID(identifier float64) (TitanAtlasVoicePool, error) {
	if a._dataTitanVoiceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TitanAtlasVoicePool{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTitanVoiceID[identifier], nil
}

// ByTitanVoicePoolID returns the TitanAtlasVoicePool uniquely identified by TitanVoicePoolID
//
// Error is only non-nil if the source errors out
func (a *TitanAtlasVoicePoolAccessor) ByTitanVoicePoolID(identifier float64) (TitanAtlasVoicePool, error) {
	if a._dataTitanVoicePoolID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TitanAtlasVoicePool{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTitanVoicePoolID[identifier], nil
}
