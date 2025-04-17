package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AlleyGrid struct {
	GridDesc           hash.Hash `json:"GridDesc"`
	GridID             float64   `json:"GridID"`
	GridIcon           string    `json:"GridIcon"`
	GridTitle          hash.Hash `json:"GridTitle"`
	GridType           string    `json:"GridType"`
	RelatedEventID     float64   `json:"RelatedEventID"`
	RelatedMainMission float64   `json:"RelatedMainMission"`
	ShopInfoIcon       string    `json:"ShopInfoIcon"`
}
type AlleyGridAccessor struct {
	_data       []AlleyGrid
	_dataGridID map[float64]AlleyGrid
}

// LoadData retrieves the data. Must be called before AlleyGrid.GroupData
func (a *AlleyGridAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AlleyGrid.json")
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
func (a *AlleyGridAccessor) Raw() ([]AlleyGrid, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AlleyGrid{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AlleyGridAccessor.LoadData to preload everything
func (a *AlleyGridAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGridID[d.GridID] = d
	}
}

// ByGridID returns the AlleyGrid uniquely identified by GridID
//
// Error is only non-nil if the source errors out
func (a *AlleyGridAccessor) ByGridID(identifier float64) (AlleyGrid, error) {
	if a._dataGridID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AlleyGrid{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGridID[identifier], nil
}
