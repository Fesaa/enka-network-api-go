package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MuseumItem struct {
	AreaID             float64   `json:"AreaID"`
	CollectedReward    float64   `json:"CollectedReward"`
	DisplayOrder       float64   `json:"DisplayOrder"`
	EvidenceInfoTextID hash.Hash `json:"EvidenceInfoTextID"`
	HideGetHint        bool      `json:"HideGetHint"`
	ItemID             float64   `json:"ItemID"`
	ItemSkillList      []float64 `json:"ItemSkillList"`
	MuseumItemDesc     hash.Hash `json:"MuseumItemDesc"`
	RenewPoint         float64   `json:"RenewPoint"`
	SceneGroupID       float64   `json:"SceneGroupID"`
	ScenePropID        float64   `json:"ScenePropID"`
	UnlockPhase        float64   `json:"UnlockPhase"`
}
type MuseumItemAccessor struct {
	_data       []MuseumItem
	_dataItemID map[float64]MuseumItem
}

// LoadData retrieves the data. Must be called before MuseumItem.GroupData
func (a *MuseumItemAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MuseumItem.json")
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
func (a *MuseumItemAccessor) Raw() ([]MuseumItem, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MuseumItem{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MuseumItemAccessor.LoadData to preload everything
func (a *MuseumItemAccessor) GroupData() {
	for _, d := range a._data {
		a._dataItemID[d.ItemID] = d
	}
}

// ByItemID returns the MuseumItem uniquely identified by ItemID
//
// Error is only non-nil if the source errors out
func (a *MuseumItemAccessor) ByItemID(identifier float64) (MuseumItem, error) {
	if a._dataItemID == nil {
		err := a.LoadData()
		if err != nil {
			return MuseumItem{}, err
		}
		a.GroupData()
	}
	return a._dataItemID[identifier], nil
}
