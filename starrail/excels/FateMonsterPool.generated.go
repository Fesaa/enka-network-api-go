package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type FateMonsterPool struct {
	GDLDHCCGPNM []float64 `json:"GDLDHCCGPNM"`
	HJJGPMNHCDJ []float64 `json:"HJJGPMNHCDJ"`
	HOCLLKOJDGO []float64 `json:"HOCLLKOJDGO"`
	MNOBIEDPCHI []float64 `json:"MNOBIEDPCHI"`
	MOLGEJJJMEA []float64 `json:"MOLGEJJJMEA"`
	PMMHDBALPBN float64   `json:"PMMHDBALPBN"`
}
type FateMonsterPoolAccessor struct {
	_data            []FateMonsterPool
	_dataPMMHDBALPBN map[float64]FateMonsterPool
}

// LoadData retrieves the data. Must be called before FateMonsterPool.GroupData
func (a *FateMonsterPoolAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FateMonsterPool.json")
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
func (a *FateMonsterPoolAccessor) Raw() ([]FateMonsterPool, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FateMonsterPool{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FateMonsterPoolAccessor.LoadData to preload everything
func (a *FateMonsterPoolAccessor) GroupData() {
	a._dataPMMHDBALPBN = map[float64]FateMonsterPool{}
	for _, d := range a._data {
		a._dataPMMHDBALPBN[d.PMMHDBALPBN] = d
	}
}

// ByPMMHDBALPBN returns the FateMonsterPool uniquely identified by PMMHDBALPBN
//
// Error is only non-nil if the source errors out
func (a *FateMonsterPoolAccessor) ByPMMHDBALPBN(identifier float64) (FateMonsterPool, error) {
	if a._dataPMMHDBALPBN == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FateMonsterPool{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPMMHDBALPBN[identifier], nil
}
