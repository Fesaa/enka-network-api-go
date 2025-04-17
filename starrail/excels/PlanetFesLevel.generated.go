package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesLevel struct {
	BasicBuffIDList []float64                     `json:"BasicBuffIDList"`
	BuffIDList      []float64                     `json:"BuffIDList"`
	CostNum         PlanetFesLevelCostNum         `json:"CostNum"`
	Description     hash.Hash                     `json:"Description"`
	GrantGold       PlanetFesLevelGrantGold       `json:"GrantGold"`
	GrantItemList   []PlanetFesLevelGrantItemList `json:"GrantItemList"`
	Level           float64                       `json:"Level"`
	NewTipsList     []float64                     `json:"NewTipsList"`
	QuestID         float64                       `json:"QuestID"`
}
type PlanetFesLevelCostNum struct {
	BaseValue float64 `json:"base_value"`
	Unit      string  `json:"unit"`
}
type PlanetFesLevelGrantGold struct {
	BaseValue float64 `json:"base_value"`
	Unit      string  `json:"unit"`
}
type PlanetFesLevelGrantItemList struct {
	ItemID  float64 `json:"ItemID"`
	ItemNum float64 `json:"ItemNum"`
}
type PlanetFesLevelAccessor struct {
	_data      []PlanetFesLevel
	_dataLevel map[float64]PlanetFesLevel
}

// LoadData retrieves the data. Must be called before PlanetFesLevel.GroupData
func (a *PlanetFesLevelAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesLevel.json")
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
func (a *PlanetFesLevelAccessor) Raw() ([]PlanetFesLevel, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesLevel{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesLevelAccessor.LoadData to preload everything
func (a *PlanetFesLevelAccessor) GroupData() {
	for _, d := range a._data {
		a._dataLevel[d.Level] = d
	}
}

// ByLevel returns the PlanetFesLevel uniquely identified by Level
//
// Error is only non-nil if the source errors out
func (a *PlanetFesLevelAccessor) ByLevel(identifier float64) (PlanetFesLevel, error) {
	if a._dataLevel == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesLevel{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLevel[identifier], nil
}
