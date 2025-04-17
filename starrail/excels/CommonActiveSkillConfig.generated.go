package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type CommonActiveSkillConfig struct {
	AbilityName         string  `json:"AbilityName"`
	CommonActiveSkillID float64 `json:"CommonActiveSkillID"`
}
type CommonActiveSkillConfigAccessor struct {
	_data                    []CommonActiveSkillConfig
	_dataCommonActiveSkillID map[float64]CommonActiveSkillConfig
}

// LoadData retrieves the data. Must be called before CommonActiveSkillConfig.GroupData
func (a *CommonActiveSkillConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/CommonActiveSkillConfig.json")
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
func (a *CommonActiveSkillConfigAccessor) Raw() ([]CommonActiveSkillConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []CommonActiveSkillConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with CommonActiveSkillConfigAccessor.LoadData to preload everything
func (a *CommonActiveSkillConfigAccessor) GroupData() {
	a._dataCommonActiveSkillID = map[float64]CommonActiveSkillConfig{}
	for _, d := range a._data {
		a._dataCommonActiveSkillID[d.CommonActiveSkillID] = d
	}
}

// ByCommonActiveSkillID returns the CommonActiveSkillConfig uniquely identified by CommonActiveSkillID
//
// Error is only non-nil if the source errors out
func (a *CommonActiveSkillConfigAccessor) ByCommonActiveSkillID(identifier float64) (CommonActiveSkillConfig, error) {
	if a._dataCommonActiveSkillID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return CommonActiveSkillConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCommonActiveSkillID[identifier], nil
}
