package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MarbleRandomBuff struct {
	ConditionList    []float64 `json:"ConditionList"`
	Desc             hash.Hash `json:"Desc"`
	EffectParam      float64   `json:"EffectParam"`
	EffectType       float64   `json:"EffectType"`
	ID               float64   `json:"ID"`
	IconPath         string    `json:"IconPath"`
	IsRepeat         float64   `json:"IsRepeat"`
	Name             hash.Hash `json:"Name"`
	ParamList        []float64 `json:"ParamList"`
	SetInactive      float64   `json:"SetInactive"`
	UnlockSubMission float64   `json:"UnlockSubMission"`
	Weight           float64   `json:"Weight"`
}
type MarbleRandomBuffAccessor struct {
	_data            []MarbleRandomBuff
	_dataID          map[float64]MarbleRandomBuff
	_dataEffectParam map[float64]MarbleRandomBuff
}

// LoadData retrieves the data. Must be called before MarbleRandomBuff.GroupData
func (a *MarbleRandomBuffAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MarbleRandomBuff.json")
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
func (a *MarbleRandomBuffAccessor) Raw() ([]MarbleRandomBuff, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MarbleRandomBuff{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MarbleRandomBuffAccessor.LoadData to preload everything
func (a *MarbleRandomBuffAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataEffectParam[d.EffectParam] = d
	}
}

// ByID returns the MarbleRandomBuff uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MarbleRandomBuffAccessor) ByID(identifier float64) (MarbleRandomBuff, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return MarbleRandomBuff{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByEffectParam returns the MarbleRandomBuff uniquely identified by EffectParam
//
// Error is only non-nil if the source errors out
func (a *MarbleRandomBuffAccessor) ByEffectParam(identifier float64) (MarbleRandomBuff, error) {
	if a._dataEffectParam == nil {
		err := a.LoadData()
		if err != nil {
			return MarbleRandomBuff{}, err
		}
		a.GroupData()
	}
	return a._dataEffectParam[identifier], nil
}
