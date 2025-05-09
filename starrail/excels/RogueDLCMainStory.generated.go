package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCMainStory struct {
	BonusToast          hash.Hash `json:"BonusToast"`
	IsBonusUnlock       bool      `json:"IsBonusUnlock"`
	Layer               float64   `json:"Layer"`
	MainStoryButtonIcon string    `json:"MainStoryButtonIcon"`
	MainStoryID         float64   `json:"MainStoryID"`
	MainStoryName       hash.Hash `json:"MainStoryName"`
	MainStoryToastType  string    `json:"MainStoryToastType"`
	UnlockAeonDimension float64   `json:"UnlockAeonDimension"`
	UnlockPoint         float64   `json:"UnlockPoint"`
}
type RogueDLCMainStoryAccessor struct {
	_data                    []RogueDLCMainStory
	_dataMainStoryButtonIcon map[string]RogueDLCMainStory
	_dataMainStoryID         map[float64]RogueDLCMainStory
}

// LoadData retrieves the data. Must be called before RogueDLCMainStory.GroupData
func (a *RogueDLCMainStoryAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCMainStory.json")
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
func (a *RogueDLCMainStoryAccessor) Raw() ([]RogueDLCMainStory, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCMainStory{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCMainStoryAccessor.LoadData to preload everything
func (a *RogueDLCMainStoryAccessor) GroupData() {
	a._dataMainStoryButtonIcon = map[string]RogueDLCMainStory{}
	a._dataMainStoryID = map[float64]RogueDLCMainStory{}
	for _, d := range a._data {
		a._dataMainStoryButtonIcon[d.MainStoryButtonIcon] = d
		a._dataMainStoryID[d.MainStoryID] = d
	}
}

// ByMainStoryButtonIcon returns the RogueDLCMainStory uniquely identified by MainStoryButtonIcon
//
// Error is only non-nil if the source errors out
func (a *RogueDLCMainStoryAccessor) ByMainStoryButtonIcon(identifier string) (RogueDLCMainStory, error) {
	if a._dataMainStoryButtonIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCMainStory{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMainStoryButtonIcon[identifier], nil
}

// ByMainStoryID returns the RogueDLCMainStory uniquely identified by MainStoryID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCMainStoryAccessor) ByMainStoryID(identifier float64) (RogueDLCMainStory, error) {
	if a._dataMainStoryID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCMainStory{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMainStoryID[identifier], nil
}
