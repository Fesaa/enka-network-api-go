package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueCommonModeTitle struct {
	SubMode        string    `json:"SubMode"`
	TitleIconPath  string    `json:"TitleIconPath"`
	TitleTextmapID hash.Hash `json:"TitleTextmapID"`
}
type RogueCommonModeTitleAccessor struct {
	_data              []RogueCommonModeTitle
	_dataSubMode       map[string]RogueCommonModeTitle
	_dataTitleIconPath map[string]RogueCommonModeTitle
}

// LoadData retrieves the data. Must be called before RogueCommonModeTitle.GroupData
func (a *RogueCommonModeTitleAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueCommonModeTitle.json")
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
func (a *RogueCommonModeTitleAccessor) Raw() ([]RogueCommonModeTitle, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueCommonModeTitle{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueCommonModeTitleAccessor.LoadData to preload everything
func (a *RogueCommonModeTitleAccessor) GroupData() {
	a._dataSubMode = map[string]RogueCommonModeTitle{}
	a._dataTitleIconPath = map[string]RogueCommonModeTitle{}
	for _, d := range a._data {
		a._dataSubMode[d.SubMode] = d
		a._dataTitleIconPath[d.TitleIconPath] = d
	}
}

// BySubMode returns the RogueCommonModeTitle uniquely identified by SubMode
//
// Error is only non-nil if the source errors out
func (a *RogueCommonModeTitleAccessor) BySubMode(identifier string) (RogueCommonModeTitle, error) {
	if a._dataSubMode == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueCommonModeTitle{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSubMode[identifier], nil
}

// ByTitleIconPath returns the RogueCommonModeTitle uniquely identified by TitleIconPath
//
// Error is only non-nil if the source errors out
func (a *RogueCommonModeTitleAccessor) ByTitleIconPath(identifier string) (RogueCommonModeTitle, error) {
	if a._dataTitleIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueCommonModeTitle{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTitleIconPath[identifier], nil
}
