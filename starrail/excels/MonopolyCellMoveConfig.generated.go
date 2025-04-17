package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MonopolyCellMoveConfig struct {
	CellID    float64  `json:"CellID"`
	MapID     float64  `json:"MapID"`
	MoveParam []string `json:"MoveParam"`
}
type MonopolyCellMoveConfigAccessor struct {
	_data []MonopolyCellMoveConfig
}

// LoadData retrieves the data. Must be called before MonopolyCellMoveConfig.GroupData
func (a *MonopolyCellMoveConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonopolyCellMoveConfig.json")
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
func (a *MonopolyCellMoveConfigAccessor) Raw() ([]MonopolyCellMoveConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonopolyCellMoveConfig{}, err
		}
	}
	return a._data, nil
}
