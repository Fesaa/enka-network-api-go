package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type HeartDialScript struct {
	ControlDialogueID float64                           `json:"ControlDialogueID"`
	DefaultEmoType    string                            `json:"DefaultEmoType"`
	FullDialogueID    float64                           `json:"FullDialogueID"`
	LockDialogueID    float64                           `json:"LockDialogueID"`
	MissingDialogueID float64                           `json:"MissingDialogueID"`
	MissingEmoList    []string                          `json:"MissingEmoList"`
	RaidID            float64                           `json:"RaidID"`
	ScriptID          float64                           `json:"ScriptID"`
	StepList          []string                          `json:"StepList"`
	TotalEmoInfoList  []HeartDialScriptTotalEmoInfoList `json:"TotalEmoInfoList"`
	UnLockDialogueID  float64                           `json:"UnLockDialogueID"`
}
type HeartDialScriptTotalEmoInfoList struct {
	BNNBAAPNHCF string  `json:"BNNBAAPNHCF"`
	DKDHKFKAAHN float64 `json:"DKDHKFKAAHN"`
}
type HeartDialScriptAccessor struct {
	_data         []HeartDialScript
	_dataScriptID map[float64]HeartDialScript
}

// LoadData retrieves the data. Must be called before HeartDialScript.GroupData
func (a *HeartDialScriptAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HeartDialScript.json")
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
func (a *HeartDialScriptAccessor) Raw() ([]HeartDialScript, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HeartDialScript{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HeartDialScriptAccessor.LoadData to preload everything
func (a *HeartDialScriptAccessor) GroupData() {
	for _, d := range a._data {
		a._dataScriptID[d.ScriptID] = d
	}
}

// ByScriptID returns the HeartDialScript uniquely identified by ScriptID
//
// Error is only non-nil if the source errors out
func (a *HeartDialScriptAccessor) ByScriptID(identifier float64) (HeartDialScript, error) {
	if a._dataScriptID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeartDialScript{}, err
			}
		}
		a.GroupData()
	}
	return a._dataScriptID[identifier], nil
}
