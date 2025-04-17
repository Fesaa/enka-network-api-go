package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MessageItemImage struct {
	FemaleImagePath string  `json:"FemaleImagePath"`
	ID              float64 `json:"ID"`
	ImagePath       string  `json:"ImagePath"`
}
type MessageItemImageAccessor struct {
	_data   []MessageItemImage
	_dataID map[float64]MessageItemImage
}

// LoadData retrieves the data. Must be called before MessageItemImage.GroupData
func (a *MessageItemImageAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MessageItemImage.json")
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
func (a *MessageItemImageAccessor) Raw() ([]MessageItemImage, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MessageItemImage{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MessageItemImageAccessor.LoadData to preload everything
func (a *MessageItemImageAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MessageItemImage uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MessageItemImageAccessor) ByID(identifier float64) (MessageItemImage, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return MessageItemImage{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
