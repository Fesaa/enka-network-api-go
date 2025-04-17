package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type SysMailConfig struct {
	MailDetail   map[string]json.Number `json:"MailDetail"`
	MailID       json.Number            `json:"MailID"`
	MailLifeTime json.Number            `json:"MailLifeTime"`
	MailSender   map[string]json.Number `json:"MailSender"`
	MailTitle    map[string]json.Number `json:"MailTitle"`
	Type         string                 `json:"Type"`
}
type SysMailConfigAccessor struct {
	_data       []SysMailConfig
	_dataMailID map[json.Number]SysMailConfig
}

// LoadData retrieves the data. Must be called before SysMailConfig.GroupData
func (a *SysMailConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SysMailConfig.json")
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
func (a *SysMailConfigAccessor) Raw() ([]SysMailConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SysMailConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SysMailConfigAccessor.LoadData to preload everything
func (a *SysMailConfigAccessor) GroupData() {
	a._dataMailID = map[json.Number]SysMailConfig{}
	for _, d := range a._data {
		a._dataMailID[d.MailID] = d
	}
}

// ByMailID returns the SysMailConfig uniquely identified by MailID
//
// Error is only non-nil if the source errors out
func (a *SysMailConfigAccessor) ByMailID(identifier json.Number) (SysMailConfig, error) {
	if a._dataMailID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SysMailConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMailID[identifier], nil
}
