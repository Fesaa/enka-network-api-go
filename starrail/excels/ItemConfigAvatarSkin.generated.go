package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ItemConfigAvatarSkin struct {
	CustomDataList       []interface{} `json:"CustomDataList"`
	ID                   float64       `json:"ID"`
	InventoryDisplayTag  float64       `json:"InventoryDisplayTag"`
	ItemAvatarIconPath   string        `json:"ItemAvatarIconPath"`
	ItemCurrencyIconPath string        `json:"ItemCurrencyIconPath"`
	ItemDesc             hash.Hash     `json:"ItemDesc"`
	ItemFigureIconPath   string        `json:"ItemFigureIconPath"`
	ItemIconPath         string        `json:"ItemIconPath"`
	ItemMainType         string        `json:"ItemMainType"`
	ItemName             hash.Hash     `json:"ItemName"`
	ItemSubType          string        `json:"ItemSubType"`
	PileLimit            float64       `json:"PileLimit"`
	Rarity               string        `json:"Rarity"`
	ReturnItemIDList     []interface{} `json:"ReturnItemIDList"`
	IsVisible            bool          `json:"isVisible"`
}
type ItemConfigAvatarSkinAccessor struct {
	_data                     []ItemConfigAvatarSkin
	_dataID                   map[float64]ItemConfigAvatarSkin
	_dataInventoryDisplayTag  map[float64]ItemConfigAvatarSkin
	_dataItemAvatarIconPath   map[string]ItemConfigAvatarSkin
	_dataItemCurrencyIconPath map[string]ItemConfigAvatarSkin
	_dataItemFigureIconPath   map[string]ItemConfigAvatarSkin
	_dataItemIconPath         map[string]ItemConfigAvatarSkin
	_dataItemMainType         map[string]ItemConfigAvatarSkin
	_dataItemSubType          map[string]ItemConfigAvatarSkin
	_dataPileLimit            map[float64]ItemConfigAvatarSkin
	_dataRarity               map[string]ItemConfigAvatarSkin
}

// LoadData retrieves the data. Must be called before ItemConfigAvatarSkin.GroupData
func (a *ItemConfigAvatarSkinAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemConfigAvatarSkin.json")
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
func (a *ItemConfigAvatarSkinAccessor) Raw() ([]ItemConfigAvatarSkin, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemConfigAvatarSkin{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ItemConfigAvatarSkinAccessor.LoadData to preload everything
func (a *ItemConfigAvatarSkinAccessor) GroupData() {
	a._dataID = map[float64]ItemConfigAvatarSkin{}
	a._dataInventoryDisplayTag = map[float64]ItemConfigAvatarSkin{}
	a._dataItemAvatarIconPath = map[string]ItemConfigAvatarSkin{}
	a._dataItemCurrencyIconPath = map[string]ItemConfigAvatarSkin{}
	a._dataItemFigureIconPath = map[string]ItemConfigAvatarSkin{}
	a._dataItemIconPath = map[string]ItemConfigAvatarSkin{}
	a._dataItemMainType = map[string]ItemConfigAvatarSkin{}
	a._dataItemSubType = map[string]ItemConfigAvatarSkin{}
	a._dataPileLimit = map[float64]ItemConfigAvatarSkin{}
	a._dataRarity = map[string]ItemConfigAvatarSkin{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataInventoryDisplayTag[d.InventoryDisplayTag] = d
		a._dataItemAvatarIconPath[d.ItemAvatarIconPath] = d
		a._dataItemCurrencyIconPath[d.ItemCurrencyIconPath] = d
		a._dataItemFigureIconPath[d.ItemFigureIconPath] = d
		a._dataItemIconPath[d.ItemIconPath] = d
		a._dataItemMainType[d.ItemMainType] = d
		a._dataItemSubType[d.ItemSubType] = d
		a._dataPileLimit[d.PileLimit] = d
		a._dataRarity[d.Rarity] = d
	}
}

// ByID returns the ItemConfigAvatarSkin uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ItemConfigAvatarSkinAccessor) ByID(identifier float64) (ItemConfigAvatarSkin, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigAvatarSkin{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByInventoryDisplayTag returns the ItemConfigAvatarSkin uniquely identified by InventoryDisplayTag
//
// Error is only non-nil if the source errors out
func (a *ItemConfigAvatarSkinAccessor) ByInventoryDisplayTag(identifier float64) (ItemConfigAvatarSkin, error) {
	if a._dataInventoryDisplayTag == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigAvatarSkin{}, err
			}
		}
		a.GroupData()
	}
	return a._dataInventoryDisplayTag[identifier], nil
}

// ByItemAvatarIconPath returns the ItemConfigAvatarSkin uniquely identified by ItemAvatarIconPath
//
// Error is only non-nil if the source errors out
func (a *ItemConfigAvatarSkinAccessor) ByItemAvatarIconPath(identifier string) (ItemConfigAvatarSkin, error) {
	if a._dataItemAvatarIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigAvatarSkin{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemAvatarIconPath[identifier], nil
}

// ByItemCurrencyIconPath returns the ItemConfigAvatarSkin uniquely identified by ItemCurrencyIconPath
//
// Error is only non-nil if the source errors out
func (a *ItemConfigAvatarSkinAccessor) ByItemCurrencyIconPath(identifier string) (ItemConfigAvatarSkin, error) {
	if a._dataItemCurrencyIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigAvatarSkin{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemCurrencyIconPath[identifier], nil
}

// ByItemFigureIconPath returns the ItemConfigAvatarSkin uniquely identified by ItemFigureIconPath
//
// Error is only non-nil if the source errors out
func (a *ItemConfigAvatarSkinAccessor) ByItemFigureIconPath(identifier string) (ItemConfigAvatarSkin, error) {
	if a._dataItemFigureIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigAvatarSkin{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemFigureIconPath[identifier], nil
}

// ByItemIconPath returns the ItemConfigAvatarSkin uniquely identified by ItemIconPath
//
// Error is only non-nil if the source errors out
func (a *ItemConfigAvatarSkinAccessor) ByItemIconPath(identifier string) (ItemConfigAvatarSkin, error) {
	if a._dataItemIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigAvatarSkin{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemIconPath[identifier], nil
}

// ByItemMainType returns the ItemConfigAvatarSkin uniquely identified by ItemMainType
//
// Error is only non-nil if the source errors out
func (a *ItemConfigAvatarSkinAccessor) ByItemMainType(identifier string) (ItemConfigAvatarSkin, error) {
	if a._dataItemMainType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigAvatarSkin{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemMainType[identifier], nil
}

// ByItemSubType returns the ItemConfigAvatarSkin uniquely identified by ItemSubType
//
// Error is only non-nil if the source errors out
func (a *ItemConfigAvatarSkinAccessor) ByItemSubType(identifier string) (ItemConfigAvatarSkin, error) {
	if a._dataItemSubType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigAvatarSkin{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemSubType[identifier], nil
}

// ByPileLimit returns the ItemConfigAvatarSkin uniquely identified by PileLimit
//
// Error is only non-nil if the source errors out
func (a *ItemConfigAvatarSkinAccessor) ByPileLimit(identifier float64) (ItemConfigAvatarSkin, error) {
	if a._dataPileLimit == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigAvatarSkin{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPileLimit[identifier], nil
}

// ByRarity returns the ItemConfigAvatarSkin uniquely identified by Rarity
//
// Error is only non-nil if the source errors out
func (a *ItemConfigAvatarSkinAccessor) ByRarity(identifier string) (ItemConfigAvatarSkin, error) {
	if a._dataRarity == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ItemConfigAvatarSkin{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRarity[identifier], nil
}
