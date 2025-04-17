package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MissionGotoConfig struct {
	Desc   hash.Hash `json:"Desc"`
	GotoID float64   `json:"GotoID"`
}
type MissionGotoConfigAccessor struct {
	_data       []MissionGotoConfig
	_dataGotoID map[float64]MissionGotoConfig
}

// LoadData retrieves the data. Must be called before MissionGotoConfig.GroupData
func (a *MissionGotoConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MissionGotoConfig.json")
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
func (a *MissionGotoConfigAccessor) Raw() ([]MissionGotoConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MissionGotoConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MissionGotoConfigAccessor.LoadData to preload everything
func (a *MissionGotoConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGotoID[d.GotoID] = d
	}
}

// ByGotoID returns the MissionGotoConfig uniquely identified by GotoID
//
// Error is only non-nil if the source errors out
func (a *MissionGotoConfigAccessor) ByGotoID(identifier float64) (MissionGotoConfig, error) {
	if a._dataGotoID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MissionGotoConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGotoID[identifier], nil
}
