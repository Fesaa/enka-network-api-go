package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MuseumStuff struct {
	CollectedReward      float64   `json:"CollectedReward"`
	DisplayOrder         float64   `json:"DisplayOrder"`
	EvidenceInfoTextID   hash.Hash `json:"EvidenceInfoTextID"`
	IsInitial            bool      `json:"IsInitial"`
	IsTargetReward       bool      `json:"IsTargetReward"`
	ItemID               float64   `json:"ItemID"`
	MuseumStuffDesc      hash.Hash `json:"MuseumStuffDesc"`
	RecruitPrice         float64   `json:"RecruitPrice"`
	RecruitUnlockMission float64   `json:"RecruitUnlockMission"`
	SceneGroupID         float64   `json:"SceneGroupID"`
	ScenePropID          float64   `json:"ScenePropID"`
	StatsA               float64   `json:"StatsA"`
	StatsB               float64   `json:"StatsB"`
	StatsC               float64   `json:"StatsC"`
	Type                 string    `json:"Type"`
	UnlockPhase          float64   `json:"UnlockPhase"`
}
type MuseumStuffAccessor struct {
	_data            []MuseumStuff
	_dataScenePropID map[float64]MuseumStuff
	_dataItemID      map[float64]MuseumStuff
}

// LoadData retrieves the data. Must be called before MuseumStuff.GroupData
func (a *MuseumStuffAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MuseumStuff.json")
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
func (a *MuseumStuffAccessor) Raw() ([]MuseumStuff, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MuseumStuff{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MuseumStuffAccessor.LoadData to preload everything
func (a *MuseumStuffAccessor) GroupData() {
	for _, d := range a._data {
		a._dataScenePropID[d.ScenePropID] = d
		a._dataItemID[d.ItemID] = d
	}
}

// ByScenePropID returns the MuseumStuff uniquely identified by ScenePropID
//
// Error is only non-nil if the source errors out
func (a *MuseumStuffAccessor) ByScenePropID(identifier float64) (MuseumStuff, error) {
	if a._dataScenePropID == nil {
		err := a.LoadData()
		if err != nil {
			return MuseumStuff{}, err
		}
		a.GroupData()
	}
	return a._dataScenePropID[identifier], nil
}

// ByItemID returns the MuseumStuff uniquely identified by ItemID
//
// Error is only non-nil if the source errors out
func (a *MuseumStuffAccessor) ByItemID(identifier float64) (MuseumStuff, error) {
	if a._dataItemID == nil {
		err := a.LoadData()
		if err != nil {
			return MuseumStuff{}, err
		}
		a.GroupData()
	}
	return a._dataItemID[identifier], nil
}
