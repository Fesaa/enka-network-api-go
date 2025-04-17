package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TravelBrochureConfig struct {
	BackgroundPrefab     string                           `json:"BackgroundPrefab"`
	Conditions           []TravelBrochureConfigConditions `json:"Conditions"`
	DiaryGroupID         float64                          `json:"DiaryGroupID"`
	DirectoryName        hash.Hash                        `json:"DirectoryName"`
	FinishQuestID        float64                          `json:"FinishQuestID"`
	FrontPrefab          string                           `json:"FrontPrefab"`
	ID                   float64                          `json:"ID"`
	PasterAchievementPic string                           `json:"PasterAchievementPic"`
	PicPath              string                           `json:"PicPath"`
	ShowInDirectory      bool                             `json:"ShowInDirectory"`
	ShowUnlockToast      bool                             `json:"ShowUnlockToast"`
	Sort                 float64                          `json:"Sort"`
	Type                 string                           `json:"Type"`
}
type TravelBrochureConfigConditions struct {
	Param string `json:"Param"`
	Type  string `json:"Type"`
}
type TravelBrochureConfigAccessor struct {
	_data             []TravelBrochureConfig
	_dataSort         map[float64]TravelBrochureConfig
	_dataID           map[float64]TravelBrochureConfig
	_dataDiaryGroupID map[float64]TravelBrochureConfig
}

// LoadData retrieves the data. Must be called before TravelBrochureConfig.GroupData
func (a *TravelBrochureConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TravelBrochureConfig.json")
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
func (a *TravelBrochureConfigAccessor) Raw() ([]TravelBrochureConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TravelBrochureConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TravelBrochureConfigAccessor.LoadData to preload everything
func (a *TravelBrochureConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataSort[d.Sort] = d
		a._dataID[d.ID] = d
		a._dataDiaryGroupID[d.DiaryGroupID] = d
	}
}

// BySort returns the TravelBrochureConfig uniquely identified by Sort
//
// Error is only non-nil if the source errors out
func (a *TravelBrochureConfigAccessor) BySort(identifier float64) (TravelBrochureConfig, error) {
	if a._dataSort == nil {
		err := a.LoadData()
		if err != nil {
			return TravelBrochureConfig{}, err
		}
		a.GroupData()
	}
	return a._dataSort[identifier], nil
}

// ByID returns the TravelBrochureConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TravelBrochureConfigAccessor) ByID(identifier float64) (TravelBrochureConfig, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return TravelBrochureConfig{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByDiaryGroupID returns the TravelBrochureConfig uniquely identified by DiaryGroupID
//
// Error is only non-nil if the source errors out
func (a *TravelBrochureConfigAccessor) ByDiaryGroupID(identifier float64) (TravelBrochureConfig, error) {
	if a._dataDiaryGroupID == nil {
		err := a.LoadData()
		if err != nil {
			return TravelBrochureConfig{}, err
		}
		a.GroupData()
	}
	return a._dataDiaryGroupID[identifier], nil
}
