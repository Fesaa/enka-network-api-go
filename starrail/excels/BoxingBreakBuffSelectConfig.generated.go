package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type BoxingBreakBuffSelectConfig struct {
	BoxingClubBuffID     float64   `json:"BoxingClubBuffID"`
	BoxingClubNatureType string    `json:"BoxingClubNatureType"`
	ExtraEffectIDList    []float64 `json:"ExtraEffectIDList"`
}
type BoxingBreakBuffSelectConfigAccessor struct {
	_data                 []BoxingBreakBuffSelectConfig
	_dataBoxingClubBuffID map[float64]BoxingBreakBuffSelectConfig
}

// LoadData retrieves the data. Must be called before BoxingBreakBuffSelectConfig.GroupData
func (a *BoxingBreakBuffSelectConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BoxingBreakBuffSelectConfig.json")
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
func (a *BoxingBreakBuffSelectConfigAccessor) Raw() ([]BoxingBreakBuffSelectConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BoxingBreakBuffSelectConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BoxingBreakBuffSelectConfigAccessor.LoadData to preload everything
func (a *BoxingBreakBuffSelectConfigAccessor) GroupData() {
	a._dataBoxingClubBuffID = map[float64]BoxingBreakBuffSelectConfig{}
	for _, d := range a._data {
		a._dataBoxingClubBuffID[d.BoxingClubBuffID] = d
	}
}

// ByBoxingClubBuffID returns the BoxingBreakBuffSelectConfig uniquely identified by BoxingClubBuffID
//
// Error is only non-nil if the source errors out
func (a *BoxingBreakBuffSelectConfigAccessor) ByBoxingClubBuffID(identifier float64) (BoxingBreakBuffSelectConfig, error) {
	if a._dataBoxingClubBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BoxingBreakBuffSelectConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBoxingClubBuffID[identifier], nil
}
