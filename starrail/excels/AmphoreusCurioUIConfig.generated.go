package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AmphoreusCurioUIConfig struct {
	Desc          hash.Hash   `json:"Desc"`
	ID            float64     `json:"ID"`
	IconPath      string      `json:"IconPath"`
	Name          hash.Hash   `json:"Name"`
	NameAfter     hash.Hash   `json:"NameAfter"`
	ReplyIDList   []float64   `json:"ReplyIDList"`
	Tag           float64     `json:"Tag"`
	TextmapIDList []hash.Hash `json:"TextmapIDList"`
}
type AmphoreusCurioUIConfigAccessor struct {
	_data         []AmphoreusCurioUIConfig
	_dataID       map[float64]AmphoreusCurioUIConfig
	_dataIconPath map[string]AmphoreusCurioUIConfig
}

// LoadData retrieves the data. Must be called before AmphoreusCurioUIConfig.GroupData
func (a *AmphoreusCurioUIConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AmphoreusCurioUIConfig.json")
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
func (a *AmphoreusCurioUIConfigAccessor) Raw() ([]AmphoreusCurioUIConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AmphoreusCurioUIConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AmphoreusCurioUIConfigAccessor.LoadData to preload everything
func (a *AmphoreusCurioUIConfigAccessor) GroupData() {
	a._dataID = map[float64]AmphoreusCurioUIConfig{}
	a._dataIconPath = map[string]AmphoreusCurioUIConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataIconPath[d.IconPath] = d
	}
}

// ByID returns the AmphoreusCurioUIConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *AmphoreusCurioUIConfigAccessor) ByID(identifier float64) (AmphoreusCurioUIConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AmphoreusCurioUIConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByIconPath returns the AmphoreusCurioUIConfig uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *AmphoreusCurioUIConfigAccessor) ByIconPath(identifier string) (AmphoreusCurioUIConfig, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AmphoreusCurioUIConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}
