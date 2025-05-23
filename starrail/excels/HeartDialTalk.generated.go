package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type HeartDialTalk struct {
	FloorIDList []float64 `json:"FloorIDList"`
	ID          float64   `json:"ID"`
	IsKaomoji   bool      `json:"IsKaomoji"`
	SDFText     hash.Hash `json:"SDFText"`
	VoiceID     float64   `json:"VoiceID"`
}
type HeartDialTalkAccessor struct {
	_data   []HeartDialTalk
	_dataID map[float64]HeartDialTalk
}

// LoadData retrieves the data. Must be called before HeartDialTalk.GroupData
func (a *HeartDialTalkAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HeartDialTalk.json")
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
func (a *HeartDialTalkAccessor) Raw() ([]HeartDialTalk, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HeartDialTalk{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HeartDialTalkAccessor.LoadData to preload everything
func (a *HeartDialTalkAccessor) GroupData() {
	a._dataID = map[float64]HeartDialTalk{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the HeartDialTalk uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *HeartDialTalkAccessor) ByID(identifier float64) (HeartDialTalk, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeartDialTalk{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
