package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ChallengeBossGroupExtra struct {
	BossPatternPrefabPath       string    `json:"BossPatternPrefabPath"`
	BossPositionPrefabPath1     string    `json:"BossPositionPrefabPath1"`
	BossPositionPrefabPath2     string    `json:"BossPositionPrefabPath2"`
	BuffList1                   []float64 `json:"BuffList1"`
	BuffList2                   []float64 `json:"BuffList2"`
	GroupID                     float64   `json:"GroupID"`
	ThemeIconPicPath            string    `json:"ThemeIconPicPath"`
	ThemePosterBgPicPath        string    `json:"ThemePosterBgPicPath"`
	ThemePosterEffectPrefabPath string    `json:"ThemePosterEffectPrefabPath"`
	ThemePosterTabPicPath       string    `json:"ThemePosterTabPicPath"`
	ThemeToastPicPath           string    `json:"ThemeToastPicPath"`
}
type ChallengeBossGroupExtraAccessor struct {
	_data        []ChallengeBossGroupExtra
	_dataGroupID map[float64]ChallengeBossGroupExtra
}

// LoadData retrieves the data. Must be called before ChallengeBossGroupExtra.GroupData
func (a *ChallengeBossGroupExtraAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChallengeBossGroupExtra.json")
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
func (a *ChallengeBossGroupExtraAccessor) Raw() ([]ChallengeBossGroupExtra, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChallengeBossGroupExtra{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChallengeBossGroupExtraAccessor.LoadData to preload everything
func (a *ChallengeBossGroupExtraAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGroupID[d.GroupID] = d
	}
}

// ByGroupID returns the ChallengeBossGroupExtra uniquely identified by GroupID
//
// Error is only non-nil if the source errors out
func (a *ChallengeBossGroupExtraAccessor) ByGroupID(identifier float64) (ChallengeBossGroupExtra, error) {
	if a._dataGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChallengeBossGroupExtra{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGroupID[identifier], nil
}
