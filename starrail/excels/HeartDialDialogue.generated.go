package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type HeartDialDialogue struct {
	ControlTalkList []float64 `json:"ControlTalkList"`
	ID              float64   `json:"ID"`
	RewardID        float64   `json:"RewardID"`
}
type HeartDialDialogueAccessor struct {
	_data   []HeartDialDialogue
	_dataID map[float64]HeartDialDialogue
}

// LoadData retrieves the data. Must be called before HeartDialDialogue.GroupData
func (a *HeartDialDialogueAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HeartDialDialogue.json")
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
func (a *HeartDialDialogueAccessor) Raw() ([]HeartDialDialogue, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HeartDialDialogue{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HeartDialDialogueAccessor.LoadData to preload everything
func (a *HeartDialDialogueAccessor) GroupData() {
	a._dataID = map[float64]HeartDialDialogue{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the HeartDialDialogue uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *HeartDialDialogueAccessor) ByID(identifier float64) (HeartDialDialogue, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeartDialDialogue{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
