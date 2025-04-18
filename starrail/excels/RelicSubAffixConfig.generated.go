package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RelicSubAffixConfig struct {
	AffixID   float64    `json:"AffixID"`
	BaseValue hash.Value `json:"BaseValue"`
	GroupID   float64    `json:"GroupID"`
	Property  string     `json:"Property"`
	StepNum   float64    `json:"StepNum"`
	StepValue hash.Value `json:"StepValue"`
}
type RelicSubAffixConfigAccessor struct {
	_data []RelicSubAffixConfig
}

// LoadData retrieves the data. Must be called before RelicSubAffixConfig.GroupData
func (a *RelicSubAffixConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RelicSubAffixConfig.json")
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
func (a *RelicSubAffixConfigAccessor) Raw() ([]RelicSubAffixConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RelicSubAffixConfig{}, err
		}
	}
	return a._data, nil
}
