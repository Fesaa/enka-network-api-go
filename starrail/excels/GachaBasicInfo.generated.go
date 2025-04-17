package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type GachaBasicInfo struct {
	EndTime               string                 `json:"EndTime"`
	GachaID               json.Number            `json:"GachaID"`
	GachaType             string                 `json:"GachaType"`
	PoolDesc              map[string]json.Number `json:"PoolDesc"`
	PoolLabelIcon         string                 `json:"PoolLabelIcon"`
	PoolLabelIconSelected string                 `json:"PoolLabelIconSelected"`
	PoolName              map[string]json.Number `json:"PoolName"`
	PrefabPath            string                 `json:"PrefabPath"`
	SortID                json.Number            `json:"SortID"`
	StartTime             string                 `json:"StartTime"`
	TypeTitle             map[string]json.Number `json:"TypeTitle"`
}
type GachaBasicInfoAccessor struct {
	_data []GachaBasicInfo
}

// LoadData retrieves the data. Must be called before GachaBasicInfo.GroupData
func (a *GachaBasicInfoAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/GachaBasicInfo.json")
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
func (a *GachaBasicInfoAccessor) Raw() ([]GachaBasicInfo, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []GachaBasicInfo{}, err
		}
	}
	return a._data, nil
}
