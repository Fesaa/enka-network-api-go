package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCMainStoryBranch struct {
	AeonID            float64 `json:"AeonID"`
	MainStoryBranchID float64 `json:"MainStoryBranchID"`
	RogueNPCID        float64 `json:"RogueNPCID"`
}
type RogueDLCMainStoryBranchAccessor struct {
	_data                  []RogueDLCMainStoryBranch
	_dataMainStoryBranchID map[float64]RogueDLCMainStoryBranch
}

// LoadData retrieves the data. Must be called before RogueDLCMainStoryBranch.GroupData
func (a *RogueDLCMainStoryBranchAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCMainStoryBranch.json")
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
func (a *RogueDLCMainStoryBranchAccessor) Raw() ([]RogueDLCMainStoryBranch, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCMainStoryBranch{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCMainStoryBranchAccessor.LoadData to preload everything
func (a *RogueDLCMainStoryBranchAccessor) GroupData() {
	a._dataMainStoryBranchID = map[float64]RogueDLCMainStoryBranch{}
	for _, d := range a._data {
		a._dataMainStoryBranchID[d.MainStoryBranchID] = d
	}
}

// ByMainStoryBranchID returns the RogueDLCMainStoryBranch uniquely identified by MainStoryBranchID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCMainStoryBranchAccessor) ByMainStoryBranchID(identifier float64) (RogueDLCMainStoryBranch, error) {
	if a._dataMainStoryBranchID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCMainStoryBranch{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMainStoryBranchID[identifier], nil
}
