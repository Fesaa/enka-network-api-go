package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ItemConfigBook struct {
	CustomDataList       []interface{} `json:"CustomDataList"`
	ID                   float64       `json:"ID"`
	InventoryDisplayTag  float64       `json:"InventoryDisplayTag"`
	ItemAvatarIconPath   string        `json:"ItemAvatarIconPath"`
	ItemBGDesc           hash.Hash     `json:"ItemBGDesc"`
	ItemCurrencyIconPath string        `json:"ItemCurrencyIconPath"`
	ItemDesc             hash.Hash     `json:"ItemDesc"`
	ItemFigureIconPath   string        `json:"ItemFigureIconPath"`
	ItemIconPath         string        `json:"ItemIconPath"`
	ItemMainType         string        `json:"ItemMainType"`
	ItemName             hash.Hash     `json:"ItemName"`
	ItemSubType          string        `json:"ItemSubType"`
	PileLimit            float64       `json:"PileLimit"`
	PurposeType          float64       `json:"PurposeType"`
	Rarity               string        `json:"Rarity"`
	ReturnItemIDList     []interface{} `json:"ReturnItemIDList"`
	UseMethod            string        `json:"UseMethod"`
	IsVisible            bool          `json:"isVisible"`
}
type ItemConfigBookAccessor struct {
	_data   []ItemConfigBook
	_dataID map[float64]ItemConfigBook
}

// LoadData retrieves the data. Must be called before ItemConfigBook.GroupData
func (a *ItemConfigBookAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemConfigBook.json")
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
func (a *ItemConfigBookAccessor) Raw() ([]ItemConfigBook, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemConfigBook{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ItemConfigBookAccessor.LoadData to preload everything
func (a *ItemConfigBookAccessor) GroupData() {
	a._dataID = map[float64]ItemConfigBook{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the ItemConfigBook uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ItemConfigBookAccessor) ByID(identifier float64) (ItemConfigBook, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigBook{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
