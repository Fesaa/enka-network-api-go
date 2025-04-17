package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TrainPartyGridType struct {
	GridType        string    `json:"GridType"`
	GridTypeName    hash.Hash `json:"GridTypeName"`
	GridTypeTipInfo hash.Hash `json:"GridTypeTipInfo"`
}
type TrainPartyGridTypeAccessor struct {
	_data         []TrainPartyGridType
	_dataGridType map[string]TrainPartyGridType
}

// LoadData retrieves the data. Must be called before TrainPartyGridType.GroupData
func (a *TrainPartyGridTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyGridType.json")
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
func (a *TrainPartyGridTypeAccessor) Raw() ([]TrainPartyGridType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyGridType{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartyGridTypeAccessor.LoadData to preload everything
func (a *TrainPartyGridTypeAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGridType[d.GridType] = d
	}
}

// ByGridType returns the TrainPartyGridType uniquely identified by GridType
//
// Error is only non-nil if the source errors out
func (a *TrainPartyGridTypeAccessor) ByGridType(identifier string) (TrainPartyGridType, error) {
	if a._dataGridType == nil {
		err := a.LoadData()
		if err != nil {
			return TrainPartyGridType{}, err
		}
		a.GroupData()
	}
	return a._dataGridType[identifier], nil
}
