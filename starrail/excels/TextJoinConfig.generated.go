package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TextJoinConfig struct {
	DefaultItem      float64   `json:"DefaultItem"`
	IsOverride       bool      `json:"IsOverride"`
	TextJoinID       float64   `json:"TextJoinID"`
	TextJoinItemList []float64 `json:"TextJoinItemList"`
	Type             string    `json:"Type"`
}
type TextJoinConfigAccessor struct {
	_data           []TextJoinConfig
	_dataTextJoinID map[float64]TextJoinConfig
}

// LoadData retrieves the data. Must be called before TextJoinConfig.GroupData
func (a *TextJoinConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TextJoinConfig.json")
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
func (a *TextJoinConfigAccessor) Raw() ([]TextJoinConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TextJoinConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TextJoinConfigAccessor.LoadData to preload everything
func (a *TextJoinConfigAccessor) GroupData() {
	a._dataTextJoinID = map[float64]TextJoinConfig{}
	for _, d := range a._data {
		a._dataTextJoinID[d.TextJoinID] = d
	}
}

// ByTextJoinID returns the TextJoinConfig uniquely identified by TextJoinID
//
// Error is only non-nil if the source errors out
func (a *TextJoinConfigAccessor) ByTextJoinID(identifier float64) (TextJoinConfig, error) {
	if a._dataTextJoinID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TextJoinConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTextJoinID[identifier], nil
}
