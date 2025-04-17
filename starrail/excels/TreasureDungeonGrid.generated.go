package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TreasureDungeonGrid struct {
	EffectType            string    `json:"EffectType"`
	GridID                float64   `json:"GridID"`
	GridSubType           float64   `json:"GridSubType"`
	GridType              string    `json:"GridType"`
	IconPath              string    `json:"IconPath"`
	IconPath2D            string    `json:"IconPath2D"`
	Name                  hash.Hash `json:"Name"`
	OpenBuff              float64   `json:"OpenBuff"`
	ParamInt              float64   `json:"ParamInt"`
	ReplaceGridID         float64   `json:"ReplaceGridID"`
	TutorialTriggerString string    `json:"TutorialTriggerString"`
	TutorialTriggerType   float64   `json:"TutorialTriggerType"`
	TypeParam             []float64 `json:"TypeParam"`
}
type TreasureDungeonGridAccessor struct {
	_data []TreasureDungeonGrid
}

// LoadData retrieves the data. Must be called before TreasureDungeonGrid.GroupData
func (a *TreasureDungeonGridAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TreasureDungeonGrid.json")
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
func (a *TreasureDungeonGridAccessor) Raw() ([]TreasureDungeonGrid, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TreasureDungeonGrid{}, err
		}
	}
	return a._data, nil
}
