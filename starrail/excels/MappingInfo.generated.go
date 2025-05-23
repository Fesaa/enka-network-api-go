package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MappingInfo struct {
	Desc            hash.Hash                    `json:"Desc"`
	DisplayItemList []MappingInfoDisplayItemList `json:"DisplayItemList"`
	FarmType        string                       `json:"FarmType"`
	ID              float64                      `json:"ID"`
	IsShowInFog     bool                         `json:"IsShowInFog"`
	Name            hash.Hash                    `json:"Name"`
	ShowMonsterList []float64                    `json:"ShowMonsterList"`
	Type            string                       `json:"Type"`
	WorldLevel      float64                      `json:"WorldLevel"`
}
type MappingInfoDisplayItemList struct {
	ItemID  float64 `json:"ItemID"`
	ItemNum float64 `json:"ItemNum"`
}
type MappingInfoAccessor struct {
	_data []MappingInfo
}

// LoadData retrieves the data. Must be called before MappingInfo.GroupData
func (a *MappingInfoAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MappingInfo.json")
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
func (a *MappingInfoAccessor) Raw() ([]MappingInfo, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MappingInfo{}, err
		}
	}
	return a._data, nil
}
