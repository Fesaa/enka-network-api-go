package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MonopolyGoodsConfig struct {
	Cost              float64   `json:"Cost"`
	Desc              hash.Hash `json:"Desc"`
	GoodsID           float64   `json:"GoodsID"`
	GoodsType         string    `json:"GoodsType"`
	IconPath          string    `json:"IconPath"`
	Name              hash.Hash `json:"Name"`
	TextDisplayParam1 float64   `json:"TextDisplayParam1"`
	TextDisplayParam2 float64   `json:"TextDisplayParam2"`
}
type MonopolyGoodsConfigAccessor struct {
	_data        []MonopolyGoodsConfig
	_dataGoodsID map[float64]MonopolyGoodsConfig
}

// LoadData retrieves the data. Must be called before MonopolyGoodsConfig.GroupData
func (a *MonopolyGoodsConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonopolyGoodsConfig.json")
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
func (a *MonopolyGoodsConfigAccessor) Raw() ([]MonopolyGoodsConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonopolyGoodsConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonopolyGoodsConfigAccessor.LoadData to preload everything
func (a *MonopolyGoodsConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGoodsID[d.GoodsID] = d
	}
}

// ByGoodsID returns the MonopolyGoodsConfig uniquely identified by GoodsID
//
// Error is only non-nil if the source errors out
func (a *MonopolyGoodsConfigAccessor) ByGoodsID(identifier float64) (MonopolyGoodsConfig, error) {
	if a._dataGoodsID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MonopolyGoodsConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGoodsID[identifier], nil
}
