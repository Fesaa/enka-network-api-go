package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type HeliobusSpecialPost struct {
	HeliobusSpecialPostID float64                    `json:"HeliobusSpecialPostID"`
	Likes                 []HeliobusSpecialPostLikes `json:"Likes"`
	PostImgIDList         []float64                  `json:"PostImgIDList"`
	SubMissionID          float64                    `json:"SubMissionID"`
	TemplateIDList        []float64                  `json:"TemplateIDList"`
}
type HeliobusSpecialPostLikes struct {
	HDILBDIPGHO float64 `json:"HDILBDIPGHO"`
	HEPALNIOJNP float64 `json:"HEPALNIOJNP"`
}
type HeliobusSpecialPostAccessor struct {
	_data                      []HeliobusSpecialPost
	_dataHeliobusSpecialPostID map[float64]HeliobusSpecialPost
	_dataSubMissionID          map[float64]HeliobusSpecialPost
}

// LoadData retrieves the data. Must be called before HeliobusSpecialPost.GroupData
func (a *HeliobusSpecialPostAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HeliobusSpecialPost.json")
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
func (a *HeliobusSpecialPostAccessor) Raw() ([]HeliobusSpecialPost, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HeliobusSpecialPost{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HeliobusSpecialPostAccessor.LoadData to preload everything
func (a *HeliobusSpecialPostAccessor) GroupData() {
	a._dataHeliobusSpecialPostID = map[float64]HeliobusSpecialPost{}
	a._dataSubMissionID = map[float64]HeliobusSpecialPost{}
	for _, d := range a._data {
		a._dataHeliobusSpecialPostID[d.HeliobusSpecialPostID] = d
		a._dataSubMissionID[d.SubMissionID] = d
	}
}

// ByHeliobusSpecialPostID returns the HeliobusSpecialPost uniquely identified by HeliobusSpecialPostID
//
// Error is only non-nil if the source errors out
func (a *HeliobusSpecialPostAccessor) ByHeliobusSpecialPostID(identifier float64) (HeliobusSpecialPost, error) {
	if a._dataHeliobusSpecialPostID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusSpecialPost{}, err
			}
		}
		a.GroupData()
	}
	return a._dataHeliobusSpecialPostID[identifier], nil
}

// BySubMissionID returns the HeliobusSpecialPost uniquely identified by SubMissionID
//
// Error is only non-nil if the source errors out
func (a *HeliobusSpecialPostAccessor) BySubMissionID(identifier float64) (HeliobusSpecialPost, error) {
	if a._dataSubMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusSpecialPost{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSubMissionID[identifier], nil
}
