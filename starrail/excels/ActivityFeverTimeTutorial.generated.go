package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityFeverTimeTutorial struct {
	P2AvailableBuffID   float64   `json:"P2AvailableBuffID"`
	RecommendAvatarList []float64 `json:"RecommendAvatarList"`
	TutorialID          float64   `json:"TutorialID"`
}
type ActivityFeverTimeTutorialAccessor struct {
	_data                  []ActivityFeverTimeTutorial
	_dataP2AvailableBuffID map[float64]ActivityFeverTimeTutorial
}

// LoadData retrieves the data. Must be called before ActivityFeverTimeTutorial.GroupData
func (a *ActivityFeverTimeTutorialAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityFeverTimeTutorial.json")
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
func (a *ActivityFeverTimeTutorialAccessor) Raw() ([]ActivityFeverTimeTutorial, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityFeverTimeTutorial{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityFeverTimeTutorialAccessor.LoadData to preload everything
func (a *ActivityFeverTimeTutorialAccessor) GroupData() {
	a._dataP2AvailableBuffID = map[float64]ActivityFeverTimeTutorial{}
	for _, d := range a._data {
		a._dataP2AvailableBuffID[d.P2AvailableBuffID] = d
	}
}

// ByP2AvailableBuffID returns the ActivityFeverTimeTutorial uniquely identified by P2AvailableBuffID
//
// Error is only non-nil if the source errors out
func (a *ActivityFeverTimeTutorialAccessor) ByP2AvailableBuffID(identifier float64) (ActivityFeverTimeTutorial, error) {
	if a._dataP2AvailableBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityFeverTimeTutorial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataP2AvailableBuffID[identifier], nil
}
