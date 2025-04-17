package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type VersionReviewMission struct {
	PreMainMissionID     float64 `json:"PreMainMissionID"`
	ReviewMainMissionID  float64 `json:"ReviewMainMissionID"`
	StoryPerformanceID   float64 `json:"StoryPerformanceID"`
	StoryStartEntranceID float64 `json:"StoryStartEntranceID"`
}
type VersionReviewMissionAccessor struct {
	_data                     []VersionReviewMission
	_dataReviewMainMissionID  map[float64]VersionReviewMission
	_dataPreMainMissionID     map[float64]VersionReviewMission
	_dataStoryPerformanceID   map[float64]VersionReviewMission
	_dataStoryStartEntranceID map[float64]VersionReviewMission
}

// LoadData retrieves the data. Must be called before VersionReviewMission.GroupData
func (a *VersionReviewMissionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/VersionReviewMission.json")
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
func (a *VersionReviewMissionAccessor) Raw() ([]VersionReviewMission, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []VersionReviewMission{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with VersionReviewMissionAccessor.LoadData to preload everything
func (a *VersionReviewMissionAccessor) GroupData() {
	for _, d := range a._data {
		a._dataReviewMainMissionID[d.ReviewMainMissionID] = d
		a._dataPreMainMissionID[d.PreMainMissionID] = d
		a._dataStoryPerformanceID[d.StoryPerformanceID] = d
		a._dataStoryStartEntranceID[d.StoryStartEntranceID] = d
	}
}

// ByReviewMainMissionID returns the VersionReviewMission uniquely identified by ReviewMainMissionID
//
// Error is only non-nil if the source errors out
func (a *VersionReviewMissionAccessor) ByReviewMainMissionID(identifier float64) (VersionReviewMission, error) {
	if a._dataReviewMainMissionID == nil {
		err := a.LoadData()
		if err != nil {
			return VersionReviewMission{}, err
		}
		a.GroupData()
	}
	return a._dataReviewMainMissionID[identifier], nil
}

// ByPreMainMissionID returns the VersionReviewMission uniquely identified by PreMainMissionID
//
// Error is only non-nil if the source errors out
func (a *VersionReviewMissionAccessor) ByPreMainMissionID(identifier float64) (VersionReviewMission, error) {
	if a._dataPreMainMissionID == nil {
		err := a.LoadData()
		if err != nil {
			return VersionReviewMission{}, err
		}
		a.GroupData()
	}
	return a._dataPreMainMissionID[identifier], nil
}

// ByStoryPerformanceID returns the VersionReviewMission uniquely identified by StoryPerformanceID
//
// Error is only non-nil if the source errors out
func (a *VersionReviewMissionAccessor) ByStoryPerformanceID(identifier float64) (VersionReviewMission, error) {
	if a._dataStoryPerformanceID == nil {
		err := a.LoadData()
		if err != nil {
			return VersionReviewMission{}, err
		}
		a.GroupData()
	}
	return a._dataStoryPerformanceID[identifier], nil
}

// ByStoryStartEntranceID returns the VersionReviewMission uniquely identified by StoryStartEntranceID
//
// Error is only non-nil if the source errors out
func (a *VersionReviewMissionAccessor) ByStoryStartEntranceID(identifier float64) (VersionReviewMission, error) {
	if a._dataStoryStartEntranceID == nil {
		err := a.LoadData()
		if err != nil {
			return VersionReviewMission{}, err
		}
		a.GroupData()
	}
	return a._dataStoryStartEntranceID[identifier], nil
}
