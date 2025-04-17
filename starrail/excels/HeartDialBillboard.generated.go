package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type HeartDialBillboard struct {
	EmoType   string  `json:"EmoType"`
	MapIconID float64 `json:"MapIconID"`
	StepType  string  `json:"StepType"`
}
type HeartDialBillboardAccessor struct {
	_data []HeartDialBillboard
}

// LoadData retrieves the data. Must be called before HeartDialBillboard.GroupData
func (a *HeartDialBillboardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HeartDialBillboard.json")
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
func (a *HeartDialBillboardAccessor) Raw() ([]HeartDialBillboard, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HeartDialBillboard{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}
