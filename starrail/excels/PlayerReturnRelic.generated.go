package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PlayerReturnRelic struct {
	IsRelicMatchMainAffix bool    `json:"IsRelicMatchMainAffix"`
	RelicLevel            float64 `json:"RelicLevel"`
	RelicRarity           string  `json:"RelicRarity"`
	WorldLevel            float64 `json:"WorldLevel"`
}
type PlayerReturnRelicAccessor struct {
	_data []PlayerReturnRelic
}

// LoadData retrieves the data. Must be called before PlayerReturnRelic.GroupData
func (a *PlayerReturnRelicAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlayerReturnRelic.json")
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
func (a *PlayerReturnRelicAccessor) Raw() ([]PlayerReturnRelic, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlayerReturnRelic{}, err
		}
	}
	return a._data, nil
}
