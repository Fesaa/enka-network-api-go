package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ExtraEffectConfig struct {
	DescParamList       []hash.IntValue `json:"DescParamList"`
	ExtraEffectDesc     hash.Hash       `json:"ExtraEffectDesc"`
	ExtraEffectID       float64         `json:"ExtraEffectID"`
	ExtraEffectIconPath string          `json:"ExtraEffectIconPath"`
	ExtraEffectName     hash.Hash       `json:"ExtraEffectName"`
	ExtraEffectType     float64         `json:"ExtraEffectType"`
}
type ExtraEffectConfigAccessor struct {
	_data              []ExtraEffectConfig
	_dataExtraEffectID map[float64]ExtraEffectConfig
}

// LoadData retrieves the data. Must be called before ExtraEffectConfig.GroupData
func (a *ExtraEffectConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ExtraEffectConfig.json")
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
func (a *ExtraEffectConfigAccessor) Raw() ([]ExtraEffectConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ExtraEffectConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ExtraEffectConfigAccessor.LoadData to preload everything
func (a *ExtraEffectConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataExtraEffectID[d.ExtraEffectID] = d
	}
}

// ByExtraEffectID returns the ExtraEffectConfig uniquely identified by ExtraEffectID
//
// Error is only non-nil if the source errors out
func (a *ExtraEffectConfigAccessor) ByExtraEffectID(identifier float64) (ExtraEffectConfig, error) {
	if a._dataExtraEffectID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ExtraEffectConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataExtraEffectID[identifier], nil
}
