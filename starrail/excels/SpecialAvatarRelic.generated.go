package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type SpecialAvatarRelic struct {
	Comment2          string                          `json:"Comment2"`
	RelicIDList       []SpecialAvatarRelicRelicIDList `json:"RelicIDList"`
	RelicPropertyType float64                         `json:"RelicPropertyType"`
}
type SpecialAvatarRelicRelicIDList struct {
	GOBEGPKDLLF float64 `json:"GOBEGPKDLLF"`
	JHPOAELKHED float64 `json:"JHPOAELKHED"`
}
type SpecialAvatarRelicAccessor struct {
	_data                  []SpecialAvatarRelic
	_dataRelicPropertyType map[float64]SpecialAvatarRelic
}

// LoadData retrieves the data. Must be called before SpecialAvatarRelic.GroupData
func (a *SpecialAvatarRelicAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SpecialAvatarRelic.json")
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
func (a *SpecialAvatarRelicAccessor) Raw() ([]SpecialAvatarRelic, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SpecialAvatarRelic{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SpecialAvatarRelicAccessor.LoadData to preload everything
func (a *SpecialAvatarRelicAccessor) GroupData() {
	a._dataRelicPropertyType = map[float64]SpecialAvatarRelic{}
	for _, d := range a._data {
		a._dataRelicPropertyType[d.RelicPropertyType] = d
	}
}

// ByRelicPropertyType returns the SpecialAvatarRelic uniquely identified by RelicPropertyType
//
// Error is only non-nil if the source errors out
func (a *SpecialAvatarRelicAccessor) ByRelicPropertyType(identifier float64) (SpecialAvatarRelic, error) {
	if a._dataRelicPropertyType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpecialAvatarRelic{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRelicPropertyType[identifier], nil
}
