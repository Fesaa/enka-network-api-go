package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesGachaBasic struct {
	CostGemNum             float64   `json:"CostGemNum"`
	CostItemID             float64   `json:"CostItemID"`
	GachaID                float64   `json:"GachaID"`
	GachaType              string    `json:"GachaType"`
	MultiGachaCount        float64   `json:"MultiGachaCount"`
	MultiGachaUnlockIDList []float64 `json:"MultiGachaUnlockIDList"`
	UnlockIDList           []float64 `json:"UnlockIDList"`
}
type PlanetFesGachaBasicAccessor struct {
	_data        []PlanetFesGachaBasic
	_dataGachaID map[float64]PlanetFesGachaBasic
}

// LoadData retrieves the data. Must be called before PlanetFesGachaBasic.GroupData
func (a *PlanetFesGachaBasicAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesGachaBasic.json")
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
func (a *PlanetFesGachaBasicAccessor) Raw() ([]PlanetFesGachaBasic, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesGachaBasic{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesGachaBasicAccessor.LoadData to preload everything
func (a *PlanetFesGachaBasicAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGachaID[d.GachaID] = d
	}
}

// ByGachaID returns the PlanetFesGachaBasic uniquely identified by GachaID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesGachaBasicAccessor) ByGachaID(identifier float64) (PlanetFesGachaBasic, error) {
	if a._dataGachaID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesGachaBasic{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGachaID[identifier], nil
}
