package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TarotFiles struct {
	ID       float64   `json:"ID"`
	Sentence hash.Hash `json:"Sentence"`
	VoiceID  float64   `json:"VoiceID"`
}
type TarotFilesAccessor struct {
	_data        []TarotFiles
	_dataID      map[float64]TarotFiles
	_dataVoiceID map[float64]TarotFiles
}

// LoadData retrieves the data. Must be called before TarotFiles.GroupData
func (a *TarotFilesAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TarotFiles.json")
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
func (a *TarotFilesAccessor) Raw() ([]TarotFiles, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TarotFiles{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TarotFilesAccessor.LoadData to preload everything
func (a *TarotFilesAccessor) GroupData() {
	a._dataID = map[float64]TarotFiles{}
	a._dataVoiceID = map[float64]TarotFiles{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataVoiceID[d.VoiceID] = d
	}
}

// ByID returns the TarotFiles uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TarotFilesAccessor) ByID(identifier float64) (TarotFiles, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TarotFiles{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByVoiceID returns the TarotFiles uniquely identified by VoiceID
//
// Error is only non-nil if the source errors out
func (a *TarotFilesAccessor) ByVoiceID(identifier float64) (TarotFiles, error) {
	if a._dataVoiceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TarotFiles{}, err
			}
		}
		a.GroupData()
	}
	return a._dataVoiceID[identifier], nil
}
