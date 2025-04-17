package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueMagicUnlock struct {
	RogueUnlockDetail map[string]json.Number `json:"RogueUnlockDetail"`
	RogueUnlockID     json.Number            `json:"RogueUnlockID"`
	UnlockFinishWay   json.Number            `json:"UnlockFinishWay"`
}
type RogueMagicUnlockAccessor struct {
	_data []RogueMagicUnlock
}

// LoadData retrieves the data. Must be called before RogueMagicUnlock.GroupData
func (a *RogueMagicUnlockAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicUnlock.json")
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
func (a *RogueMagicUnlockAccessor) Raw() ([]RogueMagicUnlock, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicUnlock{}, err
		}
	}
	return a._data, nil
}
