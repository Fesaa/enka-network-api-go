package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ItemConfigTrainDynamic struct {
	CustomDataList       []interface{} `json:"CustomDataList"`
	ID                   float64       `json:"ID"`
	InventoryDisplayTag  float64       `json:"InventoryDisplayTag"`
	ItemAvatarIconPath   string        `json:"ItemAvatarIconPath"`
	ItemBGDesc           hash.Hash     `json:"ItemBGDesc"`
	ItemCurrencyIconPath string        `json:"ItemCurrencyIconPath"`
	ItemFigureIconPath   string        `json:"ItemFigureIconPath"`
	ItemGroup            float64       `json:"ItemGroup"`
	ItemIconPath         string        `json:"ItemIconPath"`
	ItemMainType         string        `json:"ItemMainType"`
	ItemName             hash.Hash     `json:"ItemName"`
	ItemSubType          string        `json:"ItemSubType"`
	PileLimit            float64       `json:"PileLimit"`
	Rarity               string        `json:"Rarity"`
	ReturnItemIDList     []interface{} `json:"ReturnItemIDList"`
}
type ItemConfigTrainDynamicAccessor struct {
	_data                   []ItemConfigTrainDynamic
	_dataID                 map[float64]ItemConfigTrainDynamic
	_dataItemFigureIconPath map[string]ItemConfigTrainDynamic
	_dataItemIconPath       map[string]ItemConfigTrainDynamic
}

// LoadData retrieves the data. Must be called before ItemConfigTrainDynamic.GroupData
func (a *ItemConfigTrainDynamicAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemConfigTrainDynamic.json")
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
func (a *ItemConfigTrainDynamicAccessor) Raw() ([]ItemConfigTrainDynamic, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemConfigTrainDynamic{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ItemConfigTrainDynamicAccessor.LoadData to preload everything
func (a *ItemConfigTrainDynamicAccessor) GroupData() {
	a._dataID = map[float64]ItemConfigTrainDynamic{}
	a._dataItemFigureIconPath = map[string]ItemConfigTrainDynamic{}
	a._dataItemIconPath = map[string]ItemConfigTrainDynamic{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataItemFigureIconPath[d.ItemFigureIconPath] = d
		a._dataItemIconPath[d.ItemIconPath] = d
	}
}

// ByID returns the ItemConfigTrainDynamic uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ItemConfigTrainDynamicAccessor) ByID(identifier float64) (ItemConfigTrainDynamic, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigTrainDynamic{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByItemFigureIconPath returns the ItemConfigTrainDynamic uniquely identified by ItemFigureIconPath
//
// Error is only non-nil if the source errors out
func (a *ItemConfigTrainDynamicAccessor) ByItemFigureIconPath(identifier string) (ItemConfigTrainDynamic, error) {
	if a._dataItemFigureIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigTrainDynamic{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemFigureIconPath[identifier], nil
}

// ByItemIconPath returns the ItemConfigTrainDynamic uniquely identified by ItemIconPath
//
// Error is only non-nil if the source errors out
func (a *ItemConfigTrainDynamicAccessor) ByItemIconPath(identifier string) (ItemConfigTrainDynamic, error) {
	if a._dataItemIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigTrainDynamic{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemIconPath[identifier], nil
}
