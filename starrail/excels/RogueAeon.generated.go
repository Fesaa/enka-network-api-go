package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueAeon struct {
	AeonID                      json.Number            `json:"AeonID"`
	ArrivedTalkDialogueGroupID  json.Number            `json:"ArrivedTalkDialogueGroupID"`
	BattleEventBuffGroup        json.Number            `json:"BattleEventBuffGroup"`
	BattleEventEnhanceBuffGroup json.Number            `json:"BattleEventEnhanceBuffGroup"`
	DisplayID                   json.Number            `json:"DisplayID"`
	EffectDesc1                 map[string]json.Number `json:"EffectDesc1"`
	EffectDesc2                 map[string]json.Number `json:"EffectDesc2"`
	RogueBuffType               json.Number            `json:"RogueBuffType"`
	RogueVersion                json.Number            `json:"RogueVersion"`
	Sort                        json.Number            `json:"Sort"`
	UnlockID                    json.Number            `json:"UnlockID"`
}
type RogueAeonAccessor struct {
	_data []RogueAeon
}

// LoadData retrieves the data. Must be called before RogueAeon.GroupData
func (a *RogueAeonAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueAeon.json")
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
func (a *RogueAeonAccessor) Raw() ([]RogueAeon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueAeon{}, err
		}
	}
	return a._data, nil
}
