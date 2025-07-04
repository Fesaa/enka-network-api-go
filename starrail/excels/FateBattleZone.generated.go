package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type FateBattleZone struct {
	ALDLKHGHFBF float64 `json:"ALDLKHGHFBF"`
	DFMKHGBHABJ float64 `json:"DFMKHGBHABJ"`
	OAIPACNPKCA float64 `json:"OAIPACNPKCA"`
	PAPFKIFLCKC string  `json:"PAPFKIFLCKC"`
	PCNCAJNPDJF string  `json:"PCNCAJNPDJF"`
	PPPFDAMIPAC float64 `json:"PPPFDAMIPAC"`
}
type FateBattleZoneAccessor struct {
	_data            []FateBattleZone
	_dataPPPFDAMIPAC map[float64]FateBattleZone
}

// LoadData retrieves the data. Must be called before FateBattleZone.GroupData
func (a *FateBattleZoneAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FateBattleZone.json")
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
func (a *FateBattleZoneAccessor) Raw() ([]FateBattleZone, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FateBattleZone{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FateBattleZoneAccessor.LoadData to preload everything
func (a *FateBattleZoneAccessor) GroupData() {
	a._dataPPPFDAMIPAC = map[float64]FateBattleZone{}
	for _, d := range a._data {
		a._dataPPPFDAMIPAC[d.PPPFDAMIPAC] = d
	}
}

// ByPPPFDAMIPAC returns the FateBattleZone uniquely identified by PPPFDAMIPAC
//
// Error is only non-nil if the source errors out
func (a *FateBattleZoneAccessor) ByPPPFDAMIPAC(identifier float64) (FateBattleZone, error) {
	if a._dataPPPFDAMIPAC == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FateBattleZone{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPPPFDAMIPAC[identifier], nil
}
