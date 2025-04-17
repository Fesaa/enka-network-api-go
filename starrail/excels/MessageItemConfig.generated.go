package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MessageItemConfig struct {
	ContactsID     float64   `json:"ContactsID"`
	ID             float64   `json:"ID"`
	ItemContentID  float64   `json:"ItemContentID"`
	ItemType       string    `json:"ItemType"`
	MainText       hash.Hash `json:"MainText"`
	NextItemIDList []float64 `json:"NextItemIDList"`
	OptionText     hash.Hash `json:"OptionText"`
	SectionID      float64   `json:"SectionID"`
	Sender         string    `json:"Sender"`
}
type MessageItemConfigAccessor struct {
	_data   []MessageItemConfig
	_dataID map[float64]MessageItemConfig
}

// LoadData retrieves the data. Must be called before MessageItemConfig.GroupData
func (a *MessageItemConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MessageItemConfig.json")
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
func (a *MessageItemConfigAccessor) Raw() ([]MessageItemConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MessageItemConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MessageItemConfigAccessor.LoadData to preload everything
func (a *MessageItemConfigAccessor) GroupData() {
	a._dataID = map[float64]MessageItemConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MessageItemConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MessageItemConfigAccessor) ByID(identifier float64) (MessageItemConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MessageItemConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
