package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCSubStoryGroup struct {
	ShowGroup         float64   `json:"ShowGroup"`
	SubStoryGroupID   float64   `json:"SubStoryGroupID"`
	SubStoryGroupName hash.Hash `json:"SubStoryGroupName"`
	SubStoryList      []float64 `json:"SubStoryList"`
	UnlockID          float64   `json:"UnlockID"`
}
type RogueDLCSubStoryGroupAccessor struct {
	_data                []RogueDLCSubStoryGroup
	_dataSubStoryGroupID map[float64]RogueDLCSubStoryGroup
}

// LoadData retrieves the data. Must be called before RogueDLCSubStoryGroup.GroupData
func (a *RogueDLCSubStoryGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCSubStoryGroup.json")
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
func (a *RogueDLCSubStoryGroupAccessor) Raw() ([]RogueDLCSubStoryGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCSubStoryGroup{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCSubStoryGroupAccessor.LoadData to preload everything
func (a *RogueDLCSubStoryGroupAccessor) GroupData() {
	for _, d := range a._data {
		a._dataSubStoryGroupID[d.SubStoryGroupID] = d
	}
}

// BySubStoryGroupID returns the RogueDLCSubStoryGroup uniquely identified by SubStoryGroupID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCSubStoryGroupAccessor) BySubStoryGroupID(identifier float64) (RogueDLCSubStoryGroup, error) {
	if a._dataSubStoryGroupID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueDLCSubStoryGroup{}, err
		}
		a.GroupData()
	}
	return a._dataSubStoryGroupID[identifier], nil
}
