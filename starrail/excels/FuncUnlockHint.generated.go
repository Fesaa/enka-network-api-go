package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type FuncUnlockHint struct {
	Desc     hash.Hash `json:"Desc"`
	IconPath string    `json:"IconPath"`
	SubTitle hash.Hash `json:"SubTitle"`
	Title    hash.Hash `json:"Title"`
	Type     string    `json:"Type"`
	UnlockID float64   `json:"UnlockID"`
}
type FuncUnlockHintAccessor struct {
	_data         []FuncUnlockHint
	_dataUnlockID map[float64]FuncUnlockHint
}

// LoadData retrieves the data. Must be called before FuncUnlockHint.GroupData
func (a *FuncUnlockHintAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FuncUnlockHint.json")
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
func (a *FuncUnlockHintAccessor) Raw() ([]FuncUnlockHint, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FuncUnlockHint{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FuncUnlockHintAccessor.LoadData to preload everything
func (a *FuncUnlockHintAccessor) GroupData() {
	for _, d := range a._data {
		a._dataUnlockID[d.UnlockID] = d
	}
}

// ByUnlockID returns the FuncUnlockHint uniquely identified by UnlockID
//
// Error is only non-nil if the source errors out
func (a *FuncUnlockHintAccessor) ByUnlockID(identifier float64) (FuncUnlockHint, error) {
	if a._dataUnlockID == nil {
		err := a.LoadData()
		if err != nil {
			return FuncUnlockHint{}, err
		}
		a.GroupData()
	}
	return a._dataUnlockID[identifier], nil
}
