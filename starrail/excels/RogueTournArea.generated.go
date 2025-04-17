package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournArea struct {
	AreaGroupID        string                  `json:"AreaGroupID"`
	AreaID             float64                 `json:"AreaID"`
	AreaNameID         hash.Hash               `json:"AreaNameID"`
	Difficulty         string                  `json:"Difficulty"`
	DifficultyIDList   []float64               `json:"DifficultyIDList"`
	DivisionLevel      float64                 `json:"DivisionLevel"`
	ExpScoreID         float64                 `json:"ExpScoreID"`
	FirstReward        float64                 `json:"FirstReward"`
	FirstRoom          RogueTournAreaFirstRoom `json:"FirstRoom"`
	IsHard             bool                    `json:"IsHard"`
	LayerIDList        []float64               `json:"LayerIDList"`
	MonsterEliteDropID float64                 `json:"MonsterEliteDropID"`
	TournMode          string                  `json:"TournMode"`
	UnlockID           float64                 `json:"UnlockID"`
	WorldLevelLimit    float64                 `json:"WorldLevelLimit"`
}
type RogueTournAreaFirstRoom struct {
	NBOCIPLJMHI string `json:"NBOCIPLJMHI"`
}
type RogueTournAreaAccessor struct {
	_data       []RogueTournArea
	_dataAreaID map[float64]RogueTournArea
}

// LoadData retrieves the data. Must be called before RogueTournArea.GroupData
func (a *RogueTournAreaAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournArea.json")
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
func (a *RogueTournAreaAccessor) Raw() ([]RogueTournArea, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournArea{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournAreaAccessor.LoadData to preload everything
func (a *RogueTournAreaAccessor) GroupData() {
	for _, d := range a._data {
		a._dataAreaID[d.AreaID] = d
	}
}

// ByAreaID returns the RogueTournArea uniquely identified by AreaID
//
// Error is only non-nil if the source errors out
func (a *RogueTournAreaAccessor) ByAreaID(identifier float64) (RogueTournArea, error) {
	if a._dataAreaID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournArea{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAreaID[identifier], nil
}
