package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournTitanBless struct {
	BlessBattleDisplayCategoryList []string  `json:"BlessBattleDisplayCategoryList"`
	BlessRatio                     float64   `json:"BlessRatio"`
	ExtraEffectIDList              []float64 `json:"ExtraEffectIDList"`
	MazeBuffID                     float64   `json:"MazeBuffID"`
	TitanBlessID                   float64   `json:"TitanBlessID"`
	TitanBlessLevel                float64   `json:"TitanBlessLevel"`
	TitanType                      string    `json:"TitanType"`
}
type RogueTournTitanBlessAccessor struct {
	_data             []RogueTournTitanBless
	_dataMazeBuffID   map[float64]RogueTournTitanBless
	_dataTitanBlessID map[float64]RogueTournTitanBless
}

// LoadData retrieves the data. Must be called before RogueTournTitanBless.GroupData
func (a *RogueTournTitanBlessAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournTitanBless.json")
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
func (a *RogueTournTitanBlessAccessor) Raw() ([]RogueTournTitanBless, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournTitanBless{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournTitanBlessAccessor.LoadData to preload everything
func (a *RogueTournTitanBlessAccessor) GroupData() {
	a._dataMazeBuffID = map[float64]RogueTournTitanBless{}
	a._dataTitanBlessID = map[float64]RogueTournTitanBless{}
	for _, d := range a._data {
		a._dataMazeBuffID[d.MazeBuffID] = d
		a._dataTitanBlessID[d.TitanBlessID] = d
	}
}

// ByMazeBuffID returns the RogueTournTitanBless uniquely identified by MazeBuffID
//
// Error is only non-nil if the source errors out
func (a *RogueTournTitanBlessAccessor) ByMazeBuffID(identifier float64) (RogueTournTitanBless, error) {
	if a._dataMazeBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournTitanBless{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMazeBuffID[identifier], nil
}

// ByTitanBlessID returns the RogueTournTitanBless uniquely identified by TitanBlessID
//
// Error is only non-nil if the source errors out
func (a *RogueTournTitanBlessAccessor) ByTitanBlessID(identifier float64) (RogueTournTitanBless, error) {
	if a._dataTitanBlessID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournTitanBless{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTitanBlessID[identifier], nil
}
