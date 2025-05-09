package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AvatarPropertyConfig struct {
	IconPath              string    `json:"IconPath"`
	IsDisplay             bool      `json:"IsDisplay"`
	MainRelicFilter       float64   `json:"MainRelicFilter"`
	Order                 float64   `json:"Order"`
	PropertyClassify      float64   `json:"PropertyClassify"`
	PropertyInstructionID float64   `json:"PropertyInstructionID"`
	PropertyName          hash.Hash `json:"PropertyName"`
	PropertyNameFilter    hash.Hash `json:"PropertyNameFilter"`
	PropertyNameRelic     hash.Hash `json:"PropertyNameRelic"`
	PropertyNameSkillTree hash.Hash `json:"PropertyNameSkillTree"`
	PropertyType          string    `json:"PropertyType"`
	SubRelicFilter        float64   `json:"SubRelicFilter"`
	IsBattleDisplay       bool      `json:"isBattleDisplay"`
}
type AvatarPropertyConfigAccessor struct {
	_data             []AvatarPropertyConfig
	_dataPropertyType map[string]AvatarPropertyConfig
}

// LoadData retrieves the data. Must be called before AvatarPropertyConfig.GroupData
func (a *AvatarPropertyConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarPropertyConfig.json")
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
func (a *AvatarPropertyConfigAccessor) Raw() ([]AvatarPropertyConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarPropertyConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarPropertyConfigAccessor.LoadData to preload everything
func (a *AvatarPropertyConfigAccessor) GroupData() {
	a._dataPropertyType = map[string]AvatarPropertyConfig{}
	for _, d := range a._data {
		a._dataPropertyType[d.PropertyType] = d
	}
}

// ByPropertyType returns the AvatarPropertyConfig uniquely identified by PropertyType
//
// Error is only non-nil if the source errors out
func (a *AvatarPropertyConfigAccessor) ByPropertyType(identifier string) (AvatarPropertyConfig, error) {
	if a._dataPropertyType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarPropertyConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPropertyType[identifier], nil
}
