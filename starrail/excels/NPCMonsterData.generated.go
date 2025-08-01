package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type NPCMonsterData struct {
	ConfigEntityPath string    `json:"ConfigEntityPath"`
	DefaultAIPath    string    `json:"DefaultAIPath"`
	ID               float64   `json:"ID"`
	IsMazeLink       bool      `json:"IsMazeLink"`
	JsonPath         string    `json:"JsonPath"`
	MappingInfoID    float64   `json:"MappingInfoID"`
	MiniMapIconType  float64   `json:"MiniMapIconType"`
	NPCName          hash.Hash `json:"NPCName"`
	PrototypeID      float64   `json:"PrototypeID"`
	Rank             string    `json:"Rank"`
}
type NPCMonsterDataAccessor struct {
	_data   []NPCMonsterData
	_dataID map[float64]NPCMonsterData
}

// LoadData retrieves the data. Must be called before NPCMonsterData.GroupData
func (a *NPCMonsterDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/NPCMonsterData.json")
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
func (a *NPCMonsterDataAccessor) Raw() ([]NPCMonsterData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []NPCMonsterData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with NPCMonsterDataAccessor.LoadData to preload everything
func (a *NPCMonsterDataAccessor) GroupData() {
	a._dataID = map[float64]NPCMonsterData{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the NPCMonsterData uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *NPCMonsterDataAccessor) ByID(identifier float64) (NPCMonsterData, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return NPCMonsterData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
