package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCBlockIntro struct {
	BlockIntroDesc           hash.Hash `json:"BlockIntroDesc"`
	BlockIntroID             float64   `json:"BlockIntroID"`
	BlockIntroIcon           string    `json:"BlockIntroIcon"`
	BlockIntroName           hash.Hash `json:"BlockIntroName"`
	BlockTypeChessBoardColor string    `json:"BlockTypeChessBoardColor"`
	IntroGroup               float64   `json:"IntroGroup"`
	Sort                     float64   `json:"Sort"`
	SubType                  []string  `json:"SubType"`
}
type RogueDLCBlockIntroAccessor struct {
	_data               []RogueDLCBlockIntro
	_dataBlockIntroID   map[float64]RogueDLCBlockIntro
	_dataBlockIntroIcon map[string]RogueDLCBlockIntro
	_dataSort           map[float64]RogueDLCBlockIntro
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
	for _, d := range a._data {
		a._dataBlockIntroID[d.BlockIntroID] = d
		a._dataBlockIntroIcon[d.BlockIntroIcon] = d
		a._dataSort[d.Sort] = d
	}
}

// ByBlockIntroID returns the RogueDLCBlockIntro uniquely identified by BlockIntroID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCBlockIntroAccessor) ByBlockIntroID(identifier float64) (RogueDLCBlockIntro, error) {
	if a._dataBlockIntroID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCBlockIntro{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBlockIntroID[identifier], nil
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

// BySort returns the RogueDLCBlockIntro uniquely identified by Sort
//
// Error is only non-nil if the source errors out
func (a *RogueDLCBlockIntroAccessor) BySort(identifier float64) (RogueDLCBlockIntro, error) {
	if a._dataSort == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCBlockIntro{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSort[identifier], nil
}
