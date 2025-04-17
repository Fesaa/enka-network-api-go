package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ItemConfig struct {
	CustomDataList       []json.Number          `json:"CustomDataList"`
	ID                   json.Number            `json:"ID"`
	InventoryDisplayTag  json.Number            `json:"InventoryDisplayTag"`
	IsShowRedDot         bool                   `json:"IsShowRedDot"`
	ItemAvatarIconPath   string                 `json:"ItemAvatarIconPath"`
	ItemBGDesc           map[string]json.Number `json:"ItemBGDesc"`
	ItemCurrencyIconPath string                 `json:"ItemCurrencyIconPath"`
	ItemDesc             map[string]json.Number `json:"ItemDesc"`
	ItemFigureIconPath   string                 `json:"ItemFigureIconPath"`
	ItemGroup            json.Number            `json:"ItemGroup"`
	ItemIconPath         string                 `json:"ItemIconPath"`
	ItemMainType         string                 `json:"ItemMainType"`
	ItemName             map[string]json.Number `json:"ItemName"`
	ItemSubType          string                 `json:"ItemSubType"`
	PileLimit            json.Number            `json:"PileLimit"`
	PurposeType          json.Number            `json:"PurposeType"`
	Rarity               string                 `json:"Rarity"`
	ReturnItemIDList     []interface{}          `json:"ReturnItemIDList"`
	SellType             string                 `json:"SellType"`
	UseMethod            string                 `json:"UseMethod"`
	IsVisible            bool                   `json:"isVisible"`
}
type ItemConfigAccessor struct {
	_data []ItemConfig
}

// LoadData retrieves the data. Must be called before ItemConfig.GroupData
func (a *ItemConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemConfig.json")
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
func (a *ItemConfigAccessor) Raw() ([]ItemConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemConfig{}, err
		}
	}
	return a._data, nil
}
