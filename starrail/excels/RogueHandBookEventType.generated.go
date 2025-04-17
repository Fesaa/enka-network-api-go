package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueHandBookEventType struct {
	ActivityModuleID       float64   `json:"ActivityModuleID"`
	RogueEventTypeTitle    hash.Hash `json:"RogueEventTypeTitle"`
	RogueHandBookEventType float64   `json:"RogueHandBookEventType"`
	TypeIcon               string    `json:"TypeIcon"`
}
type RogueHandBookEventTypeAccessor struct {
	_data                       []RogueHandBookEventType
	_dataRogueHandBookEventType map[float64]RogueHandBookEventType
	_dataTypeIcon               map[string]RogueHandBookEventType
}

// LoadData retrieves the data. Must be called before RogueHandBookEventType.GroupData
func (a *RogueHandBookEventTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueHandBookEventType.json")
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
func (a *RogueHandBookEventTypeAccessor) Raw() ([]RogueHandBookEventType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueHandBookEventType{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueHandBookEventTypeAccessor.LoadData to preload everything
func (a *RogueHandBookEventTypeAccessor) GroupData() {
	a._dataRogueHandBookEventType = map[float64]RogueHandBookEventType{}
	a._dataTypeIcon = map[string]RogueHandBookEventType{}
	for _, d := range a._data {
		a._dataRogueHandBookEventType[d.RogueHandBookEventType] = d
		a._dataTypeIcon[d.TypeIcon] = d
	}
}

// ByRogueHandBookEventType returns the RogueHandBookEventType uniquely identified by RogueHandBookEventType
//
// Error is only non-nil if the source errors out
func (a *RogueHandBookEventTypeAccessor) ByRogueHandBookEventType(identifier float64) (RogueHandBookEventType, error) {
	if a._dataRogueHandBookEventType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueHandBookEventType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueHandBookEventType[identifier], nil
}

// ByTypeIcon returns the RogueHandBookEventType uniquely identified by TypeIcon
//
// Error is only non-nil if the source errors out
func (a *RogueHandBookEventTypeAccessor) ByTypeIcon(identifier string) (RogueHandBookEventType, error) {
	if a._dataTypeIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueHandBookEventType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTypeIcon[identifier], nil
}
