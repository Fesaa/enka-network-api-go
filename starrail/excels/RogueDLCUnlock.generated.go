package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCUnlock struct {
	RogueUnlockDetail hash.Hash `json:"RogueUnlockDetail"`
	RogueUnlockID     float64   `json:"RogueUnlockID"`
	UnlockFinishWay   float64   `json:"UnlockFinishWay"`
}
type RogueDLCUnlockAccessor struct {
	_data              []RogueDLCUnlock
	_dataRogueUnlockID map[float64]RogueDLCUnlock
}

// LoadData retrieves the data. Must be called before RogueDLCUnlock.GroupData
func (a *RogueDLCUnlockAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCUnlock.json")
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
func (a *RogueDLCUnlockAccessor) Raw() ([]RogueDLCUnlock, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCUnlock{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCUnlockAccessor.LoadData to preload everything
func (a *RogueDLCUnlockAccessor) GroupData() {
	a._dataRogueUnlockID = map[float64]RogueDLCUnlock{}
	for _, d := range a._data {
		a._dataRogueUnlockID[d.RogueUnlockID] = d
	}
}

// ByRogueUnlockID returns the RogueDLCUnlock uniquely identified by RogueUnlockID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCUnlockAccessor) ByRogueUnlockID(identifier float64) (RogueDLCUnlock, error) {
	if a._dataRogueUnlockID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCUnlock{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueUnlockID[identifier], nil
}
