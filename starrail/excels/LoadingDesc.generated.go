package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type LoadingDesc struct {
	DescTextmapID    map[string]json.Number        `json:"DescTextmapID"`
	ForceParam       []LoadingDescForceParam       `json:"ForceParam"`
	ForceParamForOr  []LoadingDescForceParamForOr  `json:"ForceParamForOr"`
	Group            string                        `json:"Group"`
	ID               json.Number                   `json:"ID"`
	ImageID          json.Number                   `json:"ImageID"`
	LockParam        []LoadingDescLockParam        `json:"LockParam"`
	LockParamForOr   []LoadingDescLockParamForOr   `json:"LockParamForOr"`
	MaxLevel         json.Number                   `json:"MaxLevel"`
	MinLevel         json.Number                   `json:"MinLevel"`
	TitleTextmapID   map[string]json.Number        `json:"TitleTextmapID"`
	UnlockParam      []LoadingDescUnlockParam      `json:"UnlockParam"`
	UnlockParamForOr []LoadingDescUnlockParamForOr `json:"UnlockParamForOr"`
	Weight           json.Number                   `json:"Weight"`
}
type LoadingDescForceParam struct {
	CAOAPDCCPCA json.Number `json:"CAOAPDCCPCA"`
	PICHIHHCOCB string      `json:"PICHIHHCOCB"`
}
type LoadingDescForceParamForOr struct {
	CAOAPDCCPCA json.Number `json:"CAOAPDCCPCA"`
	PICHIHHCOCB string      `json:"PICHIHHCOCB"`
}
type LoadingDescLockParam struct {
	CAOAPDCCPCA json.Number `json:"CAOAPDCCPCA"`
	PICHIHHCOCB string      `json:"PICHIHHCOCB"`
}
type LoadingDescLockParamForOr struct {
	CAOAPDCCPCA json.Number `json:"CAOAPDCCPCA"`
	PICHIHHCOCB string      `json:"PICHIHHCOCB"`
}
type LoadingDescUnlockParam struct {
	CAOAPDCCPCA json.Number `json:"CAOAPDCCPCA"`
	PICHIHHCOCB string      `json:"PICHIHHCOCB"`
}
type LoadingDescUnlockParamForOr struct {
	CAOAPDCCPCA json.Number `json:"CAOAPDCCPCA"`
	PICHIHHCOCB string      `json:"PICHIHHCOCB"`
}
type LoadingDescAccessor struct {
	_data []LoadingDesc
}

// LoadData retrieves the data. Must be called before LoadingDesc.GroupData
func (a *LoadingDescAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/LoadingDesc.json")
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
func (a *LoadingDescAccessor) Raw() ([]LoadingDesc, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []LoadingDesc{}, err
		}
	}
	return a._data, nil
}
