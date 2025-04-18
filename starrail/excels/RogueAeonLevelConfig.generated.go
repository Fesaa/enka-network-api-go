package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueAeonLevelConfig struct {
	AeonStory     hash.Hash `json:"AeonStory"`
	AeonStoryID   float64   `json:"AeonStoryID"`
	AeonStoryName hash.Hash `json:"AeonStory_Name"`
	Exp           float64   `json:"Exp"`
	Level         float64   `json:"Level"`
	RogueAeonID   float64   `json:"RogueAeonID"`
	UnlockID      float64   `json:"UnlockID"`
}
type RogueAeonLevelConfigAccessor struct {
	_data []RogueAeonLevelConfig
}

// LoadData retrieves the data. Must be called before RogueAeonLevelConfig.GroupData
func (a *RogueAeonLevelConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueAeonLevelConfig.json")
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
func (a *RogueAeonLevelConfigAccessor) Raw() ([]RogueAeonLevelConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueAeonLevelConfig{}, err
		}
	}
	return a._data, nil
}
