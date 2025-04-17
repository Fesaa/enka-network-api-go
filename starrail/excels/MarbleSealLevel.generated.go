package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MarbleSealLevel struct {
	ID             float64   `json:"ID"`
	Level          float64   `json:"Level"`
	LevelUpDesc    hash.Hash `json:"LevelUpDesc"`
	SkillParamList []float64 `json:"SkillParamList"`
	UnlockSkillID  float64   `json:"UnlockSkillID"`
}
type MarbleSealLevelAccessor struct {
	_data []MarbleSealLevel
}

// LoadData retrieves the data. Must be called before MarbleSealLevel.GroupData
func (a *MarbleSealLevelAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MarbleSealLevel.json")
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
func (a *MarbleSealLevelAccessor) Raw() ([]MarbleSealLevel, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MarbleSealLevel{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}
