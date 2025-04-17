package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ItemComposeConfig struct {
	CoinCost                  json.Number              `json:"CoinCost"`
	FormulaRequire            json.Number              `json:"FormulaRequire"`
	FormulaType               string                   `json:"FormulaType"`
	FuncType                  string                   `json:"FuncType"`
	ID                        json.Number              `json:"ID"`
	IsShowHoldNumber          bool                     `json:"IsShowHoldNumber"`
	ItemComposeTag            []json.Number            `json:"ItemComposeTag"`
	ItemID                    json.Number              `json:"ItemID"`
	LimitType                 string                   `json:"LimitType"`
	LimitValue                json.Number              `json:"LimitValue"`
	MaterialCost              []map[string]json.Number `json:"MaterialCost"`
	MaxCount                  json.Number              `json:"MaxCount"`
	Order                     json.Number              `json:"Order"`
	RelicList                 []json.Number            `json:"RelicList"`
	SpecialMaterialCost       []json.Number            `json:"SpecialMaterialCost"`
	SpecialMaterialCostNumber json.Number              `json:"SpecialMaterialCostNumber"`
	Type                      json.Number              `json:"Type"`
	WorldLevelRequire         json.Number              `json:"WorldLevelRequire"`
}
type ItemComposeConfigAccessor struct {
	_data []ItemComposeConfig
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
