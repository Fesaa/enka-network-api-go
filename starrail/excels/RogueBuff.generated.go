package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueBuff struct {
	ActivityModuleID    float64   `json:"ActivityModuleID"`
	AeonCrossIcon       string    `json:"AeonCrossIcon"`
	AeonID              float64   `json:"AeonID"`
	BattleEventBuffType string    `json:"BattleEventBuffType"`
	ExtraEffectIDList   []float64 `json:"ExtraEffectIDList"`
	HandbookUnlockDesc  hash.Hash `json:"HandbookUnlockDesc"`
	IsShow              bool      `json:"IsShow"`
	MazeBuffID          float64   `json:"MazeBuffID"`
	MazeBuffLevel       float64   `json:"MazeBuffLevel"`
	RogueBuffCategory   string    `json:"RogueBuffCategory"`
	RogueBuffTag        float64   `json:"RogueBuffTag"`
	RogueBuffType       float64   `json:"RogueBuffType"`
	RogueVersion        float64   `json:"RogueVersion"`
	UnlockIDList        []float64 `json:"UnlockIDList"`
}
type RogueBuffAccessor struct {
	_data             []RogueBuff
	_dataRogueBuffTag map[float64]RogueBuff
}

// LoadData retrieves the data. Must be called before RogueBuff.GroupData
func (a *RogueBuffAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueBuff.json")
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
func (a *RogueBuffAccessor) Raw() ([]RogueBuff, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueBuff{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueBuffAccessor.LoadData to preload everything
func (a *RogueBuffAccessor) GroupData() {
	a._dataRogueBuffTag = map[float64]RogueBuff{}
	for _, d := range a._data {
		a._dataRogueBuffTag[d.RogueBuffTag] = d
	}
}

// ByRogueBuffTag returns the RogueBuff uniquely identified by RogueBuffTag
//
// Error is only non-nil if the source errors out
func (a *RogueBuffAccessor) ByRogueBuffTag(identifier float64) (RogueBuff, error) {
	if a._dataRogueBuffTag == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueBuff{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueBuffTag[identifier], nil
}
