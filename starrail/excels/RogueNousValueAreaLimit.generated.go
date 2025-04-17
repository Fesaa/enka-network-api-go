package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueNousValueAreaLimit struct {
	AreaID       float64 `json:"AreaID"`
	MaxNousValue float64 `json:"MaxNousValue"`
	MinNousValue float64 `json:"MinNousValue"`
}
type RogueNousValueAreaLimitAccessor struct {
	_data       []RogueNousValueAreaLimit
	_dataAreaID map[float64]RogueNousValueAreaLimit
}

// LoadData retrieves the data. Must be called before RogueNousValueAreaLimit.GroupData
func (a *RogueNousValueAreaLimitAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueNousValueAreaLimit.json")
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
func (a *RogueNousValueAreaLimitAccessor) Raw() ([]RogueNousValueAreaLimit, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueNousValueAreaLimit{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueNousValueAreaLimitAccessor.LoadData to preload everything
func (a *RogueNousValueAreaLimitAccessor) GroupData() {
	for _, d := range a._data {
		a._dataAreaID[d.AreaID] = d
	}
}

// ByAreaID returns the RogueNousValueAreaLimit uniquely identified by AreaID
//
// Error is only non-nil if the source errors out
func (a *RogueNousValueAreaLimitAccessor) ByAreaID(identifier float64) (RogueNousValueAreaLimit, error) {
	if a._dataAreaID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueNousValueAreaLimit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAreaID[identifier], nil
}
