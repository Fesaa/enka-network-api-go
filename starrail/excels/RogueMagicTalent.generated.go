package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueMagicTalent struct {
	Cost          []RogueMagicTalentCost       `json:"Cost"`
	DescParams    []RogueMagicTalentDescParams `json:"DescParams"`
	EffectDesc    hash.Hash                    `json:"EffectDesc"`
	Level         float64                      `json:"Level"`
	NameDisplayID float64                      `json:"NameDisplayID"`
	TalentID      float64                      `json:"TalentID"`
	TalentIcon    string                       `json:"TalentIcon"`
}
type RogueMagicTalentCost struct {
	ItemID  float64 `json:"ItemID"`
	ItemNum float64 `json:"ItemNum"`
}
type RogueMagicTalentDescParams struct {
	HMCDHMFHABF string `json:"HMCDHMFHABF"`
	PICHIHHCOCB string `json:"PICHIHHCOCB"`
}
type RogueMagicTalentAccessor struct {
	_data         []RogueMagicTalent
	_dataLevel    map[float64]RogueMagicTalent
	_dataTalentID map[float64]RogueMagicTalent
}

// LoadData retrieves the data. Must be called before RogueMagicTalent.GroupData
func (a *RogueMagicTalentAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicTalent.json")
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
func (a *RogueMagicTalentAccessor) Raw() ([]RogueMagicTalent, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicTalent{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueMagicTalentAccessor.LoadData to preload everything
func (a *RogueMagicTalentAccessor) GroupData() {
	for _, d := range a._data {
		a._dataLevel[d.Level] = d
		a._dataTalentID[d.TalentID] = d
	}
}

// ByLevel returns the RogueMagicTalent uniquely identified by Level
//
// Error is only non-nil if the source errors out
func (a *RogueMagicTalentAccessor) ByLevel(identifier float64) (RogueMagicTalent, error) {
	if a._dataLevel == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueMagicTalent{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLevel[identifier], nil
}

// ByTalentID returns the RogueMagicTalent uniquely identified by TalentID
//
// Error is only non-nil if the source errors out
func (a *RogueMagicTalentAccessor) ByTalentID(identifier float64) (RogueMagicTalent, error) {
	if a._dataTalentID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueMagicTalent{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTalentID[identifier], nil
}
