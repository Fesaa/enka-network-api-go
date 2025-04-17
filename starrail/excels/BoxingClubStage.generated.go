package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type BoxingClubStage struct {
	BubbleTalkEnemy  hash.Hash `json:"BubbleTalkEnemy"`
	BubbleTalkPlayer hash.Hash `json:"BubbleTalkPlayer"`
	BuffID           float64   `json:"BuffID"`
	BuffOptionalList []float64 `json:"BuffOptionalList"`
	EventID          float64   `json:"EventID"`
	MonsterWaveIndex float64   `json:"MonsterWaveIndex"`
	Name             hash.Hash `json:"Name"`
}
type BoxingClubStageAccessor struct {
	_data        []BoxingClubStage
	_dataEventID map[float64]BoxingClubStage
}

// LoadData retrieves the data. Must be called before BoxingClubStage.GroupData
func (a *BoxingClubStageAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BoxingClubStage.json")
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
func (a *BoxingClubStageAccessor) Raw() ([]BoxingClubStage, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BoxingClubStage{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BoxingClubStageAccessor.LoadData to preload everything
func (a *BoxingClubStageAccessor) GroupData() {
	a._dataEventID = map[float64]BoxingClubStage{}
	for _, d := range a._data {
		a._dataEventID[d.EventID] = d
	}
}

// ByEventID returns the BoxingClubStage uniquely identified by EventID
//
// Error is only non-nil if the source errors out
func (a *BoxingClubStageAccessor) ByEventID(identifier float64) (BoxingClubStage, error) {
	if a._dataEventID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BoxingClubStage{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEventID[identifier], nil
}
