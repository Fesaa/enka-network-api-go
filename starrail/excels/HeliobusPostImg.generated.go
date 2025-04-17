package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type HeliobusPostImg struct {
	PostImgID   float64 `json:"PostImgID"`
	PostImgPath string  `json:"PostImgPath"`
}
type HeliobusPostImgAccessor struct {
	_data            []HeliobusPostImg
	_dataPostImgID   map[float64]HeliobusPostImg
	_dataPostImgPath map[string]HeliobusPostImg
}

// LoadData retrieves the data. Must be called before HeliobusPostImg.GroupData
func (a *HeliobusPostImgAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HeliobusPostImg.json")
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
func (a *HeliobusPostImgAccessor) Raw() ([]HeliobusPostImg, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HeliobusPostImg{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HeliobusPostImgAccessor.LoadData to preload everything
func (a *HeliobusPostImgAccessor) GroupData() {
	a._dataPostImgID = map[float64]HeliobusPostImg{}
	a._dataPostImgPath = map[string]HeliobusPostImg{}
	for _, d := range a._data {
		a._dataPostImgID[d.PostImgID] = d
		a._dataPostImgPath[d.PostImgPath] = d
	}
}

// ByPostImgID returns the HeliobusPostImg uniquely identified by PostImgID
//
// Error is only non-nil if the source errors out
func (a *HeliobusPostImgAccessor) ByPostImgID(identifier float64) (HeliobusPostImg, error) {
	if a._dataPostImgID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusPostImg{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPostImgID[identifier], nil
}

// ByPostImgPath returns the HeliobusPostImg uniquely identified by PostImgPath
//
// Error is only non-nil if the source errors out
func (a *HeliobusPostImgAccessor) ByPostImgPath(identifier string) (HeliobusPostImg, error) {
	if a._dataPostImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusPostImg{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPostImgPath[identifier], nil
}
