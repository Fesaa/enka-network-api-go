package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type SwordTrainingEventOption struct {
	EffectIDList []json.Number          `json:"EffectIDList"`
	OptionDesc   map[string]json.Number `json:"OptionDesc"`
	OptionID     json.Number            `json:"OptionID"`
	ResultAudio  string                 `json:"ResultAudio"`
	ResultDesc   map[string]json.Number `json:"ResultDesc"`
	ResultImage  string                 `json:"ResultImage"`
}
type SwordTrainingEventOptionAccessor struct {
	_data         []SwordTrainingEventOption
	_dataOptionID map[json.Number]SwordTrainingEventOption
}

// LoadData retrieves the data. Must be called before SwordTrainingEventOption.GroupData
func (a *SwordTrainingEventOptionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SwordTrainingEventOption.json")
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
func (a *SwordTrainingEventOptionAccessor) Raw() ([]SwordTrainingEventOption, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SwordTrainingEventOption{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SwordTrainingEventOptionAccessor.LoadData to preload everything
func (a *SwordTrainingEventOptionAccessor) GroupData() {
	a._dataOptionID = map[json.Number]SwordTrainingEventOption{}
	for _, d := range a._data {
		a._dataOptionID[d.OptionID] = d
	}
}

// ByOptionID returns the SwordTrainingEventOption uniquely identified by OptionID
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingEventOptionAccessor) ByOptionID(identifier json.Number) (SwordTrainingEventOption, error) {
	if a._dataOptionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SwordTrainingEventOption{}, err
			}
		}
		a.GroupData()
	}
	return a._dataOptionID[identifier], nil
}
