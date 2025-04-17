package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueNPC struct {
	NPCJsonPath string      `json:"NPCJsonPath"`
	RogueNPCID  json.Number `json:"RogueNPCID"`
}
type RogueNPCAccessor struct {
	_data []RogueNPC
}

// LoadData retrieves the data. Must be called before RogueNPC.GroupData
func (a *RogueNPCAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueNPC.json")
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
func (a *RogueNPCAccessor) Raw() ([]RogueNPC, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueNPC{}, err
		}
	}
	return a._data, nil
}
