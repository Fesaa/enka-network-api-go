package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type AlleySpecialOrder struct {
	OrderPic            string                               `json:"OrderPic"`
	OrderTips           string                               `json:"OrderTips"`
	OrderTipsTime       []float64                            `json:"OrderTipsTime"`
	SpecialOrderContent []float64                            `json:"SpecialOrderContent"`
	SpecialOrderGoods   []AlleySpecialOrderSpecialOrderGoods `json:"SpecialOrderGoods"`
	SpecialOrderID      float64                              `json:"SpecialOrderID"`
	SpecialOrderReward  float64                              `json:"SpecialOrderReward"`
	SpecialOrderShip    float64                              `json:"SpecialOrderShip"`
	SpecialOrderShopID  float64                              `json:"SpecialOrderShopID"`
	SubTitleID          string                               `json:"SubTitleID"`
	UnlockMission       float64                              `json:"UnlockMission"`
}
type AlleySpecialOrderSpecialOrderGoods struct {
	GoodsCnt float64 `json:"GoodsCnt"`
	GoodsID  float64 `json:"GoodsID"`
}
type AlleySpecialOrderAccessor struct {
	_data                   []AlleySpecialOrder
	_dataSpecialOrderShopID map[float64]AlleySpecialOrder
	_dataSpecialOrderID     map[float64]AlleySpecialOrder
	_dataOrderPic           map[string]AlleySpecialOrder
	_dataSubTitleID         map[string]AlleySpecialOrder
	_dataOrderTips          map[string]AlleySpecialOrder
}

// LoadData retrieves the data. Must be called before AlleySpecialOrder.GroupData
func (a *AlleySpecialOrderAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AlleySpecialOrder.json")
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
func (a *AlleySpecialOrderAccessor) Raw() ([]AlleySpecialOrder, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AlleySpecialOrder{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AlleySpecialOrderAccessor.LoadData to preload everything
func (a *AlleySpecialOrderAccessor) GroupData() {
	for _, d := range a._data {
		a._dataSpecialOrderShopID[d.SpecialOrderShopID] = d
		a._dataSpecialOrderID[d.SpecialOrderID] = d
		a._dataOrderPic[d.OrderPic] = d
		a._dataSubTitleID[d.SubTitleID] = d
		a._dataOrderTips[d.OrderTips] = d
	}
}

// BySpecialOrderShopID returns the AlleySpecialOrder uniquely identified by SpecialOrderShopID
//
// Error is only non-nil if the source errors out
func (a *AlleySpecialOrderAccessor) BySpecialOrderShopID(identifier float64) (AlleySpecialOrder, error) {
	if a._dataSpecialOrderShopID == nil {
		err := a.LoadData()
		if err != nil {
			return AlleySpecialOrder{}, err
		}
		a.GroupData()
	}
	return a._dataSpecialOrderShopID[identifier], nil
}

// BySpecialOrderID returns the AlleySpecialOrder uniquely identified by SpecialOrderID
//
// Error is only non-nil if the source errors out
func (a *AlleySpecialOrderAccessor) BySpecialOrderID(identifier float64) (AlleySpecialOrder, error) {
	if a._dataSpecialOrderID == nil {
		err := a.LoadData()
		if err != nil {
			return AlleySpecialOrder{}, err
		}
		a.GroupData()
	}
	return a._dataSpecialOrderID[identifier], nil
}

// ByOrderPic returns the AlleySpecialOrder uniquely identified by OrderPic
//
// Error is only non-nil if the source errors out
func (a *AlleySpecialOrderAccessor) ByOrderPic(identifier string) (AlleySpecialOrder, error) {
	if a._dataOrderPic == nil {
		err := a.LoadData()
		if err != nil {
			return AlleySpecialOrder{}, err
		}
		a.GroupData()
	}
	return a._dataOrderPic[identifier], nil
}

// BySubTitleID returns the AlleySpecialOrder uniquely identified by SubTitleID
//
// Error is only non-nil if the source errors out
func (a *AlleySpecialOrderAccessor) BySubTitleID(identifier string) (AlleySpecialOrder, error) {
	if a._dataSubTitleID == nil {
		err := a.LoadData()
		if err != nil {
			return AlleySpecialOrder{}, err
		}
		a.GroupData()
	}
	return a._dataSubTitleID[identifier], nil
}

// ByOrderTips returns the AlleySpecialOrder uniquely identified by OrderTips
//
// Error is only non-nil if the source errors out
func (a *AlleySpecialOrderAccessor) ByOrderTips(identifier string) (AlleySpecialOrder, error) {
	if a._dataOrderTips == nil {
		err := a.LoadData()
		if err != nil {
			return AlleySpecialOrder{}, err
		}
		a.GroupData()
	}
	return a._dataOrderTips[identifier], nil
}
