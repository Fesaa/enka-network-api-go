package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type BattlePassAdvertisement struct {
	Desc           map[string]json.Number `json:"Desc"`
	ID             json.Number            `json:"ID"`
	IconBundlePath string                 `json:"IconBundlePath"`
	Title          map[string]json.Number `json:"Title"`
}
type BattlePassAdvertisementAccessor struct {
	_data               []BattlePassAdvertisement
	_dataIconBundlePath map[string]BattlePassAdvertisement
}

// LoadData retrieves the data. Must be called before BattlePassAdvertisement.GroupData
func (a *BattlePassAdvertisementAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattlePassAdvertisement.json")
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
func (a *BattlePassAdvertisementAccessor) Raw() ([]BattlePassAdvertisement, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattlePassAdvertisement{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattlePassAdvertisementAccessor.LoadData to preload everything
func (a *BattlePassAdvertisementAccessor) GroupData() {
	a._dataIconBundlePath = map[string]BattlePassAdvertisement{}
	for _, d := range a._data {
		a._dataIconBundlePath[d.IconBundlePath] = d
	}
}

// ByIconBundlePath returns the BattlePassAdvertisement uniquely identified by IconBundlePath
//
// Error is only non-nil if the source errors out
func (a *BattlePassAdvertisementAccessor) ByIconBundlePath(identifier string) (BattlePassAdvertisement, error) {
	if a._dataIconBundlePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattlePassAdvertisement{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconBundlePath[identifier], nil
}
