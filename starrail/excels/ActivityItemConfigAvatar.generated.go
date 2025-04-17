package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ActivityItemConfigAvatar struct {
	CustomDataList       []interface{}          `json:"CustomDataList"`
	ID                   json.Number            `json:"ID"`
	InventoryDisplayTag  json.Number            `json:"InventoryDisplayTag"`
	ItemAvatarIconPath   string                 `json:"ItemAvatarIconPath"`
	ItemBGDesc           map[string]json.Number `json:"ItemBGDesc"`
	ItemCurrencyIconPath string                 `json:"ItemCurrencyIconPath"`
	ItemFigureIconPath   string                 `json:"ItemFigureIconPath"`
	ItemIconPath         string                 `json:"ItemIconPath"`
	ItemMainType         string                 `json:"ItemMainType"`
	ItemName             map[string]json.Number `json:"ItemName"`
	ItemSubType          string                 `json:"ItemSubType"`
	PileLimit            json.Number            `json:"PileLimit"`
	Rarity               string                 `json:"Rarity"`
	ReturnItemIDList     []interface{}          `json:"ReturnItemIDList"`
}
type ActivityItemConfigAvatarAccessor struct {
	_data                   []ActivityItemConfigAvatar
	_dataItemAvatarIconPath map[string]ActivityItemConfigAvatar
	_dataItemFigureIconPath map[string]ActivityItemConfigAvatar
	_dataItemIconPath       map[string]ActivityItemConfigAvatar
}

// LoadData retrieves the data. Must be called before ActivityItemConfigAvatar.GroupData
func (a *ActivityItemConfigAvatarAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityItemConfigAvatar.json")
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
func (a *ActivityItemConfigAvatarAccessor) Raw() ([]ActivityItemConfigAvatar, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityItemConfigAvatar{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityItemConfigAvatarAccessor.LoadData to preload everything
func (a *ActivityItemConfigAvatarAccessor) GroupData() {
	a._dataItemAvatarIconPath = map[string]ActivityItemConfigAvatar{}
	a._dataItemFigureIconPath = map[string]ActivityItemConfigAvatar{}
	a._dataItemIconPath = map[string]ActivityItemConfigAvatar{}
	for _, d := range a._data {
		a._dataItemAvatarIconPath[d.ItemAvatarIconPath] = d
		a._dataItemFigureIconPath[d.ItemFigureIconPath] = d
		a._dataItemIconPath[d.ItemIconPath] = d
	}
}

// ByItemAvatarIconPath returns the ActivityItemConfigAvatar uniquely identified by ItemAvatarIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityItemConfigAvatarAccessor) ByItemAvatarIconPath(identifier string) (ActivityItemConfigAvatar, error) {
	if a._dataItemAvatarIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityItemConfigAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemAvatarIconPath[identifier], nil
}

// ByItemFigureIconPath returns the ActivityItemConfigAvatar uniquely identified by ItemFigureIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityItemConfigAvatarAccessor) ByItemFigureIconPath(identifier string) (ActivityItemConfigAvatar, error) {
	if a._dataItemFigureIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityItemConfigAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemFigureIconPath[identifier], nil
}

// ByItemIconPath returns the ActivityItemConfigAvatar uniquely identified by ItemIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityItemConfigAvatarAccessor) ByItemIconPath(identifier string) (ActivityItemConfigAvatar, error) {
	if a._dataItemIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityItemConfigAvatar{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemIconPath[identifier], nil
}
