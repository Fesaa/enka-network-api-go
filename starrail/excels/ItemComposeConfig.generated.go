package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ItemComposeConfig struct {
	CoinCost                  float64                         `json:"CoinCost"`
	FormulaRequire            float64                         `json:"FormulaRequire"`
	FormulaType               string                          `json:"FormulaType"`
	FuncType                  string                          `json:"FuncType"`
	ID                        float64                         `json:"ID"`
	IsShowHoldNumber          bool                            `json:"IsShowHoldNumber"`
	ItemComposeTag            []float64                       `json:"ItemComposeTag"`
	ItemID                    float64                         `json:"ItemID"`
	LimitType                 string                          `json:"LimitType"`
	LimitValue                float64                         `json:"LimitValue"`
	MaterialCost              []ItemComposeConfigMaterialCost `json:"MaterialCost"`
	MaxCount                  float64                         `json:"MaxCount"`
	Order                     float64                         `json:"Order"`
	RelicList                 []float64                       `json:"RelicList"`
	SpecialMaterialCost       []float64                       `json:"SpecialMaterialCost"`
	SpecialMaterialCostNumber float64                         `json:"SpecialMaterialCostNumber"`
	Type                      float64                         `json:"Type"`
	WorldLevelRequire         float64                         `json:"WorldLevelRequire"`
}
type ItemComposeConfigMaterialCost struct {
	ItemID  float64 `json:"ItemID"`
	ItemNum float64 `json:"ItemNum"`
}
type ItemComposeConfigAccessor struct {
	_data   []ItemComposeConfig
	_dataID map[float64]ItemComposeConfig
}

// LoadData retrieves the data. Must be called before ItemComposeConfig.GroupData
func (a *ItemComposeConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemComposeConfig.json")
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
func (a *ItemComposeConfigAccessor) Raw() ([]ItemComposeConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemComposeConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ItemComposeConfigAccessor.LoadData to preload everything
func (a *ItemComposeConfigAccessor) GroupData() {
	a._dataID = map[float64]ItemComposeConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the ItemComposeConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ItemComposeConfigAccessor) ByID(identifier float64) (ItemComposeConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemComposeConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
