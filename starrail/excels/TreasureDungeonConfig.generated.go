package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TreasureDungeonConfig struct {
	Desc                 hash.Hash `json:"Desc"`
	DisplayEventID       float64   `json:"DisplayEventID"`
	DisplayMonsterIDList []float64 `json:"DisplayMonsterIDList"`
	DungeonID            float64   `json:"DungeonID"`
	EntranceIconPath     string    `json:"EntranceIconPath"`
	ExploreSubHpRatio    float64   `json:"ExploreSubHpRatio"`
	GridExploreCost      float64   `json:"GridExploreCost"`
	GridPrefabType       float64   `json:"GridPrefabType"`
	GroupID              float64   `json:"GroupID"`
	ImgPath              string    `json:"ImgPath"`
	InitialExplore       float64   `json:"InitialExplore"`
	MaxExplore           float64   `json:"MaxExplore"`
	Name                 hash.Hash `json:"Name"`
	PreDungeonID         float64   `json:"PreDungeonID"`
	RecommendNature      []string  `json:"RecommendNature"`
	SpecialAvatarIDList  []float64 `json:"SpecialAvatarIDList"`
	UnlockID             float64   `json:"UnlockID"`
}
type TreasureDungeonConfigAccessor struct {
	_data                 []TreasureDungeonConfig
	_dataDungeonID        map[float64]TreasureDungeonConfig
	_dataEntranceIconPath map[string]TreasureDungeonConfig
}

// LoadData retrieves the data. Must be called before TreasureDungeonConfig.GroupData
func (a *TreasureDungeonConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TreasureDungeonConfig.json")
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
func (a *TreasureDungeonConfigAccessor) Raw() ([]TreasureDungeonConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TreasureDungeonConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TreasureDungeonConfigAccessor.LoadData to preload everything
func (a *TreasureDungeonConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataDungeonID[d.DungeonID] = d
		a._dataEntranceIconPath[d.EntranceIconPath] = d
	}
}

// ByDungeonID returns the TreasureDungeonConfig uniquely identified by DungeonID
//
// Error is only non-nil if the source errors out
func (a *TreasureDungeonConfigAccessor) ByDungeonID(identifier float64) (TreasureDungeonConfig, error) {
	if a._dataDungeonID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TreasureDungeonConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDungeonID[identifier], nil
}

// ByEntranceIconPath returns the TreasureDungeonConfig uniquely identified by EntranceIconPath
//
// Error is only non-nil if the source errors out
func (a *TreasureDungeonConfigAccessor) ByEntranceIconPath(identifier string) (TreasureDungeonConfig, error) {
	if a._dataEntranceIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TreasureDungeonConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEntranceIconPath[identifier], nil
}
