package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MonsterDifficultyGuide struct {
	DifficultyGuideDescription map[string]json.Number `json:"DifficultyGuideDescription"`
	DifficultyGuideID          json.Number            `json:"DifficultyGuideID"`
	ParameterList              []json.Number          `json:"ParameterList"`
	SkillID                    json.Number            `json:"SkillID"`
}
type MonsterDifficultyGuideAccessor struct {
	_data []MonsterDifficultyGuide
}

// LoadData retrieves the data. Must be called before MonsterDifficultyGuide.GroupData
func (a *MonsterDifficultyGuideAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonsterDifficultyGuide.json")
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
func (a *MonsterDifficultyGuideAccessor) Raw() ([]MonsterDifficultyGuide, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonsterDifficultyGuide{}, err
		}
	}
	return a._data, nil
}
