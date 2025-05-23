package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueMagicUnit struct {
	AttachRangeTypeList []string   `json:"AttachRangeTypeList"`
	EffectTypeList      []string   `json:"EffectTypeList"`
	ExtraEffectID       []float64  `json:"ExtraEffectID"`
	FuncType            string     `json:"FuncType"`
	LimitRange          string     `json:"LimitRange"`
	MagicUnitCategory   string     `json:"MagicUnitCategory"`
	MagicUnitDesc       hash.Hash  `json:"MagicUnitDesc"`
	MagicUnitID         float64    `json:"MagicUnitID"`
	MagicUnitLevel      float64    `json:"MagicUnitLevel"`
	MagicUnitMazeBuffID float64    `json:"MagicUnitMazeBuffID"`
	MagicUnitSimpleDesc hash.Hash  `json:"MagicUnitSimpleDesc"`
	MagicUnitType       string     `json:"MagicUnitType"`
	SpecialType         string     `json:"SpecialType"`
	StyleType           string     `json:"StyleType"`
	UnitBasicPower      hash.Value `json:"UnitBasicPower"`
	UnlockID            float64    `json:"UnlockID"`
}
type RogueMagicUnitAccessor struct {
	_data []RogueMagicUnit
}

// LoadData retrieves the data. Must be called before RogueMagicUnit.GroupData
func (a *RogueMagicUnitAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicUnit.json")
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
func (a *RogueMagicUnitAccessor) Raw() ([]RogueMagicUnit, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicUnit{}, err
		}
	}
	return a._data, nil
}
