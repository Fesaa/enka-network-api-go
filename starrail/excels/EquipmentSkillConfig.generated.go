package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type EquipmentSkillConfig struct {
	AbilityName     string                                `json:"AbilityName"`
	AbilityProperty []EquipmentSkillConfigAbilityProperty `json:"AbilityProperty"`
	Level           float64                               `json:"Level"`
	ParamList       []hash.IntValue                       `json:"ParamList"`
	SkillDesc       hash.Hash                             `json:"SkillDesc"`
	SkillID         float64                               `json:"SkillID"`
	SkillName       hash.Hash                             `json:"SkillName"`
}
type EquipmentSkillConfigAbilityProperty struct {
	PropertyType string        `json:"PropertyType"`
	Value        hash.IntValue `json:"Value"`
}
type EquipmentSkillConfigAccessor struct {
	_data []EquipmentSkillConfig
}

// LoadData retrieves the data. Must be called before EquipmentSkillConfig.GroupData
func (a *EquipmentSkillConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EquipmentSkillConfig.json")
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
func (a *EquipmentSkillConfigAccessor) Raw() ([]EquipmentSkillConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EquipmentSkillConfig{}, err
		}
	}
	return a._data, nil
}
