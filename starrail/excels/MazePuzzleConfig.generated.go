package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MazePuzzleConfig struct {
	DefaultCDDuration float64   `json:"DefaultCDDuration"`
	IconPath          string    `json:"IconPath"`
	PuzzleFuncType    string    `json:"PuzzleFuncType"`
	ShowFuncBtnHint   hash.Hash `json:"ShowFuncBtnHint"`
}
type MazePuzzleConfigAccessor struct {
	_data []MazePuzzleConfig
}

// LoadData retrieves the data. Must be called before MazePuzzleConfig.GroupData
func (a *MazePuzzleConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MazePuzzleConfig.json")
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
func (a *MazePuzzleConfigAccessor) Raw() ([]MazePuzzleConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MazePuzzleConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}
