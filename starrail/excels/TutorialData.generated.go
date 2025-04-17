package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TutorialData struct {
	CanInterrupt        bool                              `json:"CanInterrupt"`
	FinishTriggerParams []TutorialDataFinishTriggerParams `json:"FinishTriggerParams"`
	Priority            float64                           `json:"Priority"`
	RestoreType         string                            `json:"RestoreType"`
	TriggerParams       []TutorialDataTriggerParams       `json:"TriggerParams"`
	TutorialID          float64                           `json:"TutorialID"`
	TutorialJsonPath    string                            `json:"TutorialJsonPath"`
}
type TutorialDataFinishTriggerParams struct {
	TriggerParam string `json:"TriggerParam"`
	TriggerType  string `json:"TriggerType"`
}
type TutorialDataTriggerParams struct {
	TriggerParam string `json:"TriggerParam"`
	TriggerType  string `json:"TriggerType"`
}
type TutorialDataAccessor struct {
	_data           []TutorialData
	_dataTutorialID map[float64]TutorialData
}

// LoadData retrieves the data. Must be called before TutorialData.GroupData
func (a *TutorialDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TutorialData.json")
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
func (a *TutorialDataAccessor) Raw() ([]TutorialData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TutorialData{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TutorialDataAccessor.LoadData to preload everything
func (a *TutorialDataAccessor) GroupData() {
	for _, d := range a._data {
		a._dataTutorialID[d.TutorialID] = d
	}
}

// ByTutorialID returns the TutorialData uniquely identified by TutorialID
//
// Error is only non-nil if the source errors out
func (a *TutorialDataAccessor) ByTutorialID(identifier float64) (TutorialData, error) {
	if a._dataTutorialID == nil {
		err := a.LoadData()
		if err != nil {
			return TutorialData{}, err
		}
		a.GroupData()
	}
	return a._dataTutorialID[identifier], nil
}
