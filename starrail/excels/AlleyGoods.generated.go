package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type AlleyGoods struct {
	GoodsConfig      string      `json:"GoodsConfig"`
	GoodsID          json.Number `json:"GoodsID"`
	GoodsPic         string      `json:"GoodsPic"`
	GoodsPicLocked   string      `json:"GoodsPicLocked"`
	GoodsProfit      json.Number `json:"GoodsProfit"`
	RotateAudioEvent string      `json:"RotateAudioEvent"`
	SettleAudioEvent string      `json:"SettleAudioEvent"`
}
type AlleyGoodsAccessor struct {
	_data        []AlleyGoods
	_dataGoodsID map[json.Number]AlleyGoods
}

// LoadData retrieves the data. Must be called before AlleyGoods.GroupData
func (a *AlleyGoodsAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AlleyGoods.json")
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
func (a *AlleyGoodsAccessor) Raw() ([]AlleyGoods, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AlleyGoods{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AlleyGoodsAccessor.LoadData to preload everything
func (a *AlleyGoodsAccessor) GroupData() {
	a._dataGoodsID = map[json.Number]AlleyGoods{}
	for _, d := range a._data {
		a._dataGoodsID[d.GoodsID] = d
	}
}

// ByGoodsID returns the AlleyGoods uniquely identified by GoodsID
//
// Error is only non-nil if the source errors out
func (a *AlleyGoodsAccessor) ByGoodsID(identifier json.Number) (AlleyGoods, error) {
	if a._dataGoodsID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AlleyGoods{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGoodsID[identifier], nil
}
