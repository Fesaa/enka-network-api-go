package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueNousDiceBranchValue struct {
	AeonID           float64      `json:"AeonID"`
	BranchEffectDesc hash.Hash    `json:"BranchEffectDesc"`
	BranchID         float64      `json:"BranchID"`
	ParamList        []hash.Value `json:"ParamList"`
}
type RogueNousDiceBranchValueAccessor struct {
	_data []RogueNousDiceBranchValue
}

// LoadData retrieves the data. Must be called before RogueNousDiceBranchValue.GroupData
func (a *RogueNousDiceBranchValueAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueNousDiceBranchValue.json")
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
func (a *RogueNousDiceBranchValueAccessor) Raw() ([]RogueNousDiceBranchValue, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueNousDiceBranchValue{}, err
		}
	}
	return a._data, nil
}
