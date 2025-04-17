package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type AlleyDeskTalk struct {
	CustomString  string      `json:"CustomString"`
	TalkID        json.Number `json:"TalkID"`
	TalkPriority  json.Number `json:"TalkPriority"`
	TalkType      string      `json:"TalkType"`
	TalkTypeParam string      `json:"TalkTypeParam"`
	TalkWeight    json.Number `json:"TalkWeight"`
	TextIDList    string      `json:"TextIDList"`
}
type AlleyDeskTalkAccessor struct {
	_data           []AlleyDeskTalk
	_dataTalkID     map[json.Number]AlleyDeskTalk
	_dataTextIDList map[string]AlleyDeskTalk
}

// LoadData retrieves the data. Must be called before AlleyDeskTalk.GroupData
func (a *AlleyDeskTalkAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AlleyDeskTalk.json")
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
func (a *AlleyDeskTalkAccessor) Raw() ([]AlleyDeskTalk, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AlleyDeskTalk{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AlleyDeskTalkAccessor.LoadData to preload everything
func (a *AlleyDeskTalkAccessor) GroupData() {
	a._dataTalkID = map[json.Number]AlleyDeskTalk{}
	a._dataTextIDList = map[string]AlleyDeskTalk{}
	for _, d := range a._data {
		a._dataTalkID[d.TalkID] = d
		a._dataTextIDList[d.TextIDList] = d
	}
}

// ByTalkID returns the AlleyDeskTalk uniquely identified by TalkID
//
// Error is only non-nil if the source errors out
func (a *AlleyDeskTalkAccessor) ByTalkID(identifier json.Number) (AlleyDeskTalk, error) {
	if a._dataTalkID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AlleyDeskTalk{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTalkID[identifier], nil
}

// ByTextIDList returns the AlleyDeskTalk uniquely identified by TextIDList
//
// Error is only non-nil if the source errors out
func (a *AlleyDeskTalkAccessor) ByTextIDList(identifier string) (AlleyDeskTalk, error) {
	if a._dataTextIDList == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AlleyDeskTalk{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTextIDList[identifier], nil
}
