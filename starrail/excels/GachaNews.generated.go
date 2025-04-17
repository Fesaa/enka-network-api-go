package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type GachaNews struct {
	AvatarList []json.Number          `json:"AvatarList"`
	DecideID   json.Number            `json:"DecideID"`
	Desc       map[string]json.Number `json:"Desc"`
	NewsID     json.Number            `json:"NewsID"`
	Title      map[string]json.Number `json:"Title"`
}
type GachaNewsAccessor struct {
	_data         []GachaNews
	_dataDecideID map[json.Number]GachaNews
	_dataNewsID   map[json.Number]GachaNews
}

// LoadData retrieves the data. Must be called before GachaNews.GroupData
func (a *GachaNewsAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/GachaNews.json")
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
func (a *GachaNewsAccessor) Raw() ([]GachaNews, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []GachaNews{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with GachaNewsAccessor.LoadData to preload everything
func (a *GachaNewsAccessor) GroupData() {
	a._dataDecideID = map[json.Number]GachaNews{}
	a._dataNewsID = map[json.Number]GachaNews{}
	for _, d := range a._data {
		a._dataDecideID[d.DecideID] = d
		a._dataNewsID[d.NewsID] = d
	}
}

// ByDecideID returns the GachaNews uniquely identified by DecideID
//
// Error is only non-nil if the source errors out
func (a *GachaNewsAccessor) ByDecideID(identifier json.Number) (GachaNews, error) {
	if a._dataDecideID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GachaNews{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDecideID[identifier], nil
}

// ByNewsID returns the GachaNews uniquely identified by NewsID
//
// Error is only non-nil if the source errors out
func (a *GachaNewsAccessor) ByNewsID(identifier json.Number) (GachaNews, error) {
	if a._dataNewsID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GachaNews{}, err
			}
		}
		a.GroupData()
	}
	return a._dataNewsID[identifier], nil
}
