package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournNPC struct {
	NPCJsonPath string  `json:"NPCJsonPath"`
	RogueNPCID  float64 `json:"RogueNPCID"`
}
type RogueTournNPCAccessor struct {
	_data            []RogueTournNPC
	_dataNPCJsonPath map[string]RogueTournNPC
	_dataRogueNPCID  map[float64]RogueTournNPC
}

// LoadData retrieves the data. Must be called before RogueTournNPC.GroupData
func (a *RogueTournNPCAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournNPC.json")
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
func (a *RogueTournNPCAccessor) Raw() ([]RogueTournNPC, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournNPC{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournNPCAccessor.LoadData to preload everything
func (a *RogueTournNPCAccessor) GroupData() {
	for _, d := range a._data {
		a._dataNPCJsonPath[d.NPCJsonPath] = d
		a._dataRogueNPCID[d.RogueNPCID] = d
	}
}

// ByNPCJsonPath returns the RogueTournNPC uniquely identified by NPCJsonPath
//
// Error is only non-nil if the source errors out
func (a *RogueTournNPCAccessor) ByNPCJsonPath(identifier string) (RogueTournNPC, error) {
	if a._dataNPCJsonPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournNPC{}, err
			}
		}
		a.GroupData()
	}
	return a._dataNPCJsonPath[identifier], nil
}

// ByRogueNPCID returns the RogueTournNPC uniquely identified by RogueNPCID
//
// Error is only non-nil if the source errors out
func (a *RogueTournNPCAccessor) ByRogueNPCID(identifier float64) (RogueTournNPC, error) {
	if a._dataRogueNPCID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournNPC{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueNPCID[identifier], nil
}
