package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueDialogueOptionDisplay struct {
	OptionDesc      hash.Hash `json:"OptionDesc"`
	OptionDisplayID float64   `json:"OptionDisplayID"`
	OptionTitle     hash.Hash `json:"OptionTitle"`
}
type RogueDialogueOptionDisplayAccessor struct {
	_data                []RogueDialogueOptionDisplay
	_dataOptionDisplayID map[float64]RogueDialogueOptionDisplay
}

// LoadData retrieves the data. Must be called before RogueDialogueOptionDisplay.GroupData
func (a *RogueDialogueOptionDisplayAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDialogueOptionDisplay.json")
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
func (a *RogueDialogueOptionDisplayAccessor) Raw() ([]RogueDialogueOptionDisplay, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDialogueOptionDisplay{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDialogueOptionDisplayAccessor.LoadData to preload everything
func (a *RogueDialogueOptionDisplayAccessor) GroupData() {
	for _, d := range a._data {
		a._dataOptionDisplayID[d.OptionDisplayID] = d
	}
}

// ByOptionDisplayID returns the RogueDialogueOptionDisplay uniquely identified by OptionDisplayID
//
// Error is only non-nil if the source errors out
func (a *RogueDialogueOptionDisplayAccessor) ByOptionDisplayID(identifier float64) (RogueDialogueOptionDisplay, error) {
	if a._dataOptionDisplayID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDialogueOptionDisplay{}, err
			}
		}
		a.GroupData()
	}
	return a._dataOptionDisplayID[identifier], nil
}
