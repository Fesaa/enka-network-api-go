package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueDLCBlockIntro struct {
	BlockIntroDesc           map[string]json.Number `json:"BlockIntroDesc"`
	BlockIntroID             json.Number            `json:"BlockIntroID"`
	BlockIntroIcon           string                 `json:"BlockIntroIcon"`
	BlockIntroName           map[string]json.Number `json:"BlockIntroName"`
	BlockTypeChessBoardColor string                 `json:"BlockTypeChessBoardColor"`
	IntroGroup               json.Number            `json:"IntroGroup"`
	Sort                     json.Number            `json:"Sort"`
	SubType                  []string               `json:"SubType"`
}
type RogueDLCBlockIntroAccessor struct {
	_data               []RogueDLCBlockIntro
	_dataBlockIntroIcon map[string]RogueDLCBlockIntro
}

// LoadData retrieves the data. Must be called before RogueDLCBlockIntro.GroupData
func (a *RogueDLCBlockIntroAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCBlockIntro.json")
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
func (a *RogueDLCBlockIntroAccessor) Raw() ([]RogueDLCBlockIntro, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCBlockIntro{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCBlockIntroAccessor.LoadData to preload everything
func (a *RogueDLCBlockIntroAccessor) GroupData() {
	a._dataBlockIntroIcon = map[string]RogueDLCBlockIntro{}
	for _, d := range a._data {
		a._dataBlockIntroIcon[d.BlockIntroIcon] = d
	}
}

// ByBlockIntroIcon returns the RogueDLCBlockIntro uniquely identified by BlockIntroIcon
//
// Error is only non-nil if the source errors out
func (a *RogueDLCBlockIntroAccessor) ByBlockIntroIcon(identifier string) (RogueDLCBlockIntro, error) {
	if a._dataBlockIntroIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCBlockIntro{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBlockIntroIcon[identifier], nil
}
