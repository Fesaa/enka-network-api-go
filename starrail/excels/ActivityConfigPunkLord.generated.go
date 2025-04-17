package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ActivityConfigPunkLord struct {
	AssistPoint   json.Number            `json:"AssistPoint"`
	ExistTime     json.Number            `json:"ExistTime"`
	GroupType     string                 `json:"GroupType"`
	ID            json.Number            `json:"ID"`
	KillPoint     json.Number            `json:"KillPoint"`
	ManikinConfig string                 `json:"ManikinConfig"`
	MonsterBuff   map[string]json.Number `json:"MonsterBuff"`
	MonsterHP     json.Number            `json:"MonsterHP"`
	MonsterLevel  json.Number            `json:"MonsterLevel"`
	MonsterPic    string                 `json:"MonsterPic"`
	MonsterRare   string                 `json:"MonsterRare"`
	PluralHP      json.Number            `json:"PluralHP"`
	RaidID        json.Number            `json:"RaidID"`
	ShowMonster   string                 `json:"ShowMonster"`
	SummonPoint   json.Number            `json:"SummonPoint"`
	TurnLimit     json.Number            `json:"TurnLimit"`
	WorldLevel    json.Number            `json:"WorldLevel"`
}
type ActivityConfigPunkLordAccessor struct {
	_data []ActivityConfigPunkLord
}

// LoadData retrieves the data. Must be called before ActivityConfigPunkLord.GroupData
func (a *ActivityConfigPunkLordAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityConfigPunkLord.json")
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
func (a *ActivityConfigPunkLordAccessor) Raw() ([]ActivityConfigPunkLord, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityConfigPunkLord{}, err
		}
	}
	return a._data, nil
}
