package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AvatarRelicRecommend struct {
	AvatarID             float64                            `json:"AvatarID"`
	LocalCriticalChance  hash.IntValue                      `json:"LocalCriticalChance"`
	PropertyList         []AvatarRelicRecommendPropertyList `json:"PropertyList"`
	PropertyList3        []string                           `json:"PropertyList3"`
	PropertyList4        []string                           `json:"PropertyList4"`
	PropertyList5        []string                           `json:"PropertyList5"`
	PropertyList6        []string                           `json:"PropertyList6"`
	ScoreRankList        []float64                          `json:"ScoreRankList"`
	Set2IDList           []float64                          `json:"Set2IDList"`
	Set4IDList           []float64                          `json:"Set4IDList"`
	SubAffixPropertyList []string                           `json:"SubAffixPropertyList"`
}
type AvatarRelicRecommendPropertyList struct {
	PropertyType string `json:"PropertyType"`
	RelicType    string `json:"RelicType"`
}
type AvatarRelicRecommendAccessor struct {
	_data         []AvatarRelicRecommend
	_dataAvatarID map[float64]AvatarRelicRecommend
}

// LoadData retrieves the data. Must be called before AvatarRelicRecommend.GroupData
func (a *AvatarRelicRecommendAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarRelicRecommend.json")
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
func (a *AvatarRelicRecommendAccessor) Raw() ([]AvatarRelicRecommend, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarRelicRecommend{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarRelicRecommendAccessor.LoadData to preload everything
func (a *AvatarRelicRecommendAccessor) GroupData() {
	for _, d := range a._data {
		a._dataAvatarID[d.AvatarID] = d
	}
}

// ByAvatarID returns the AvatarRelicRecommend uniquely identified by AvatarID
//
// Error is only non-nil if the source errors out
func (a *AvatarRelicRecommendAccessor) ByAvatarID(identifier float64) (AvatarRelicRecommend, error) {
	if a._dataAvatarID == nil {
		err := a.LoadData()
		if err != nil {
			return AvatarRelicRecommend{}, err
		}
		a.GroupData()
	}
	return a._dataAvatarID[identifier], nil
}
