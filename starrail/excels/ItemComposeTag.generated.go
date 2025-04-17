package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ItemComposeTag struct {
	BelongTypeID     float64 `json:"BelongTypeID"`
	ItemComposeTagID float64 `json:"ItemComposeTagID"`
	TagTextmapID     string  `json:"TagTextmapID"`
}
type ItemComposeTagAccessor struct {
	_data                 []ItemComposeTag
	_dataItemComposeTagID map[float64]ItemComposeTag
	_dataTagTextmapID     map[string]ItemComposeTag
}

// LoadData retrieves the data. Must be called before ItemComposeTag.GroupData
func (a *ItemComposeTagAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemComposeTag.json")
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
func (a *ItemComposeTagAccessor) Raw() ([]ItemComposeTag, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemComposeTag{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ItemComposeTagAccessor.LoadData to preload everything
func (a *ItemComposeTagAccessor) GroupData() {
	a._dataItemComposeTagID = map[float64]ItemComposeTag{}
	a._dataTagTextmapID = map[string]ItemComposeTag{}
	for _, d := range a._data {
		a._dataItemComposeTagID[d.ItemComposeTagID] = d
		a._dataTagTextmapID[d.TagTextmapID] = d
	}
}

// ByItemComposeTagID returns the ItemComposeTag uniquely identified by ItemComposeTagID
//
// Error is only non-nil if the source errors out
func (a *ItemComposeTagAccessor) ByItemComposeTagID(identifier float64) (ItemComposeTag, error) {
	if a._dataItemComposeTagID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemComposeTag{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemComposeTagID[identifier], nil
}

// ByTagTextmapID returns the ItemComposeTag uniquely identified by TagTextmapID
//
// Error is only non-nil if the source errors out
func (a *ItemComposeTagAccessor) ByTagTextmapID(identifier string) (ItemComposeTag, error) {
	if a._dataTagTextmapID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemComposeTag{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTagTextmapID[identifier], nil
}
