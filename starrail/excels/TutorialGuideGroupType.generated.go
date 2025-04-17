package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TutorialGuideGroupType struct {
	MessageIconPath string    `json:"MessageIconPath"`
	MessageTitle    hash.Hash `json:"MessageTitle"`
	TutorialType    float64   `json:"TutorialType"`
}
type TutorialGuideGroupTypeAccessor struct {
	_data                []TutorialGuideGroupType
	_dataMessageIconPath map[string]TutorialGuideGroupType
}

// LoadData retrieves the data. Must be called before TutorialGuideGroupType.GroupData
func (a *TutorialGuideGroupTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TutorialGuideGroupType.json")
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
func (a *TutorialGuideGroupTypeAccessor) Raw() ([]TutorialGuideGroupType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TutorialGuideGroupType{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TutorialGuideGroupTypeAccessor.LoadData to preload everything
func (a *TutorialGuideGroupTypeAccessor) GroupData() {
	a._dataMessageIconPath = map[string]TutorialGuideGroupType{}
	for _, d := range a._data {
		a._dataMessageIconPath[d.MessageIconPath] = d
	}
}

// ByMessageIconPath returns the TutorialGuideGroupType uniquely identified by MessageIconPath
//
// Error is only non-nil if the source errors out
func (a *TutorialGuideGroupTypeAccessor) ByMessageIconPath(identifier string) (TutorialGuideGroupType, error) {
	if a._dataMessageIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TutorialGuideGroupType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMessageIconPath[identifier], nil
}
