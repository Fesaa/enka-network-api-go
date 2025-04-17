package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type AvatarVO struct {
	ActionBegin           float64 `json:"ActionBegin"`
	ActionBeginAdvantage  float64 `json:"ActionBeginAdvantage"`
	ActionBeginHighThreat float64 `json:"ActionBeginHighThreat"`
	LightHit              float64 `json:"LightHit"`
	ReceiveHealing        float64 `json:"ReceiveHealing"`
	Revived               float64 `json:"Revived"`
	StandBy               float64 `json:"StandBy"`
	UltraReady            float64 `json:"UltraReady"`
	VOTag                 string  `json:"VOTag"`
}
type AvatarVOAccessor struct {
	_data      []AvatarVO
	_dataVOTag map[string]AvatarVO
}

// LoadData retrieves the data. Must be called before AvatarVO.GroupData
func (a *AvatarVOAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarVO.json")
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
func (a *AvatarVOAccessor) Raw() ([]AvatarVO, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarVO{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarVOAccessor.LoadData to preload everything
func (a *AvatarVOAccessor) GroupData() {
	a._dataVOTag = map[string]AvatarVO{}
	for _, d := range a._data {
		a._dataVOTag[d.VOTag] = d
	}
}

// ByVOTag returns the AvatarVO uniquely identified by VOTag
//
// Error is only non-nil if the source errors out
func (a *AvatarVOAccessor) ByVOTag(identifier string) (AvatarVO, error) {
	if a._dataVOTag == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarVO{}, err
			}
		}
		a.GroupData()
	}
	return a._dataVOTag[identifier], nil
}
